package service

import (
	"net/http"
	"time"

	dbGorm "gym-backend/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type MemberRequest struct {
	ID            int    `json:"id"`
	FullName      string `json:"fullname"`
	BirthDate     string `json:"birthdate"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
	IdentityCard  string `json:"identitycard"`
	ExpirationDay string `json:"expirationday"`
	StaffID       uint   `json:"staffid"`
	IsActive      bool   `json:"isactive"`
}

func CreateMember(c echo.Context, db *gorm.DB) error {
	r := new(MemberRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	format := "02/01/2006"
	dob, _ := time.Parse(format, r.BirthDate)
	exd, _ := time.Parse(format, r.ExpirationDay)

	n := dbGorm.Member{
		FullName:      r.FullName,
		BirthDate:     dob,
		Address:       r.Address,
		Phone:         r.Phone,
		IdentityCard:  r.IdentityCard,
		ExpirationDay: exd,
		StaffID:       r.StaffID,
		IsActive:      r.IsActive,
	}
	db.Create(&n)
	return c.JSON(http.StatusCreated, "Thêm hội viên mới thành công")
}

func GetMemberWithId(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	q := dbGorm.Member{}
	db.Where("id = ?", id).First(&q)
	if q.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy dữ liệu.",
		})
	}

	return c.JSON(http.StatusOK, q)
}

func GetAllMember(c echo.Context, db *gorm.DB) error {
	a := []dbGorm.Member{}
	db.Where(&dbGorm.Member{Active: true}).Find(&a)
	return c.JSON(http.StatusOK, a)
}

func UpdateMember(c echo.Context, db *gorm.DB) error {
	r := new(MemberRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	format := "02/01/2006"
	dob, _ := time.Parse(format, r.BirthDate)
	exd, _ := time.Parse(format, r.ExpirationDay)
	var q dbGorm.Member
	db.Where("id = ?", r.ID).First(&q)

	q.FullName = r.FullName
	q.BirthDate = dob
	q.Address = r.Address
	q.Phone = r.Phone
	q.IdentityCard = r.IdentityCard
	q.ExpirationDay = exd
	q.StaffID = r.StaffID
	q.IsActive = r.IsActive
	db.Save(&q)
	return c.JSON(http.StatusOK, "OK")
}

func DeactiveMember(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	n := dbGorm.Member{}
	db.Where("id = ?", id).First(&n)
	if n.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy hội viên.",
		})
	}

	n.Active = false
	db.Save(&n)

	return c.JSON(http.StatusOK, "Ok")
}
