package main

import "fmt"

func main(){
	fmt.Println("The program prints Ping, Pong or a number.")
	fmt.Println("Enter a number:")
	var x int
	fmt.Scan(&x)
	if x%2 == 0{
		fmt.Println("Ping")
	} else if {
		x%7 == 0{
			fmt.Println("Pong")
	} else if x%2 == 0 && x%7 == 0{
	    	fmt.Println("PingPong")
	} else{
			fmt.Println(x)
		}
	}
}