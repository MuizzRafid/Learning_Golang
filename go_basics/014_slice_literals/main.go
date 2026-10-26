package main

import "fmt"

func main() {
	//common collection type
	//dynamic and can grow
	//[]types{.....}

	result := []string{"simbah", "coco"}

	fmt.Println(result,result[0],result[len(result)-1]);
	var num []int
	num=append(num,10);
	num=append(num,50);
	num=append(num,30);

	fmt.Println(num)
}
