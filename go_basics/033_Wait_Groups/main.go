package main

import (
	"fmt"
	"sync"
)

//with out time.sleep our goroutines run but before complition our main
//func turn of. so we need something to manage our goroutines ,to do that
//we use -> wait groups

func task(id int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("doing task",id)
}

func main() {
   var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
    //its like counter time machanism
	wg.Add(1)
	//the counting of the new start goroutines
	go task(i,&wg)
	}
wg.Wait()
	
}