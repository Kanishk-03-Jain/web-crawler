package main

import "sync"

type Queue struct {
	totalQueued int        // total urls pushed till now
	number      int        // number of current elements
	elements    []string   // elements list in the queue
	mu          sync.Mutex // Mutex locks
}

// push url into queue
func (q *Queue) enqueue(url string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.elements = append(q.elements, url)
	q.totalQueued++
	q.number++
}

// pop url from queue
func (q *Queue) dequeue() string {
	q.mu.Lock()
	defer q.mu.Unlock()
	url := q.elements[0]
	q.elements = q.elements[1:]
	q.number--
	return url
}

// get size of queue
func (q *Queue) size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.number
}
