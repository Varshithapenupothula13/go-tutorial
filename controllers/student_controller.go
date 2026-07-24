package controllers

import (
	"fmt"

	"go_tutorial/models"
	"go_tutorial/utils"
)

func GetStudent() {
	student := models.Student{
		ID:    1,
		Name:  "Varshitha",
		Email: "varshitha@example.com",
	}

	if utils.IsValidEmail(student.Email) {
		fmt.Println("Student:", student)
	} else {
		fmt.Println("Invalid email")
	}
}