package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// TODO: What does the program do?
// The program writes out odd digits from the last 3 digits of a number.
func main() {
	var x uint
	fmt.Scan(&x)

	// TODO: What do the variables a, b and c store?
	x, a := x/10, x%10
	// a store the last digit of x, then x loses the last digit.
	x, b := x/10, x%10
	// b store the last digit of x, then x loses the last digit.
	x, c := x/10, x%10
	// c store the last digit of x, then x loses the last digit.

	// TODO: What do the three `if ... { ... }` statements below do?
	if a > b {
		a, b = b, a
	}
	// If a > b, then values of a and b ​​are swapped.
	if b > c {
		b, c = c, b
	}
	// If b > c, then values of b and c ​​are swapped.
	if a > b {
		a, b = b, a
	}
	// If a > b, then values of a and b ​​are swapped.

	// TODO: What do the `if ... { ... }` statements below do?
	x = 0
	if a%2 != 0 {
		x = x*10 + a
	}
	// If a is odd, x changes to x multiplied by 10 plus a.
	if b%2 != 0 {
		x = x*10 + b
	}
	// If b is odd, x changes to x multiplied by 10 plus b.
	if c%2 != 0 {
		x = x*10 + c
	}
	// If c is odd, x changes to x multiplied by 10 plus c.

	fmt.Println(x)
}
