package main

import "fmt"

func main() {
    student := map[string]int{
        "Math":    85,
        "Science": 90,
        "English": 88,
    }

    fmt.Println("Marks:", student)
    fmt.Println("Math:", student["Math"])

    student["Go"] = 95
    fmt.Println("After adding Go:", student)
}