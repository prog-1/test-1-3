package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter a number:")
	fmt.Scan(&a)
	if a%2 == 0 && a%7 == 0 {
		fmt.Println("PingPong")
	} else if a%7 == 0 {
		fmt.Println("Pong")
	} else if a%2 == 0 {
		fmt.Println("Ping")

	} else {
		fmt.Println(a)
	}

}
