package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Webpage struct {
	Url     string // URL of the webpage
	Title   string // Title of the webpage
	Content string // Content of the webpage
}

// gets links from html webpage tokens
func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		fmt.Printf("%s %s %s", a.Key, a.Namespace, a.Val)
		if a.Key == "href" {
			if len(a.Val) == 0 || !strings.HasPrefix(a.Val, "http") {
				ok = false
				href = a.Val
				return ok, href
			}
			href = a.Val
			ok = true
		}
	}
	return ok, href
}

// reads the webpage and stores data into byte channel
func fetchPage(url string, c chan []byte) {
	client := &http.Client{}

	req, err := client.NewRequest()

	res, err := http.Get(url)
	if err != nil {
		body := []byte("")
		c <- body
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Printf("%s", body)
	if err != nil {
		body = []byte("")
	}
	c <- body
}

// parge channel data and get data and new links
func parseHTML(currUrl string, content []byte, q *Queue, crawled *CrawledSet, db *DatabaseConnection) {
	z := html.NewTokenizer(bytes.NewReader(content))
	tokenCount := 0
	pageContentLength := 0
	body := false
	webpage := Webpage{Url: currUrl, Title: "", Content: ""}

	for {
		if z.Next() == html.ErrorToken || tokenCount > 500 {
			if crawled.size() < 1000 {
				db.insertWebpage(webpage)
			}
			return
		}
		t := z.Token()
		if t.Type == html.StartTagToken {
			if t.Data == "body" {
				body = true
			}
			if t.Data == "javascript" || t.Data == "script" || t.Data == "style" {
				// Skip script and style tags
				z.Next()
				continue
			}
			if t.Data == "title" {
				z.Next()
				title := z.Token().Data // data disappears after z.Token() is called
				webpage.Title = title
				fmt.Printf("Count: %d | %s -> %s\n", crawled.size(), currUrl, title)
			}

			if t.Data == "a" {
				ok, href := getHref(t)
				if !ok {
					continue
				}
				if crawled.contains(href) {
					continue
				} else {
					q.enqueue(href)
				}
			}
		}
		if body && t.Type == html.TextToken && pageContentLength < 500 {
			webpage.Content += strings.TrimSpace(t.Data)
			pageContentLength += len(t.Data)
		}
		tokenCount++
	}
}
