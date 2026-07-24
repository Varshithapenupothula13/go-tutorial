package controllers

import (
	"fmt"

	"go_tutorial/models"
)

func GetStudent() {
	student := models.Student{
		ID:    1,
		Name:  "Varshitha",
		Email: "varshitha@example.com",
	}

	fmt.Println("Student:", student)
}