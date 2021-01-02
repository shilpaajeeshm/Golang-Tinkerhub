package main

import "fmt"

func main() {
	a := 3
	b := 5
	if a > b {
		if b == 5 {
			fmt.Println("a is smaller ")
		}
	} else {
		fmt.Println("b is smaller")
	}
}
