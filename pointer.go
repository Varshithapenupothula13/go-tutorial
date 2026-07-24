package main

import "fmt"

func main() {
    number := 10
    pointer := &number

    fmt.Println("Value:", number)
    fmt.Println("Address:", pointer)
    fmt.Println("Value using pointer:", *pointer)

    *pointer = 20
    fmt.Println("Updated value:", number)
}