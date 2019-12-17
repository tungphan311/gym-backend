package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type StaffRequest struct {
	ID          int    `json:"id"`
	FullName    string `json:"fullname"`
	BirthDate   string `json:"birthdate"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	RoleID      int    `json: roleid`
	Gender      int    `json:"gender"`
	Email       string `json:"email"`
	BeginDay    string `json:"beginday"`
	StaffTypeID int    `json:"stafftypeid"`
}

// CreateStaff is used to create new account
func CreateStaff(c echo.Context, db *gorm.DB) error {
	staff := new(StaffRequest)

	if err := c.Bind(staff); err != nil {
		return err
	}

	// format := "02/01/06"
	dob, _ := time.Parse(time.RFC3339, staff.BirthDate)

	fmt.Println("staff: %v", dob)

	// newStaff := dbGorm.Staff{FullName: staff.FullName, BirthDate: staff.BirthDate, Address: staff.Address, Phone: staff.Phone,
	// 	RoleID: staff.RoleID, Gender: staff.Gender, Email: staff.Email, BeginDay: staff.BirthDate, StaffTypeID: staff.StaffTypeID}
	// db.Create(&newStaff)

	return c.JSON(http.StatusCreated, "OK")
}
