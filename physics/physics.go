package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program finds the capacity of the three capacitors")
	fmt.Println("Enter capacity of three capacitors:")
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a <= 0 {
		return
	} else if b <= 0 {
		return
	} else if c <= 0 {
		return
	} else {
		o := 0
		o = 1/a + 1/(b+c)
		fmt.Println(o)
	}
}
