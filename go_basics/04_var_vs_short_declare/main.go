package main
import (
	"fmt"
)

var x string = "sangam"
func main(){
	var city string;
	city="London"
	var channel="sangam"//inferred to string

	// := 

	subscribers := 5000
	subscribers =subscribers + 1000
	fmt.Println(city,channel,subscribers)
}