package crawlerstats

import (
	"fmt"
	"time"

	CrawledSet "github.com/Kanishk-03-Jain/web-crawler/CrawledSet"
	Queue "github.com/Kanishk-03-Jain/web-crawler/queue"
)

type CrawledStats struct {
	pagesPerMinute        string
	crawledRatioPerMinute string
	startTime             time.Time
}

func (c *CrawledStats) Update(crawled *CrawledSet.CrawledSet, queue *Queue.Queue, t time.Time) {
	c.pagesPerMinute += fmt.Sprintf("%f %d\n", t.Sub(c.startTime).Minutes(), crawled.Size())
	c.crawledRatioPerMinute += fmt.Sprintf("%f %f\n", t.Sub(c.startTime).Minutes(), float64(crawled.Size())/float64(queue.Size()))
}

func (c *CrawledStats) Print() {
	fmt.Println("Pages crawled per minute:")
	fmt.Println(c.pagesPerMinute)
	fmt.Println("Crawl to Queued Ratio per minute:")
	fmt.Println(c.crawledRatioPerMinute)
}
