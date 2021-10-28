package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program prints Ping, Pong or a number.")
	fmt.Println("Enter a number:")
	var a int
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Print("ping")
	}
	if a%7 == 0 {
		fmt.Print("pong")
	}
}
