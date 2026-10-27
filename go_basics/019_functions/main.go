package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func SumAndProduct(a int,b int) (int, int){
	sum :=a+b
	product :=a * b
	return sum, product
}
func main() {

	sum1 := add(10, 15)
	fmt.Println(sum1);
    s, p := SumAndProduct(6,5)

	//i only want sum not product so i need to give a place holder
	onlySum, _:=SumAndProduct(3,2)

	fmt.Println(onlySum,s,p)
}
