package main

//goroutines is a powerful thing of golang. here it run the code conqurently
//give each task to different light thread rather doing synchronized order
import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task",id)
}

func main() {
  for i:=0; i<10;i++{
	go task(i)
  }


  for i:=0; i<=10; i++{
	go func(i int){
		fmt.Println(i)
	}(i)
  }
  fmt.Println("main already finished")
  time.Sleep(time.Second*2)

}
