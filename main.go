package main

import (
	"fmt"

	"go_tutorial/routes"
)

func main() {
	fmt.Println("Go Tutorial Application Started")

	routes.SetupRoutes()
}