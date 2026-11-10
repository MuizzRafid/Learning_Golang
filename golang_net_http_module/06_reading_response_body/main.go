package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/2"

	resp, err := http.Get(url);
	if err !=nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode !=http.StatusOK{
		fmt.Println(resp.Status)
		return
	}

	bodyBytes,err :=io.ReadAll(resp.Body)
	if err !=nil{
		fmt.Println(err)
		return 
	}

	bodytext:=string(bodyBytes)
	max:=250
	if(len(bodytext)<max){
		max=len(bodytext)
	}

	fmt.Println(bodytext[:max])
}
