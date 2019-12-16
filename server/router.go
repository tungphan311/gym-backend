package server

import (
	"database/sql"
	"gym-backend/service"
	"net/http"

	"github.com/labstack/echo"
)

// StartRouter is use to start server
func StartRouter(db *sql.DB) {
	var (
		PORT string = "5555" //config.Ctx.GetString("port")
		e           = echo.New()
	)

	// routes
	api := e.Group("/api")

	api.GET("/intro", func(c echo.Context) error {
		return c.String(http.StatusOK, "OOAD\ngym backend")
	})

	api.POST("/accounts", func(c echo.Context) error {
		return service.CreateAccount(c, db)
	})

	e.Logger.Fatal(e.Start(":" + PORT))
}
