package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	file, err := os.Open("./sources.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		uri := scanner.Text()
		err := parseUri(uri)
		if err != nil {
			continue
		}
		getFeed(uri)
	}
}

func parseUri(uri string) error {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		fmt.Println("ERROR: Could not parse", uri)
		return err
	}
	return nil
}

func getFeed(uri string) error {
	res, err := http.Get(uri)
	if err != nil {
		fmt.Println("ERROR: Could not GET", uri)
		fmt.Println("ERROR: Status:", res.Status)
		return err
	}
	readBody(res)
	return nil
}

func readBody(res *http.Response) error {
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("ERROR: Could not read body.")
		return err
	}
	fmt.Printf("%s\n", body)
	return nil
}
