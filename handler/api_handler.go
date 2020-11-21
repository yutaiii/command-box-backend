package handler

import (
	"log"
	"net/http"

	"github.com/yutaiii/command-box-backend/usecase"

	"github.com/labstack/echo"
)

func GetCommands() echo.HandlerFunc {
	return func(c echo.Context) error {
		commands, err := usecase.GetCommands()
		if err != nil {
			log.Printf("GetCommands err %+v", err)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, commands)
	}
}
