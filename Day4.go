package main

import (
	"fmt"
)

func main() {
	var matrix [20][20]string
	var a, b int
	fmt.Println("Enter the no.of rows:\n")
	fmt.Scanln(&a)
	fmt.Println("Enter the no.of columns:\n")
	fmt.Scanln(&b)
	fmt.Print("\n Enter the matrix elements:")
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Scanf("%s", &matrix[i][j])
		}
	}
	fmt.Println("The matrix is:\n")
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			fmt.Printf("%s\t\t", matrix[i][j])
		}
		fmt.Println("\n")
	}
}
