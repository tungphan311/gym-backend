package server

import (
	"gym-backend/service"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// StartRouter is use to start server
func StartRouter(db *gorm.DB) {
	var (
		PORT string = "5555" //config.Ctx.GetString("port")
		e           = echo.New()
	)

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// config to pass cors policy
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// routes
	api := e.Group("/api")

	api.GET("/intro", func(c echo.Context) error {
		return c.String(http.StatusOK, "OOAD\ngym backend")
	})

	api.POST("/accounts", func(c echo.Context) error {
		return service.CreateAccount(c, db)
	})

	api.POST("/login", func(c echo.Context) error {
		return service.Login(c, db)
	})

	e.Logger.Fatal(e.Start(":" + PORT))
}
