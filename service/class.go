package service

import (
	"net/http"

	dbGorm "gym-backend/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type ClassRequest struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	DurationDays   int     `json:"durationdays"`
	ScheduleString string  `json:"schedulestring"`
	ClassTypeID    uint    `json:"classtypeid"`
	StaffID        uint    `json:"staffid"`
}

func CreateClass(c echo.Context, db *gorm.DB) error {
	r := new(ClassRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	n := dbGorm.Class{
		Name:           r.Name,
		Price:          r.Price,
		DurationDays:   r.DurationDays,
		ScheduleString: r.ScheduleString,
		ClassTypeID:    r.ClassTypeID,
		StaffID:        r.StaffID,
	}
	db.Create(&n)
	return c.JSON(http.StatusCreated, "Thêm gói tập mới thành công")
}

func GetClassWithId(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	q := dbGorm.Class{}
	db.Where("id = ?", id).First(&q)
	if q.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy dữ liệu.",
		})
	}

	return c.JSON(http.StatusOK, q)
}

func GetAllClass(c echo.Context, db *gorm.DB) error {
	a := []dbGorm.Class{}
	db.Where(&dbGorm.Class{Active: true}).Find(&a)
	return c.JSON(http.StatusOK, a)
}

func UpdateClass(c echo.Context, db *gorm.DB) error {
	r := new(ClassRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	var q dbGorm.Class
	db.Where("id = ?", r.ID).First(&q)

	q.Name = r.Name
	q.Price = r.Price
	q.DurationDays = r.DurationDays
	q.ScheduleString = r.ScheduleString
	q.ClassTypeID = r.ClassTypeID
	q.StaffID = r.StaffID
	db.Save(&q)

	return c.JSON(http.StatusOK, "OK")
}

func DeactiveClass(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	n := dbGorm.Class{}
	db.Where("id = ?", id).First(&n)
	if n.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy gói tập.",
		})
	}

	n.Active = false
	db.Save(&n)

	return c.JSON(http.StatusOK, "Ok")
}
