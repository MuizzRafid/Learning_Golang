package main

import (
	"errors"
	"fmt"
)

//The defer statement in Go is used to schedule a function call to run just before the surrounding function returns

func main(){


	//defer->resp.body.Close()
	fmt.Println("Case 1: success")
	if err :=doWork(false); err !=nil{
		fmt.Println("error:",err)
	}

}

func doWork(success bool) error{
	//i am getting a flag which is success
	//want to do resource related work
	//start message -> resource acquired
	//clearnup message ->resourse released


	fmt.Println("start: resource acquired")

	//this defer will guarntee this  runs at the end of this func
	//and also run the both the paths
	//- success returns -> defer late
	//errors return -> defer early
	defer fmt.Println("cleanup: resource released")

	if !success{
		return errors.New("something went wrong . i am returning early")

	}
		fmt.Println("work: this work is done")

		return nil
}