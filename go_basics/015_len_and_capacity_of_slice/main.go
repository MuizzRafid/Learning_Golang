package main

import "fmt"

func main(){
	//make([],lenght,capacity)
	scores :=make([]int,0,5)
	fmt.Println(scores,len(scores),cap(scores))

	scores=append(scores,100)
	fmt.Println("after appending 100", scores)
	
	scores=append(scores, 200, 3000)
	fmt.Println("after appending 200", scores)
	
	scores=append(scores,45,55)
	fmt.Println("after appending 45 55", scores)

	//if in case we r excedding capacity, go grows the backing array (usually double)
	scores=append(scores,60)
	fmt.Println("after appending 60",scores,len(scores),cap(scores))

	todos := []string{"do youtube","workout everyday"}

	more := []string{"learn golang"}

	todos = append(todos, more...)
	fmt.Println(todos)


}
