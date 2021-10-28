package main

import "fmt"

func main() {
	fmt.Println("The program finds the capacity of the three capacitors (see the circuit):")
	var c1, c2, c3, c, a float64
	fmt.Print("Enter capacity of three capacitors:")
	fmt.Scan(&c1, &c2, &c3, c)
	a = 1/c1 + 1/(c2+c3)
	c = 1 / a
	fmt.Print(" The total capacity of the three capacitors is", c)

}
