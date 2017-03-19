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

func parseUri(uri string) (*url.URL, error) {
	parsedUri, err := url.ParseRequestURI(uri)
	return parsedUri, err
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
		parsedUri, err := parseUri(uri)
		if err != nil {
			fmt.Printf("ERROR: Could not parse uri %s.\n", uri)
			continue
		}

		res, err := getFeed(uri)
		if err != nil {
			fmt.Printf("ERROR: Could not GET %s.\n", uri)
			continue
		}

		body, err := readBody(res)
		if err != nil {
			fmt.Printf("ERROR: Could not read body for %s.\n", uri)
			continue
		}

		feed, err := parseRssFeed(body)
		if err != nil {
			fmt.Printf("ERROR: Could not parse RSS feed for %s.\n", uri)
			continue
		}

		if _, err := os.Stat("./feeds/" + parsedUri.Host + ".txt"); err != nil {
			fmt.Printf("Creating file: %s.txt\n", parsedUri.Host)
		} else {
			fmt.Printf("Adding feeds to: %s.txt\n", parsedUri.Host)
			fmt.Printf("%s\n", feed.Title)
		}
	}
}
