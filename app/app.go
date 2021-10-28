package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// The program finds x, which is > 0
func main() {
	var x uint
	fmt.Scan(&x)

	// a, b and c stores x division without remainder
	x, a := x/10, x%10
	x, b := x/10, x%10
	x, c := x/10, x%10

	// exchange a and b(two times), b and c
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}

	// These lines calculates when x equals 0
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
