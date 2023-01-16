package account

import (
	"github.com/labstack/echo/v4"
)

func Handle(e *echo.Echo, c *ControllerInterface) {
	g := e.Group("/account")

	g.POST("", (*c).CreateAccount)
	g.GET("/:id", (*c).GetAccount)
	g.PUT("/:id", (*c).UpdateAccount)
	g.DELETE("/:id", (*c).DeleteAccount)
}
