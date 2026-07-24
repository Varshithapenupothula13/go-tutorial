package main

import "fmt"

type Speaker interface {
    Speak()
}

type Person struct {
    name string
}

func (p Person) Speak() {
    fmt.Println(p.name, "is speaking")
}

func main() {
    person := Person{name: "Varshitha"}

    var speaker Speaker = person
    speaker.Speak()
}