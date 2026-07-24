package views

import (
	"fmt"

	"go_tutorial/models"
)

func DisplayStudent(student models.Student) {
	fmt.Println("Student Details")
	fmt.Println("ID:", student.ID)
	fmt.Println("Name:", student.Name)
	fmt.Println("Email:", student.Email)
}