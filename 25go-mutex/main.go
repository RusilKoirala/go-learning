package main

import (
	"fmt"
	"sync"
)

type post struct {
	mu    sync.Mutex
	views int
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views++
}

func main() {
	var wg sync.WaitGroup

	mypost := post{views: 0}

	for range 100 {
		wg.Add(1)
		go mypost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(mypost.views)
}
