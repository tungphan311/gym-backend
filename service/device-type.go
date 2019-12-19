package service

import (
	"net/http"

	dbGorm "gym-backend/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type DeviceTypeRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateDeviceType(c echo.Context, db *gorm.DB) error {
	r := new(DeviceTypeRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	n := dbGorm.DeviceType{
		Name: r.Name,
	}
	db.Create(&n)
	return c.JSON(http.StatusCreated, "Thêm loại thiết bị mới thành công")
}

func GetDeviceTypeWithId(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	q := dbGorm.DeviceType{}
	db.Where("id = ?", id).First(&q)
	if q.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy dữ liệu.",
		})
	}

	return c.JSON(http.StatusOK, q)
}

func GetAllDeviceType(c echo.Context, db *gorm.DB) error {
	a := []dbGorm.DeviceType{}
	db.Where(&dbGorm.DeviceType{Active: true}).Find(&a)
	return c.JSON(http.StatusOK, a)
}

func UpdateDeviceType(c echo.Context, db *gorm.DB) error {
	r := new(DeviceTypeRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	var q dbGorm.DeviceType
	db.Where("id = ?", r.ID).First(&q)
	q.Name = r.Name
	db.Save(&q)

	return c.JSON(http.StatusOK, "OK")
}

func DeactiveDeviceType(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	n := dbGorm.DeviceType{}
	db.Where("id = ?", id).First(&n)
	if n.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy thiết bị.",
		})
	}

	n.Active = false
	db.Save(&n)

	return c.JSON(http.StatusOK, "Ok")
}
