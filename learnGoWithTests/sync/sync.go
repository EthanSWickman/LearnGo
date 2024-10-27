package main

import "sync"

type Counter struct {
	mu  sync.Mutex
	val int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.val
}

func NewCounter() *Counter {
	return &Counter{}
}
