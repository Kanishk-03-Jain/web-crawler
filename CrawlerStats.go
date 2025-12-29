package main

import (
	"fmt"
	"time"
)

type CrawledStats struct {
	pagesPerMinute        string    // how many pages crawled per minute
	crawledRatioPerMinute string    // how many crawled out of how many available in queue
	startTime             time.Time // start time
}

// Adding stats logs
func (c *CrawledStats) update(crawled *CrawledSet, queue *Queue, t time.Time) {
	c.pagesPerMinute += fmt.Sprintf("%f %d\n", t.Sub(c.startTime).Minutes(), crawled.size())
	c.crawledRatioPerMinute += fmt.Sprintf("%f %f\n", t.Sub(c.startTime).Minutes(), float64(crawled.size())/float64(queue.size()))
}

// printing stats
func (c *CrawledStats) print() {
	fmt.Println("Pages crawled per minute:")
	fmt.Println(c.pagesPerMinute)
	fmt.Println("Crawl to Queued Ratio per minute:")
	fmt.Println(c.crawledRatioPerMinute)
}
