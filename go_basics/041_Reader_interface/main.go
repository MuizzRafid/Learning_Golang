package main

//reader is a interface it has string ,int64 ,int and lots
//of method like Read, ReadAt
import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

func readAll(r io.Reader){
data,err:=io.ReadAll(r)
if err!=nil{
	return
}
fmt.Println(string(data))
}


func main(){
	msg:=strings.NewReader("july is here")
	readAll(msg)
	file,err:=os.Open("./important.txt")
	if err!=nil{
		return
	}
	defer file.Close()
	data,err :=io.ReadAll(file)
    if err !=nil{
		return
	}
	fmt.Println(string(data))
	
}