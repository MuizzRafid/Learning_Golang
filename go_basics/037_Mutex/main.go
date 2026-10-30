package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex
}
//race condition -> when multiple process try to change or write on a single thing
//here every goroutine try to write on views. when other works on that some misses
//to write . so the output change everytime i run. to avoid this we use mutex
//it waits other goroutine so untill the running one not finished yet
func (p *post) inc(wg *sync.WaitGroup) {
	defer  func(){

		wg.Done()
		p.mu.Unlock()
	} ()
	p.mu.Lock()
	p.views += 1
	
}
func main() {
	var wg sync.WaitGroup
	myPost := post{views:0}
	for i:=0; i<100; i++{
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)

}
