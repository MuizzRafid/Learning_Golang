package main

import "fmt"

func Counter() func() int {
	var count int= 0

	return func() int {
		//the return count is accessed by the outter function 
		//so the variable is always stays with the inner function
		//this is called clouser
		count++
		return count
	}
}

func main() {

	increment := Counter()

	fmt.Println(increment())
	fmt.Println(increment())
}