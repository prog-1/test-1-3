package main

import "fmt"

func main() {
	fmt.Println("The program finds the capacity of the three capacitors (see the circuit):")
	fmt.Println("Enter capacity of three capacitors:")
	var C1, C2, C3 float64
	fmt.Scanln(&C1, &C2, &C3)
	C := 1/C1 + 1/(C2 + C3)
	fmt.Println("The total capacity of the three capacitors is", 1/C)
}