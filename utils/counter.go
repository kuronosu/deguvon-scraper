package utils

import "sync"

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{0, sync.Mutex{}}
}

type SafeCounter struct {
	i  int
	mu sync.Mutex
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	c.i++
	c.mu.Unlock()
}

func (c *SafeCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.i
}

func (c *SafeCounter) IncAndGet() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.i++
	return c.i
}
