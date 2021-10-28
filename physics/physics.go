package main

import "fmt"

func main() {
	var C, C1, C2, C3 float64
	fmt.Print("Enter capacity of three capacitors:")
	fmt.Scan(&C1, &C2, &C3)
	if C1 < 0 || C2 < 0 || C3 < 0 {
		fmt.Println("The capacitors are invalid.")
	} else {
		C = 1/C1 + 1/(C2+C3)
		C = 1 / C
		fmt.Println("The total capacity of the three capacitors is", C)
	}
}
