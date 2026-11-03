package greet
//its a reusable library package not a executable package

import "strings"

//package greet -> is reusuable library package not a executable program

//when i give name the first letter should be capital. that means it is
//exported function
func Hello(name string) string{

	clean :=normalizeName(name)
	return "hello ,"+clean
}

func normalizeName(name string)string{
n:=strings.TrimSpace(name)
if n == ""{
	return "Guest"
}
return strings.ToUpper(n)
}