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

func parseUri(uri string) error {
	_, err := url.ParseRequestURI(uri)
	return err
}

func getFeed(uri string) (*http.Response, error) {
	res, err := http.Get(uri)
	return res, err
}

func readBody(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return body, err
}

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
			fmt.Printf("%s %s.\n", "ERROR: Could not parse uri", uri)
			continue
		}

		res, err := getFeed(uri)
		if err != nil {
			fmt.Printf("%s %s.\n", "ERROR: Could not GET", uri)
			continue
		}

		body, err := readBody(res)
		if err != nil {
			fmt.Printf("%s %s.\n", "ERROR: Could not read body for", uri)
			continue
		}

		feed, err := parseRssFeed(body)
		if err != nil {
			fmt.Printf("%s %s.\n", "ERROR: Could not parse RSS feed for", uri)
			continue
		}

		fmt.Println(feed.Title)
	}
}
