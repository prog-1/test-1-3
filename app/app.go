package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// TODO: What does the program do? User enter the number, programm it scan and do some math actions with it.
func main() {
	var x uint
	fmt.Scan(&x)

	// TODO: What do the variables a, b and c store? it stores new value with special formula
	x, a := x/10, x%10
	x, b := x/10, x%10
	x, c := x/10, x%10

	// TODO: What do the three `if ... { ... }` statements below do? If order is other (not a, b,c) statement rearranges the sequence
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}

	// TODO: What do the `if ... { ... }` statements below do?
	x = 0 // if x is 0 -  statement store it differently
	if a%2 != 0 {
		x = x*10 + a
	}
	if b%2 != 0 {
		x = x*10 + b
	}
	if c%2 != 0 {
		x = x*10 + c
	}
// if b and c not even program store it differently
	fmt.Println(x)
}
