package main

import (
	"fmt"
	_ "unsafe"
)

func main() {
	var n1 int = 10
	fmt.Println(&n1)
	var ptr *int = &n1
	fmt.Println(ptr)
}
