package util

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"

	"github.com/jinzhu/gorm"
	dbGorm "gym-backend/db"
)

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

	err := smtp.SendMail("smtp.gmail.com:587", gmailAuth, GMAIL_USERNAME, []string{toEmail}, body.Bytes())

	if err != nil {
		log.Fatal(err)
	}
}
