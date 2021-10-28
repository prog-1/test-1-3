package main

import "fmt"

func main() {
	fmt.Println("The program finds the capacity of the three capacitors.")
	fmt.Println("Enter capacity of three capacitors:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	c1 := 1 / (1/a + 1/(b+c))
	fmt.Println("The total capacity of the three capacitors", c1)
}
