package main

import "fmt"

type paymentr interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymentr
}

func (p payment) makePayment(amount float32, account string) {
	p.gateway.pay(amount)
	p.gateway.refund(amount, account)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("the amount is ",amount)
}

func (p paypal) refund(amount float32, account string){
	fmt.Printf("The amount is %v and the account is %v",amount,account)
	fmt.Println()
}


func main(){
	pay:=paypal{}

	payPayment:=payment{
		gateway: pay,
	}
	payPayment.makePayment(32432,"shahodat Gazi")
	pay.pay(3432)
	pay.refund(2343,"shahodat gazi gg")
}