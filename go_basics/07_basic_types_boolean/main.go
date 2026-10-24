package main

import (
	"fmt"
)

func main(){
	isLogged :=true;
	isAdmin :=true;
	hasSubscription :=false;

	//AND && 
	canOpenDashboard := isLogged && hasSubscription
	canDeletePost := isAdmin || (isLogged && hasSubscription)
	fmt.Println(canOpenDashboard,canDeletePost)

	age :=5
	isAdult :=age>18
	fmt.Println(isAdult)

}