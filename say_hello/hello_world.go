package main

import (
	"fmt"
	"os"
)

var Monday = 1

func main (){
	fmt.Println("Hello World!", os.Args[1])
	os.Exit(0)
}