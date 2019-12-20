package service

import (
	"net/http"
	"strconv"

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
	Haspt          bool    `json:"haspt"`
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
		Haspt:          r.Haspt,
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

func GetClassWithClassTypeId(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	q := []dbGorm.Class{}
	db.Where("class_type_id = ?", id).Find(&q)

	if len(q) == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy dữ liệu.",
		})
	}

	return c.JSON(http.StatusOK, q)
}

func GetAllClass(c echo.Context, db *gorm.DB) error {
	a := []dbGorm.Class{}
	//db.Where(&dbGorm.Class{Active: true}).Find(&a)

	haspt, _ := strconv.ParseBool(c.QueryParam("haspt"))
	classtypeid, _ := strconv.ParseUint(c.QueryParam("classtypeid"), 10, 64)
	db.Where(&dbGorm.Class{
		Active:      true,
		Haspt:       haspt,
		ClassTypeID: uint(classtypeid),
	}).Find(&a)

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
	q.Haspt = r.Haspt
	db.Save(&q)

	return c.JSON(http.StatusOK, "Cập nhật thông tin gói tập thành công")
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
