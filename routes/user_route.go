package routes

import (
	"EchoMongoDB/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	//All Routes related to user come here

	// Create User
	e.POST("/user", controllers.CreateUser)
	// GET User
	e.GET("/user", controllers.GetAUser)
	// Edit User
	e.PUT("/user/:userId", controllers.EditAUser)
	// Delete User
	e.DELETE("/user/:userId", controllers.DeleteAUser)
	// Get All Users
	e.GET("/users", controllers.GetAllUsers)
}
