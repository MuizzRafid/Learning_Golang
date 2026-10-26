package main

import "fmt"

func main() {
	points := map[string]int{
		"a": 10,
		"b": 0,
	}

	fmt.Println("a", points["a"])//10
	fmt.Println("b",points["b"])//0

	//though i don't have c i still getting out of c 0
	fmt.Println("c",points["c"])
   

	//so we have val, ok 
	valB,okB :=points["b"]
	fmt.Println(valB,okB)
	valC,okC :=points["C"]
	fmt.Println(valC,okC)

	if val, ok := points["c"]; ok{
		fmt.Println(val)
	} else{
		fmt.Println("c key is not present in the map")
	}

	prices :=map[string]int{
		"xyz" :500,
		"def" :1000,
	}

	fmt.Println(prices["xyz"])

	total :=0
	for item,price :=range prices{
		fmt.Println(item,price)
		total =total +price
	}

	fmt.Println(total)
	for item :=range prices{
		fmt.Println(item)
	}

	


}
