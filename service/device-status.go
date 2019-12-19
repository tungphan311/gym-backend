package service

import (
	"net/http"

	dbGorm "gym-backend/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type DeviceStatusRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateDeviceStatus(c echo.Context, db *gorm.DB) error {
	r := new(DeviceStatusRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	n := dbGorm.DeviceStatus{
		Name: r.Name,
	}
	db.Create(&n)
	return c.JSON(http.StatusCreated, "Thêm trạng thái thiết bị mới thành công")
}

func GetDeviceStatusWithId(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	q := dbGorm.DeviceStatus{}
	db.Where("id = ?", id).First(&q)
	if q.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy dữ liệu.",
		})
	}

	return c.JSON(http.StatusOK, q)
}

func GetAllDeviceStatus(c echo.Context, db *gorm.DB) error {
	a := []dbGorm.DeviceStatus{}
	db.Where(&dbGorm.DeviceStatus{Active: true}).Find(&a)
	return c.JSON(http.StatusOK, a)
}

func UpdateDeviceStatus(c echo.Context, db *gorm.DB) error {
	r := new(DeviceStatusRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	var q dbGorm.DeviceStatus
	db.Where("id = ?", r.ID).First(&q)
	q.Name = r.Name
	db.Save(&q)

	return c.JSON(http.StatusOK, "OK")
}

func DeactiveDeviceStatus(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	n := dbGorm.DeviceStatus{}
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
