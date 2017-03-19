package main

import (
	"encoding/xml"
	"html/template"
)

type Rss struct {
	XmlName       xml.Name `xml:"rss"`
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

func parseRssFeed(content []byte) (Rss, error) {
	feed := Rss{}
	err := xml.Unmarshal(content, &feed)
	return feed, err
}
