package main

import "fmt"

func main() {
	fmt.Println("The program prints Ping, Pong or a number.")
	fmt.Println("Enter a number:")
	var number int
	fmt.Scanln(&number)
	if number%2 == 0 {
		fmt.Println("ping")
	} else if number%7 == 0 {
		fmt.Println("Pong")
	} else if number%2 == 0 && number%7 == 0 {
		fmt.Println("PingPong")
	} else {
		fmt.Println("number")
	}
}
