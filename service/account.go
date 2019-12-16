package service

import (
	"database/sql"
	"net/http"

	// "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Account is used to read create new account
type Account struct {
	ID       int    `json:"id"`
	StaffID  int    `json:"staffId"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginAccount is used to authenticate user
type LoginAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Accounts is used to read list records of account
type Accounts struct {
	Account []Account `json:"accounts"`
}

// CreateAccount is used to create new account
func CreateAccount(c echo.Context, db *sql.DB) error {
	account := new(Account)

	if err := c.Bind(account); err != nil {
		return err
	}

	sqlStatement := `
	INSERT INTO account (staffid, username, password)
	VALUES ($1, $2, $3)
	RETURNING id`

	id := 0
	err := db.QueryRow(sqlStatement, account.StaffID, account.Username, account.Password).Scan(&id)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, "OK")
}

// Login is used to authenticate user
func Login(c echo.Context, db *sql.DB) error {
	account := new(LoginAccount)

	if err := c.Bind(account); err != nil {
		return err
	}

	sqlStatement := `
	SELECT staffid FROM account 
	WHERE username=$1 
	AND password=$2`

	staffid := 0
	rows, err := db.Query(sqlStatement, account.Username, account.Password)

	if err != nil {
		return echo.ErrUnauthorized
	}

	for rows.Next() {
		if err := rows.Scan(&staffid); err != nil {
			return echo.ErrUnauthorized
		}
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["staffid"] = err

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	// wait for staff table for create token

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
