package main

import "fmt"

//struct groups related fields into one type
type User struct{
ID int
name string
email string
age int
}
type paypal struct{

}

func (p paypal) pay (amount float32){
  fmt.Println("the amount is ",amount)
}

func (p paypal) refund (amount float32, account string ){
	fmt.Printf("this is amount %v and the account is %v",amount, account)
}

type payment struct{

}

func (p payment) makePayment(amount float32,account string){
	paypalPayment :=paypal{}
	paypalPayment.pay(amount)
	paypalPayment.refund(amount,account)
}

func main(){
	payPayment:=payment{}

	payPayment.makePayment(34324,"shahodat_Gozi");

u1:=User{
	ID:9,
	name: "july",
	email: "j@gmail.com",
	age: 6,
}

fmt.Println(u1,u1.name)
}
