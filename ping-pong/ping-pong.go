package main

import (
    "fmt"
)

func main() {
    fmt.Println("The program prints Ping, Pong or a number.")
    fmt.Print("Enter a number: ")
    var a int
    fmt.Scan(&a)
    if a % 2 == 0{
        fmt.Println("Ping")
    }else{
        if a % 7 == 0{
            fmt.Println("Pong")
        }else{
            if a % 2 == 0 && a % 7 == 0{
                fmt.Println("PingPong")
            } else {
                fmt.Println(a)
            }
        }
    }
}