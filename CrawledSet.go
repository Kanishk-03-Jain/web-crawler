package main

import (
	"hash/fnv"
	"sync"
)

type CrawledSet struct {
	data   map[uint64]bool // map [hash -> true]
	number int             // size of set
	mu     sync.Mutex
}

func (c *CrawledSet) add(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[hashUrl(url)] = true
	c.number++
}

func (c *CrawledSet) contains(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.data[hashUrl(url)]
}

func (c *CrawledSet) size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.number
}

func hashUrl(url string) uint64 { // non-cryptographic hash
	h := fnv.New64a()
	h.Write([]byte(url))
	return h.Sum64()
}
