package cache

import "sync"

//inMemoryCounter - inmemory counter struct
type inMemoryCounter struct {
	mx  sync.RWMutex
	cnt int64
}

//newCounter - constructor
func newCounter() *inMemoryCounter {
	return &inMemoryCounter{}
}

//Inc - increment counter
func (c *inMemoryCounter) Inc() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.cnt++
}

//Dec - decrement counter
func (c *inMemoryCounter) Dec() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.cnt--
}

//Count - get counter
func (c *inMemoryCounter) Count() int64 {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.cnt
}
