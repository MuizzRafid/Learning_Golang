package main

//encode is writing the json and returning back

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	res :=map[string]any{
		"ok":true,
		"message":"JSON encode successful",
		"datetime": time.Now().UTC(),
	}
       w.Write([]byte("Hello there, here i am"))
	_= json.NewEncoder(w).Encode(res)
}

func main(){
   http.HandleFunc("/ok",successHandler)
	err:=http.ListenAndServe(":5000",nil)
	fmt.Println(err)
}