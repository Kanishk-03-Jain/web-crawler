package main

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	webArchiveAccess := true
	if godotenv.Load() != nil {
		fmt.Println("Error loading .env file. No access to web archive.")
		webArchiveAccess = false
	}

	db := DatabaseConnection{access: webArchiveAccess, uri: "", client: nil, collection: nil}
	db.connect()

	crawled := CrawledSet{data: make(map[uint64]bool)}
	seed := "https://en.wikipedia.org/wiki/Kanishka"
	queue := Queue{totalQueued: 0, number: 0, elements: make([]string, 0)}

	ticker := time.NewTicker(1 * time.Minute)
	done := make(chan bool)
	crawlerStats := CrawledStats{pagesPerMinute: "0 0\n", crawledRatioPerMinute: "0 0\n", startTime: time.Now()}

	// Tick every minute
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				crawlerStats.update(&crawled, &queue, t)
			}
		}
	}()
	queue.enqueue(seed)
	url := queue.dequeue()
	crawled.add(url)
	c := make(chan []byte)

	go fetchPage(url, c)

	content := <-c
	parseHTML(url, content, &queue, &crawled, &db)

	for queue.size() > 0 && crawled.size() < 5000 {
		url := queue.dequeue()
		crawled.add(url)

		go fetchPage(url, c)
		content := <-c
		if len(content) == 0 {
			continue
		}
		go parseHTML(url, content, &queue, &crawled, &db)
	}
	ticker.Stop()
	done <- true
	db.disconnect()
	fmt.Println("\n------------CRAWLER STATS------------")
	fmt.Printf("Total queued: %d\n", queue.totalQueued)
	fmt.Printf("To be crawled (Queue) size: %d\n", queue.size())
	fmt.Printf("Crawled size: %d\n", crawled.size())
	crawlerStats.print()
}
