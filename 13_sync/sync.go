package _3_sync

import "sync"

type counter struct {
	mu    sync.Mutex
	value int
}

type Counter interface {
	Inc()
	Value() int
}

func (c *counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *counter) Value() int {
	return c.value
}

func NewCounter() Counter {
	return &counter{}
}
