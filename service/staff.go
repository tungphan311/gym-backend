package service

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
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
	RoleID      uint   `json: roleid`
	Gender      int    `json:"gender"`
	Email       string `json:"email"`
	BeginDay    string `json:"beginday"`
	StaffTypeID uint   `json:"stafftypeid"`
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

	SendRegisterMail(db, staff.Email, "password")

	return c.JSON(http.StatusCreated, "OK")
}

func SendRegisterMail(db *gorm.DB, toEmail string, password string) {
	var email dbGorm.Mail
	db.Where(&dbGorm.Mail{Username: "thanhtunga1lqd@gmail.com"}).Find(&email)

	GMAIL_USERNAME := email.Username
	GMAIL_PASSWORD := email.Password

	gmailAuth := smtp.PlainAuth("", GMAIL_USERNAME, GMAIL_PASSWORD, "smtp.gmail.com")

	t, _ := template.ParseFiles("template/register.html")

	var body bytes.Buffer

	headers := "MINE-version: 1.0;\nContent-Type: text/html;"

	body.Write([]byte(fmt.Sprintf("Subject: GỬI THÔNG TIN ĐĂNG NHẬP\n%s\n\n", headers)))

	t.Execute(&body, struct {
		Username string
		Password string
	}{
		Username: toEmail,
		Password: password,
	})

	smtp.SendMail("smtp.gmail.com:587", gmailAuth, GMAIL_USERNAME, []string{toEmail}, body.Bytes())
}
