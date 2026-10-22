package main

import (
	"fmt"
	"os"
)

// main() is the entry point
// run a go file->go looks for package main and func main

func main(){
	os.Stdout.Write([]byte("i am also hello ->"))
    fmt.Println("<- Hello go");
}