package main

import "fmt"

type Student struct {
    name string
    age  int
}

func main() {
    student1 := Student{
        name: "Varshitha",
        age:  22,
    }

    fmt.Println("Name:", student1.name)
    fmt.Println("Age:", student1.age)
}