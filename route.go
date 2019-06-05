package main

import (
	"github.com/alexyslozada/mexico/usuario"
	"github.com/labstack/echo"
)

func startRoutes(e *echo.Echo) {
	// e.POST("/zapatos", zapato.Create, usuario.ValidateJWT)
	// e.POST("/usuarios", usuario.Login)
	e.POST("/api/v1/users", usuario.Create)
	e.GET("/api/v1/users", usuario.GetAll)
	e.GET("/api/v1/users/:email", usuario.GetByEmail)
	e.GET("/api/v1/users-paginate", usuario.GetAllPaginate)
}
