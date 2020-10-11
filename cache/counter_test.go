package cache

import (
	"sync"
	"testing"
)

const (
	iter = 10000
)

func TestInMemoryCounter_Inc(t *testing.T) {
	c := newCounter()
	wg := sync.WaitGroup{}
	for i := 0; i < iter; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	cnt := c.Count()
	if cnt != iter {
		t.Fatalf("got: %d, expected: %d", cnt, iter)
	}
}

func TestInMemoryCounter_Dec(t *testing.T) {
	c := newCounter()
	wg := sync.WaitGroup{}
	for i := 0; i < iter; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()

	for i := 0; i < iter-1; i++ {
		wg.Add(1)
		go func() {
			c.Dec()
			wg.Done()
		}()
	}
	wg.Wait()

	cnt := c.Count()
	if cnt != 1 {
		t.Fatalf("got: %d, expected: %d", cnt, 1)
	}
}

func TestInMemoryCounter_Count(t *testing.T) {
	TestInMemoryCounter_Inc(t)
}
