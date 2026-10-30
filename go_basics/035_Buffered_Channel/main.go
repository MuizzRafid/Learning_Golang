package main

//A buffered channel in Go is a channel that has a fixed-size queue
//(capacity) to hold data. Unlike default, unbuffered channels—which require
//an active receiver to be ready at the exact same moment a sender transmits
//data—a buffered channel allows a goroutine to send multiple values without
// waiting for a receiver, as long as the buffer isn't completely full

import (
	"fmt"
	"time"
)

//buffered channel
//<-chan = only receive channel
//chan<- =only send channel

func emailSender(emailChan <-chan string,done chan<- bool){
	
	defer func (){done<-true}()
	for email:=range emailChan{
		fmt.Println("sending email to",email)
		time.Sleep(time.Second)
	}
}

func main(){
	emailChan:=make(chan string,5)
	done:=make(chan bool)

//   emainChan<-"1@gmail.com"
//   emainChan<-"2@gmail.com"
//   emainChan<-"3@gmail.com"
//   emainChan<-"4@gmail.com"

go emailSender(emailChan,done)
for i :=0; i<5; i++{
	emailChan<-fmt.Sprintf("%d@gmail.com",i)
}
fmt.Println("done sending")
close(emailChan)
 res:=<-done
 fmt.Println(res)


//   fmt.Println(<-emainChan)
//   fmt.Println(<-emainChan)
//   fmt.Println(<-emainChan)
//   fmt.Println(<-emainChan)

}
