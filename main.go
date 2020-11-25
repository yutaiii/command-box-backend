package main

import (
	"net/http"

	"github.com/yutaiii/command-box-backend/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//routing
	e.GET("/commands", handler.GetCommands())

	// CORS restricted
	// Allows requests from any `http://localhost:8080` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPost, http.MethodDelete},
	}))

	//exec server
	e.Start(":8000")
}
