package main

import (
	"hash/fnv"
	"sync"
)

type CrawledSet struct {
	data   map[uint64]bool // map [hash -> true]
	number int             // size of set
	mu     sync.Mutex      // Mutex lock
}

// Adding crawled URL to set
func (c *CrawledSet) add(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[hashUrl(url)] = true
	c.number++
}

// checking if URL exists in the set
func (c *CrawledSet) contains(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.data[hashUrl(url)]
}

// getting the size of set
func (c *CrawledSet) size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.number
}

// function for non-cryptographic hash for hashing URL before adding them to set
func hashUrl(url string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(url))
	return h.Sum64()
}
