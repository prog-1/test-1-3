package main

import (
    "fmt"
)

func main() {
    var c, c1, c2, c3 float64
    fmt.Println("The program finds the capacity of the three capacitors (see the circuit):")
    fmt.Print("Enter capacity of three capacitors: ")
    fmt.Scan(&c1)
    fmt.Scan(&c2)
    fmt.Scan(&c3)
    c = 1/(1/c1 + 1/(c2 + c3))
    fmt.Print("The total capacity of the three capacitors is ", c)
}