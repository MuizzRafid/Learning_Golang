package main

import "fmt"
func main(){
	//map[keyType]valueType

	ages := map[string]int{
		"july": 7,
		"simbah":2,
		"downey":4,
	}
   //make(map[k]v)
	fmt.Println(ages["july"],len(ages))
	var scores map[string]int //nil map
	fmt.Println(scores,scores["a"])
	scores=make(map[string]int)
	scores["math"]=90
	fmt.Println(scores)
	//how delete element

	users:= map[string]string{
		"u1":"muizz",
		"u2":"shahadat",
		"u3":"Rahul",
	}

	fmt.Println(users)

	delete(users,"u3")

	delete(users,"u344")
	fmt.Println(users);
}
