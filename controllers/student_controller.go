package controllers

import (
	"fmt"

	"go_tutorial/models"
	"go_tutorial/utils"
	"go_tutorial/views"
)
func CreateStudent() {
	student := models.Student{
		ID:    2,
		Name:  "Sindhu",
		Email: "sindhu@example.com",
	}

	fmt.Println("Student created successfully")
	views.DisplayStudent(student)
}

func UpdateStudent() {
	student := models.Student{
		ID:    2,
		Name:  "Sindhu",
		Email: "sindhu@example.com",
	}

	student.Name = "Sindhu Updated"
	student.Email = "sindhuupdated@example.com"

	fmt.Println("Student updated successfully")
	views.DisplayStudent(student)
}

func DeleteStudent() {
	student := models.Student{
		ID:    2,
		Name:  "Sindhu Updated",
		Email: "sindhuupdated@example.com",
	}

	fmt.Println("Deleting student:")
	views.DisplayStudent(student)

	fmt.Println("Student deleted successfully")
}

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