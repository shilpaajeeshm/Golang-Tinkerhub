package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Enter the value of x:")
	var x float64
	fmt.Scanf("%f\n", &x)

	fmt.Print("Enter the value of y:")
	var y float64
	fmt.Scanf("%f\n", &y)

	fmt.Print("Enter the value of z:")
	var z float64
	fmt.Scanf("%f\n", &z)

	fmt.Print("Enter the value of aa:")
	var aa float64
	fmt.Scanf("%f\n", &aa)

	c := ((math.Sqrt(x) + math.Sqrt(y)) * (z)) / aa
	fmt.Printf("Sqrt(%.2f)+Sqrt(%.2f)*%.2f/%.2f =%.2f", x, y, z, aa, c)

}
