package main

import "fmt"


func divide(a int , b int)(tom int, decken int){
	tom =a/b
	decken= a + b

	//this is called named return
	return
}
func main(){
q,r :=divide(15,15)
fmt.Println(q,r)
}
