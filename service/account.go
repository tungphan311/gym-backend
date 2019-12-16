package service

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

// Account is used to read create new account
type Account struct {
	ID       int    `json:"id"`
	StaffID  int    `json:"staffId"`
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
