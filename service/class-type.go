package service

import (
	"net/http"

	dbGorm "gym-backend/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type ClassTypeRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateClassType(c echo.Context, db *gorm.DB) error {
	r := new(ClassTypeRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	n := dbGorm.ClassType{
		Name: r.Name,
	}
	db.Create(&n)
	return c.JSON(http.StatusCreated, "Thêm loại gói tập mới thành công")
}

func GetClassTypeWithId(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	q := dbGorm.ClassType{}
	db.Where("id = ?", id).First(&q)
	if q.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy loại gói tập.",
		})
	}

	return c.JSON(http.StatusOK, q)
}

func GetAllClassType(c echo.Context, db *gorm.DB) error {
	a := []dbGorm.ClassType{}
	db.Where(&dbGorm.ClassType{Active: true}).Find(&a)
	return c.JSON(http.StatusOK, a)
}

func UpdateClassType(c echo.Context, db *gorm.DB) error {
	r := new(ClassTypeRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	var q dbGorm.ClassType
	db.Where("id = ?", r.ID).First(&q)
	q.Name = r.Name
	db.Save(&q)

	return c.JSON(http.StatusOK, "OK")
}

func DeactiveClassType(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	n := dbGorm.ClassType{}
	db.Where("id = ?", id).First(&n)
	if n.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy loại gói tập.",
		})
	}

	n.Active = false
	db.Save(&n)

	return c.JSON(http.StatusOK, "Ok")
}
