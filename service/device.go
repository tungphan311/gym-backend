package service

import (
	"net/http"
	"strconv"
	"time"

	dbGorm "gym-backend/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type DeviceRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	InputDate string `json:"inputdate"`

	DeviceStatusID uint   `json:"devicestatusid"`
	DeviceTypeID   uint   `json:"devicetypeid"`
	Description    string `json:"description"`
	Active         bool   `json:"active"`
}

func CreateDevice(c echo.Context, db *gorm.DB) error {
	r := new(DeviceRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	format := "02/01/2006"
	ipd, _ := time.Parse(format, r.InputDate)

	n := dbGorm.Device{
		Name:           r.Name,
		InputDate:      ipd,
		DeviceStatusID: r.DeviceStatusID,
		DeviceTypeID:   r.DeviceTypeID,
		Description:    r.Description,
		Active:         r.Active,
	}
	db.Create(&n)
	return c.JSON(http.StatusCreated, "Thêm thiết bị mới thành công")
}

func GetDeviceWithId(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	q := dbGorm.Device{}
	db.Where("id = ?", id).First(&q)
	if q.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy dữ liệu.",
		})
	}

	return c.JSON(http.StatusOK, q)
}

func GetAllDevice(c echo.Context, db *gorm.DB) error {
	a := []dbGorm.Device{}

	devicetypeid, _ := strconv.ParseUint(c.QueryParam("devicetypeid"), 10, 64)
	devicestatusid, _ := strconv.ParseUint(c.QueryParam("devicestatusid"), 10, 64)
	db.Where(&dbGorm.Device{
		Active:         true,
		DeviceTypeID:   uint(devicetypeid),
		DeviceStatusID: uint(devicestatusid),
	}).Find(&a)

	return c.JSON(http.StatusOK, a)
}

func UpdateDevice(c echo.Context, db *gorm.DB) error {
	r := new(DeviceRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	var q dbGorm.Device
	db.Where("id = ?", r.ID).First(&q)

	format := "02/01/2006"
	ipd, _ := time.Parse(format, r.InputDate)
	q.Name = r.Name
	q.InputDate = ipd
	q.DeviceStatusID = r.DeviceStatusID
	q.DeviceTypeID = r.DeviceTypeID
	q.Description = r.Description
	db.Save(&q)

	return c.JSON(http.StatusOK, "OK")
}

func DeactiveDevice(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	n := dbGorm.Device{}
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
