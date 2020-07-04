package main

import "fmt"

func main() {
	sum := 1
	demo := 1
	for sum < 1000 {
		sum += sum
	}
	for i := 0; i < 1000; i++ {
		demo += i
	}
	fmt.Println(sum)
	fmt.Println(demo)
}
