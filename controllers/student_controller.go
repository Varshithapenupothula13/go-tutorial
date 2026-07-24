package controllers

import (
	"fmt"

	"go_tutorial/models"
	"go_tutorial/utils"
	"go_tutorial/views"
)

func GetStudent() {
	student := models.Student{
		ID:    1,
		Name:  "Varshitha",
		Email: "varshitha@example.com",
	}

	if utils.IsValidEmail(student.Email) {
		views.DisplayStudent(student)
	} else {
		fmt.Println("Invalid email")
	}
}