package main

import "fmt"

func main() {
    fruits := []string{"Apple", "Banana", "Mango"}

    fmt.Println("Fruits:", fruits)

    fruits = append(fruits, "Orange")

    fmt.Println("After append:", fruits)
    fmt.Println("First fruit:", fruits[0])
}