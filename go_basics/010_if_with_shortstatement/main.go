package main

import "fmt"

func main(){
	items:=3;
	pricePerItem :=49;

	if total :=items * pricePerItem; total>=100{
		fmt.Println("eligible for shiping");
	}else {
		fmt.Println("not eligible for shiping");
	}
}