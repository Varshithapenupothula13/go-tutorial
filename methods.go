package main

import "fmt"

type Person struct {
    name string
    age  int
}

func (p Person) display() {
    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
}

func main() {
    person1 := Person{
        name: "Varshitha",
        age:  22,
    }

    person1.display()
}