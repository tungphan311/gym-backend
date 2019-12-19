package service

import (
	"net/http"
	"time"

	dbGorm "gym-backend/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type BillRequest struct {
	ID          int     `json:"id"`
	Amount      float64 `json:"amount"`
	CreatedTime string  `json:"createdtime"`
}

func CreateBill(c echo.Context, db *gorm.DB) error {
	r := new(BillRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	format := "02/01/2006"
	createdTime, _ := time.Parse(format, r.CreatedTime)

	n := dbGorm.Bill{
		Amount:      r.Amount,
		CreatedTime: createdTime,
	}
	db.Create(&n)
	return c.JSON(http.StatusCreated, "Thêm hoá đơn mới thành công")
}

func GetBillWithId(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	q := dbGorm.Bill{}
	db.Where("id = ?", id).First(&q)
	if q.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy dữ liệu.",
		})
	}

	return c.JSON(http.StatusOK, q)
}

func GetAllBill(c echo.Context, db *gorm.DB) error {
	a := []dbGorm.Bill{}
	db.Where(&dbGorm.Bill{Active: true}).Find(&a)
	return c.JSON(http.StatusOK, a)
}

func UpdateBill(c echo.Context, db *gorm.DB) error {
	r := new(BillRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	format := "02/01/2006"
	createdTime, _ := time.Parse(format, r.CreatedTime)

	var q dbGorm.Bill
	db.Where("id = ?", r.ID).First(&q)
	q.Amount = r.Amount
	q.CreatedTime = createdTime
	db.Save(&q)

	return c.JSON(http.StatusOK, "OK")
}

func DeactiveBill(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	n := dbGorm.Bill{}
	db.Where("id = ?", id).First(&n)
	if n.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy hoá đơn.",
		})
	}

	n.Active = false
	db.Save(&n)

	return c.JSON(http.StatusOK, "Ok")
}
