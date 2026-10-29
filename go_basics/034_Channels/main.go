package main

//4:57
import (
	"fmt"
	"time"
)
func processNum(numChan chan int){
	for num:=range numChan{
		
		fmt.Println("Processing Number",num)
		time.Sleep(time.Second*1)
	}
}

func sum(result chan int,num1 int, num2 int){
numResult:= num1+num2
result <-numResult+5

}

//goroutine synchronizer
func task(done chan bool){
	defer func (){done<-true}()
	fmt.Println("processing...")
}

func main() {
	// create a channel(this type of channel called unbuffered channel)
// 	messageChan := make(chan string)
// 	//sending data from channel
// 	messageChan <- "ping" //channels are blocking
// 	//receiving data from the channel
// 	msg := <-messageChan
// //i create a  channel for sending and receiving but there no where to send
// //so sending and receiving req are in one line which is a deadlock
// 	fmt.Println(msg)
// 	time.Sleep(time.Second*1)

	// numChan:=make(chan int)
	// go processNum(numChan)

	// for{ 
	// 	numChan <- rand.Intn(100)
	// }
	// time.Sleep(time.Second*2)

//   result:=make(chan int)

//   go sum(result,4,5)

//   res := <-result
//   fmt.Println(res)

//   time.Sleep(time.Second*1)

done:=make(chan bool)
go task(done)

<-done//block untill no one seed data to task parameter chanel


}
