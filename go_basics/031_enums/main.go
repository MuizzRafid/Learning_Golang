package main

import "fmt"

type OrderStatus int
type OrderStatusS string
//iota is only for integer
const (
	Received OrderStatus = iota
	Conformed
	Prepared
	Delivered
)
const (
	Recvd OrderStatusS = "recieved"
	Cfd ="confirmed"
	Prd="Prepared"
	Dlvd="Delivered"
	
)
func changeOrderStatus(status OrderStatus,sStatus OrderStatusS) {
	fmt.Println("Changing order status to ", status, sStatus)
}

func main() {
changeOrderStatus(Prepared,Prd)
}

//Use Cases
//Restrict function arguments to meaningful values
//switch statement over a closed set of values
