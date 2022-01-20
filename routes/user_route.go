package routes

import (
	"EchoMongoDB/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	//All Routes related to user come here

	// POST /users
	e.POST("/user", controllers.CreateUser)

	// GET /users
	e.GET("/users", controllers.GetAUser)
}
