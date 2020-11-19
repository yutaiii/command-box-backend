package main

import (
	"github.com/yutaiii/command-box-backend/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//routing
	e.GET("/hello", handler.Hello())

	//exec server
	e.Start(":8000")
}
