package service

import (
	"net/http"

	dbGorm "gym-backend/dbGorm"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// AccountRequest is used to read create new account
type AccountRequest struct {
	ID       int    `json:"id"`
	StaffID  int    `json:"staffid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginAccount is used to authenticate user
type LoginAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Accounts is used to read list records of account
// type Accounts struct {
// 	Account []Account `json:"accounts"`
// }

// CreateAccount is used to create new account
func CreateAccount(c echo.Context, db *gorm.DB) error {
	account := new(AccountRequest)

	if err := c.Bind(account); err != nil {
		return err
	}

	newAccount := dbGorm.Account{StaffID: account.StaffID, Username: account.Username, Password: account.Password}
	db.Create(&newAccount)

	return c.JSON(http.StatusCreated, "OK")
}

// Login is used to authenticate user
func Login(c echo.Context, db *gorm.DB) error {
	account := new(LoginAccount)

	if err := c.Bind(account); err != nil {
		return err
	}

	query := dbGorm.Account{Username: account.Username, Password: account.Password}
	var user dbGorm.Account

	db.Where(&query).Find(&user)

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// // set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["staffid"] = user.StaffID

	// // Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	// wait for staff table for create token

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

	// return c.JSON(http.StatusOK, user)
}
