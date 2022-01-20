package main

import (
	"EchoMongoDB/configs"
	"EchoMongoDB/routes"

	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":6000"))
}