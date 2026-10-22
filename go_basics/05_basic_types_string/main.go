package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Muizz"
	lastName := "Rafid"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)

	fmt.Println(strings.ToUpper(fullName))

}
