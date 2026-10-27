package main

import "fmt"

//method will receive a copy of user struct
type User struct{
	Name string
	Age int
}
func main(){
u:=User{
	Name: "jjj",
	Age: 6,
}

fmt.Println(u.Intro())
fmt.Println("after",u.Age)

}

//this method receives a copy of the u type User
//(u User)=>Receiver: binds the function to the User struct
//Intro(): The name of the method you call via u.Intro()
//method returns string value
//%s -> for String and %d -> for Decimal/Integer
//so same like float/decimal ->$%.2f , charecter ->%c , for pointer ->%p, for anything -> %v
func (u User) Intro() string{
	u.Age++
	return fmt.Sprintf("Hi , j is %s and i am %d years old",u.Name,u.Age)
}