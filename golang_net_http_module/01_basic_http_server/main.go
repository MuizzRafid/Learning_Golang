package main

//server component has http.Server, http.ListenAndServe objects
//router component has http.serveMux, http.NewServeMux() objects
//client component has http.Client, http.Transport objects

import (
	"fmt"
	"net/http"
)
func helloHandler(w http.ResponseWriter,r *http.Request ){
	if r.Method != http.MethodGet{
		http.Error(w,"only get is allowed",http.StatusMethodNotAllowed)
		return
	}
  //on write -> Write writes the data to the connection as part of an HTTP reply.
	_, _=w.Write([]byte("hello from Go net/http server"))

}

func main(){
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("try going to 6000 port")

	err :=http.ListenAndServe(":6000",nil)
	if err!=nil{

	}
}