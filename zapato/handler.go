package zapato

import (
	"net/http"

	"github.com/alexyslozada/mexico/respuesta"
	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	m := &Model{}
	err := c.Bind(m)
	if err != nil {
		r := respuesta.Model{
			MensajeError: respuesta.MensajeError{
				"E102",
				"El objeto zapato está mal enviado",
			},
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	d := storage.Create(m)
	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			"A001",
			"Zapato creado correctamente",
		},
		Data: d,
	}
	return c.JSON(http.StatusCreated, r)
}
