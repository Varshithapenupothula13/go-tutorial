package main

import (
	"fmt"

	"go_tutorial/database"
	"go_tutorial/routes"
)

func main() {
	fmt.Println("Go Tutorial Application Started")

	database.ConnectDatabase()
	routes.SetupRoutes()
}