package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// TODO: What does the program do?
func main() {
	var x uint
	fmt.Scan(&x)

	// TODO: What do the variables a, b and c store?
	x, a := x/10, x%10
	x, b := x/10, x%10
	x, c := x/10, x%10

	// TODO: What do the three `if ... { ... }` statements below do?
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
	x = 0
	if a%2 != 0 {
		x = x*10 + a
	}
	if b%2 != 0 {
		x = x*10 + b
	}
	if c%2 != 0 {
		x = x*10 + c
	}

	fmt.Println(x)
}
