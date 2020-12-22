package main

import "fmt"

func main() {
	fmt.Println("Enter the value of x:")
	var x int64
	fmt.Scanf("%d\n", &x)
	fmt.Println("Enter the value of y:")
	var y int64
	fmt.Scanf("%d\n", &y)
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y && x == y)
	fmt.Println(!(x <= y))
	fmt.Println(x == y || x < y)

}
