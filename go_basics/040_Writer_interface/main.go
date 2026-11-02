//Writer is the interface that wraps the basic Write method

//Write writes len(p) bytes from p to the underlying data stream
//it returns the number of bytes written from p (0<=n<=len(p))
//and any error encountered that caused the write to stop early
//Write must return a non-nil error if it returns n < len(p)
//Write must not modify the slice data even temporarily

//implementataions must not retain p

package main

import (
	"fmt"
	"io"
	"os"
)

type Writer interface{
	Write(p []byte)(n int, err error)
}

func writeMessage(w io.Writer){
	w.Write([]byte("Hello from io.writer\n"))
}

func main(){
	writeMessage(os.Stdout)
	file,err:=os.Create("./important.txt")//also can create like this os.Create("gg.txt")

	if err!=nil{
      fmt.Println("not successful")
		return 
	}
	defer file.Close()
	writeMessage(file)
}