package main

import "fmt"

func main() {
	fmt.Println("The program prints Ping, Pong or a number.")
	fmt.Println("Enter the number:")
	var a int 
	fmt.Scanln(&a)
	if a%2 == 0 {
		if a%7 == 0 {
			fmt.Println("PingPong")
		} else {
			fmt.Println("Ping")
		}
	} else if a%7 == 0 {
		fmt.Println("Pong")
	} else {
		fmt.Println(a)
	}
}