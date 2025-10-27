package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	wg := new(sync.WaitGroup)
	var c Counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				c.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("最终计数：%d\n", c.Value())
}
