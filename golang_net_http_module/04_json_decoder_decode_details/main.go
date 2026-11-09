package main

//04_json_decoder_decode_details

//decoder is for read the data
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)


func writeJSON(w http.ResponseWriter, status int, data any){
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	_=json.NewEncoder(w).Encode(data)

}

//The backticks ` ` define a raw string literal in Go. Go uses the text inside them as
//metadata associated with the struct field.
type TestRequest struct{
	Name string `json:"name"`
}
func testHandler(w http.ResponseWriter,r *http.Request){
	if r.Method !=http.MethodPost{
      writeJSON(w,http.StatusMethodNotAllowed,map[string]any{
		"ok":false,
		"error": "only post is allowed",
	  })
	}

	defer r.Body.Close()

	var req TestRequest
	dec :=json.NewDecoder(r.Body)
	if err :=dec.Decode(&req); err!=nil{
		writeJSON(w,http.StatusBadRequest,map[string]any{
			"ok" : false,
			"error": "invalid json format",
	})
	return
	}

	req.Name= strings.TrimSpace(req.Name)

	if req.Name ==""{
		writeJSON(w,http.StatusBadRequest, map[string]any{
			"ok":false,
			"error": "name must not be empty!",
		})
		return
	}
	writeJSON(w,http.StatusOK,map[string]any{
		"ok": true,
		"data":req,
		"timestamp": time.Now().UTC(),
	})

}


func main(){
	http.HandleFunc("/test",testHandler)
	err:=http.ListenAndServe(":5000",nil)
	fmt.Println(err)
}