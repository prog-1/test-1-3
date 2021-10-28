package main

import "fmt"

func main(){
	fmt.Println("The program finds the capacity of the three capacitors (see the circuit):")
	fmt.Println("Enter capacity of three capacitors:")
	var c1, c2, c3 int
	fmt.Scan(&c1, &c2, &c3)
	cap := 1/c1 + 1/(c2 + c3)
	fmt.Println("The total capacity of the three capacitors is", 1/cap)
}