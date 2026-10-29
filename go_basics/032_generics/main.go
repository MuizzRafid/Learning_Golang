package main

import "fmt"

func printSlice(items []string) {

	for _, item := range items {
		fmt.Println(item)
	}

}



func printSliceGenerics[T string|int](items []T){
 for _,item:=range items{
	fmt.Println(item)
 }
}


type stack[T any] struct{
	elements []T
}

func main() {
	printSliceGenerics([]string{"alpha", "sigma"})

	mystack:=stack[string]{
		elements : []string{"july"},
	}

	fmt.Println(mystack.elements);
}