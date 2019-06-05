package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	//e.GET("/", holaMundo)
	//e.GET("/mexico", saludaMexico)
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	startRoutes(e)
	err := e.Start(":8080")
	if err != nil {
		fmt.Printf("No pude subir el servidor %v", err)
	}
}

func holaMundo(c echo.Context) error {
	return c.String(http.StatusOK, "hola mundo")
}

func saludaMexico(c echo.Context) error {
	return c.String(http.StatusOK, "Viva México Cuarones!!!")
}
