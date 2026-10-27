package main

import "fmt"

//address operator -> &
//derefence operator -> *
func main(){
	//pointer to shore the memory address of any value

	//&name -> address of name variable(makes a pointer)
	//*p -> deref (go to that address and read/write)

	score:=10
	fmt.Println("before",score)
	addScore(&score)
	fmt.Println("after",score)
}


func addScore(score *int){
  *(score)=20
  
}