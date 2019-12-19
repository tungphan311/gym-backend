package service

import (
	"net/http"

	db "gym-backend/db"

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

type ChangePasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	StaffID  uint   `json:"staffid"`
}

// Accounts is used to read list records of account
// type Accounts struct {
// 	Account []Account `json:"accounts"`
// }

// CreateAccount is used to create new account
func CreateAccount(c echo.Context, dbGorm *gorm.DB) error {
	account := new(AccountRequest)

	if err := c.Bind(account); err != nil {
		return err
	}

	newAccount := db.Account{StaffID: account.StaffID, Username: account.Username, Password: account.Password}
	dbGorm.Create(&newAccount)

	return c.JSON(http.StatusCreated, "OK")
}

// Login is used to authenticate user
func Login(c echo.Context, dbGorm *gorm.DB) error {
	account := new(LoginAccount)

	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Tên đăng nhập hoặc mật khẩu không đúng. Vui lòng kiểm tra lại",
		})
	}

	var (
		user  db.Account
		staff db.Staff
	)

	dbGorm.Where("username = ?", account.Username).First(&user)
	dbGorm.Where("id = ?", user.StaffID).First(&staff)

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// // set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["name"] = staff.FullName
	claims["isnew"] = staff.IsNew
	claims["roleid"] = staff.RoleID

	// // Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	// wait for staff table for create token

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func ChangePassword(c echo.Context, dbGorm *gorm.DB) error {
	account := new(ChangePasswordRequest)

	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Thông tin không hợp lệ",
		})
	}

	var queryAccount db.Account
	dbGorm.Where(&db.Account{Username: account.Username}).First(&queryAccount)

	queryAccount.Password = account.Password
	dbGorm.Save(&queryAccount)

	var staffQuery db.Staff
	dbGorm.Where("id = ?", account.StaffID).First(&staffQuery)

	staffQuery.IsNew = false
	dbGorm.Save(&staffQuery)

	return c.JSON(http.StatusOK, "Cập nhật mật khẩu thành công")
}
