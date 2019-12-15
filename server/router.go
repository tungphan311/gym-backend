package server

import (
	"github.com/labstack/echo"
	"net/http"
)

func StartRouter() {
	var (
		PORT string = "5555" //config.Ctx.GetString("port")
		e           = echo.New()
	)

	// routes
	api := e.Group("/api")

	api.GET("/intro", func(c echo.Context) error {
		return c.String(http.StatusOK, "OOAD\ngym backend")
	})

	e.Logger.Fatal(e.Start(":" + PORT))
}
