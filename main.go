package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Rss struct {
	XMLName       xml.Name `xml:"rss"`
	Version       string   `xml:"version,attr"`
	Title         string   `xml:"channel>title"`
	Link          string   `xml:"channel>link"`
	Description   string   `xml:"channel>description"`
	LastBuildDate string   `xml:"channel>lastBuildDate"`
	ItemList      []Item   `xml:"channel>item"`
}

type Item struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	PubDate     string        `xml:"pubDate"`
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
	parseRssFeed(body)
	return nil
}

func parseRssFeed(content []byte) error {
	feed := Rss{}
	err := xml.Unmarshal(content, &feed)
	if err != nil {
		fmt.Println("ERROR: Could not parse RSS feed.")
		return err
	}
	fmt.Println(feed)
	return nil
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
			continue
		}
		getFeed(uri)
	}
}
