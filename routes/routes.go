package routes

import "go_tutorial/controllers"

func SetupRoutes() {
	controllers.GetStudent()
	controllers.CreateStudent()
	controllers.UpdateStudent()
	controllers.DeleteStudent()
}