package service

import (
	"net/http"
	"time"

	dbGorm "gym-backend/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type StaffRequest struct {
	ID          int    `json:"id"`
	FullName    string `json:"fullname"`
	BirthDate   string `json:"birthdate"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	RoleID      uint   `json:"roleid"`
	Gender      int    `json:"gender"`
	Email       string `json:"email"`
	BeginDay    string `json:"beginday"`
	StaffTypeID uint   `json:"stafftypeid"`
}

type StaffID struct {
	ID int `json:"id"`
}

const (
	errMessage = "Email đã được đăng ký"
)

// CreateStaff is used to create new account
func CreateStaff(c echo.Context, db *gorm.DB) error {
	staff := new(StaffRequest)

	if err := c.Bind(staff); err != nil {
		return err
	}

	var queryStaff dbGorm.Staff
	db.Where(&dbGorm.Staff{Email: staff.Email}).Find(&queryStaff)

	if queryStaff.ID != 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    errMessage,
		})
	}

	format := "02/01/2006"
	dob, _ := time.Parse(format, staff.BirthDate)
	begin, _ := time.Parse(format, staff.BeginDay)

	newStaff := dbGorm.Staff{FullName: staff.FullName, BirthDate: dob, Address: staff.Address, Phone: staff.Phone,
		RoleID: staff.RoleID, Gender: staff.Gender, Email: staff.Email, BeginDay: begin, StaffTypeID: staff.StaffTypeID, IsNew: true}

	var createdStaff dbGorm.Staff
	db.Create(&newStaff).Last(&createdStaff)

	staffID := int(createdStaff.ID)
	newAccount := dbGorm.Account{StaffID: staffID, Username: staff.Email, Password: "password"}
	db.Create(&newAccount)

	// util.SendRegisterMail(db, staff.Email, "password")

	return c.JSON(http.StatusCreated, "Thêm nhân viên mới thành công")
}

func GetStaffWithId(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	staff := dbGorm.Staff{}
	db.Where("id = ?", id).First(&staff)
	if staff.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy nhân viên.",
		})
	}

	return c.JSON(http.StatusOK, staff)
}

func GetAllStaff(c echo.Context, db *gorm.DB) error {
	a := []dbGorm.Staff{}
	db.Where(&dbGorm.Class{Active: true}).Find(&a)
	return c.JSON(http.StatusOK, a)
}

func UpdateStaff(c echo.Context, db *gorm.DB) error {
	staff := new(StaffRequest)

	if err := c.Bind(staff); err != nil {
		return err
	}

	var queryStaff dbGorm.Staff
	db.Where(&dbGorm.Staff{Email: staff.Email}).First(&queryStaff)

	format := "02/01/2006"
	dob, _ := time.Parse(format, staff.BirthDate)
	begin, _ := time.Parse(format, staff.BeginDay)

	queryStaff.FullName = staff.FullName
	queryStaff.BirthDate = dob
	queryStaff.Address = staff.Address
	queryStaff.Phone = staff.Phone
	queryStaff.Gender = staff.Gender
	queryStaff.BeginDay = begin
	queryStaff.RoleID = staff.RoleID
	queryStaff.StaffTypeID = staff.RoleID
	db.Save(&queryStaff)

	return c.JSON(http.StatusOK, "Ok")
}
