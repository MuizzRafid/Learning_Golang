package main

import (
	"fmt"
	"os"
)

func main() {

	f,err:=os.Open("example.txt");
	if err!=nil{
		panic("o mayamaya go")
	}
	defer f.Close()
//Check if file exists
	fileInfo,err:= f.Stat()
	if err!=nil{
		panic("o mayamaya go")
	}
	fmt.Println("the file name is :",fileInfo.Name())
	fmt.Println("file or folder :",fileInfo.IsDir())
	fmt.Println("file size :",fileInfo.Size())
	fmt.Println("file permission :",fileInfo.Mode())
	fmt.Println("file last access time :",fileInfo.ModTime())

	buf:=make([]byte,9 )
	//create a byte slice of length 9 and initialize as 0

	d,err:=f.Read(buf)
	//Read reads upto buf len(9) and store it in buf return two thing d
	//and err, d is->how many byte it could read and the error
	if err!=nil{
		panic(err)
	}
    
	for i:=0; i<len(buf);i++{
    println("data",d,string(buf[i]))

	}
	println("data",d,buf)
}
