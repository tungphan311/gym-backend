package service

import (
	"fmt"
	"net/http"
	"time"

	dbGorm "gym-backend/db"
	util "gym-backend/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type StaffRequest struct {
	ID          int    `json:"id"`
	FullName    string `json:"fullname"`
	BirthDate   string `json:"birthdate"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	RoleID      uint   `json: roleid`
	Gender      int    `json:"gender"`
	Email       string `json:"email"`
	BeginDay    string `json:"beginday"`
	StaffTypeID uint   `json:"stafftypeid"`
}

type StaffID struct {
	ID int `json:"id"`
}

type ErrorResponse struct {
	StatusCode uint
	Message    string
}

const (
	errMessage = "Email is already taken"
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

	util.SendRegisterMail(db, staff.Email, "password")

	return c.JSON(http.StatusCreated, "OK")
}

func GetStaff(c echo.Context, db *gorm.DB) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	roleid := claims["roleid"].(uint)

	fmt.Println("%d", roleid)

	staffID := new(StaffID)

	if err := c.Bind(staffID); err != nil {
		return err
	}

	var queryStaff dbGorm.Staff
	db.Where(&dbGorm.Staff{}, staffID).Find(&queryStaff)

	fmt.Println("%v", queryStaff)

	return c.JSON(http.StatusOK, "Ok")
}
