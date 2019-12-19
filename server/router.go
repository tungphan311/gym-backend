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

	// Account
	api.POST("/accounts", func(c echo.Context) error {
		return service.CreateAccount(c, db)
	})

	api.POST("/login", func(c echo.Context) error {
		return service.Login(c, db)
	})

	api.POST("/change-password", func(c echo.Context) error {
		return service.ChangePassword(c, db)
	})

	// Staffs
	api.POST("/staffs", func(c echo.Context) error {
		return service.CreateStaff(c, db)
	})

	api.PUT("/staffs", func(c echo.Context) error {
		return service.UpdateStaff(c, db)
	})

	api.GET("/staffs", func(c echo.Context) error {
		return service.GetAllStaff(c, db)
	})

	api.GET("/staffs/:id", func(c echo.Context) error {
		return service.GetStaffWithId(c, db)
	})

	api.GET("/staffs/delete/:id", func(c echo.Context) error {
		return service.DeactiveStaff(c, db)
	})

	// r := e.Group("/restricted")
	// r.Use(middleware.JWT([]byte("secret")))

	// r.GET("/staffs/:id", func(c echo.Context) error {
	// 	return service.GetStaff(c, db)
	// })

	// Classes
	api.POST("/classes", func(c echo.Context) error {
		return service.CreateClass(c, db)
	})

	api.PUT("/classes", func(c echo.Context) error {
		return service.UpdateClass(c, db)
	})

	api.GET("/classes", func(c echo.Context) error {
		return service.GetAllClass(c, db)
	})

	api.GET("/classes/:id", func(c echo.Context) error {
		return service.GetClassWithId(c, db)
	})

	api.GET("/classes/delete/:id", func(c echo.Context) error {
		return service.DeactiveClass(c, db)
	})

	// Members
	api.POST("/members", func(c echo.Context) error {
		return service.CreateMember(c, db)
	})

	api.PUT("/members", func(c echo.Context) error {
		return service.UpdateMember(c, db)
	})

	api.GET("/members", func(c echo.Context) error {
		return service.GetAllMember(c, db)
	})

	api.GET("/members/:id", func(c echo.Context) error {
		return service.GetMemberWithId(c, db)
	})

	api.GET("/members/delete/:id", func(c echo.Context) error {
		return service.DeactiveMember(c, db)
	})

	// Devices
	api.POST("/devices", func(c echo.Context) error {
		return service.CreateDevice(c, db)
	})

	api.PUT("/devices", func(c echo.Context) error {
		return service.UpdateDevice(c, db)
	})

	api.GET("/devices", func(c echo.Context) error {
		return service.GetAllDevice(c, db)
	})

	api.GET("/devices/:id", func(c echo.Context) error {
		return service.GetDeviceWithId(c, db)
	})

	api.GET("/devices/delete/:id", func(c echo.Context) error {
		return service.DeactiveDevice(c, db)
	})

	// Devices type
	api.POST("/devicetypes", func(c echo.Context) error {
		return service.CreateDeviceType(c, db)
	})

	api.PUT("/devicetypes", func(c echo.Context) error {
		return service.UpdateDeviceType(c, db)
	})

	api.GET("/devicetypes", func(c echo.Context) error {
		return service.GetAllDeviceType(c, db)
	})

	api.GET("/devicetypes/:id", func(c echo.Context) error {
		return service.GetDeviceTypeWithId(c, db)
	})

	api.GET("/devicetypes/delete/:id", func(c echo.Context) error {
		return service.DeactiveDeviceType(c, db)
	})

	// Devices status
	api.POST("/devicestatus", func(c echo.Context) error {
		return service.CreateDeviceStatus(c, db)
	})

	api.PUT("/devicestatus", func(c echo.Context) error {
		return service.UpdateDeviceStatus(c, db)
	})

	api.GET("/devicestatus", func(c echo.Context) error {
		return service.GetAllDeviceStatus(c, db)
	})

	api.GET("/devicestatus/:id", func(c echo.Context) error {
		return service.GetDeviceStatusWithId(c, db)
	})

	api.GET("/devicestatus/delete/:id", func(c echo.Context) error {
		return service.DeactiveDeviceStatus(c, db)
	})

	// Class type
	api.POST("/classtypes", func(c echo.Context) error {
		return service.CreateClassType(c, db)
	})

	api.PUT("/classtypes", func(c echo.Context) error {
		return service.UpdateClassType(c, db)
	})

	api.GET("/classtypes", func(c echo.Context) error {
		return service.GetAllClassType(c, db)
	})

	api.GET("/classtypes/:id", func(c echo.Context) error {
		return service.GetClassTypeWithId(c, db)
	})

	api.GET("/classtypes/delete/:id", func(c echo.Context) error {
		return service.DeactiveClassType(c, db)
	})

	// Bill
	api.POST("/bills", func(c echo.Context) error {
		return service.CreateBill(c, db)
	})

	api.PUT("/bills", func(c echo.Context) error {
		return service.UpdateBill(c, db)
	})

	api.GET("/bills", func(c echo.Context) error {
		return service.GetAllBill(c, db)
	})

	api.GET("/bills/:id", func(c echo.Context) error {
		return service.GetBillWithId(c, db)
	})

	api.GET("/bills/delete/:id", func(c echo.Context) error {
		return service.DeactiveBill(c, db)
	})

	api.POST("/bills/buy", func(c echo.Context) error {
		return service.BuyClass(c, db)
	})

	e.Logger.Fatal(e.Start(":" + PORT))
}
