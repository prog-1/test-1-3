package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number:")
	fmt.Scan(&n)
	if n%2 == 0 && n%7 == 0 {
		fmt.Println("PingPong")
	} else if n%2 == 0 {
		fmt.Println("Ping")
	} else if n%7 == 0 {
		fmt.Println("Pong")
	} else {
		fmt.Println(n)
	}
}
