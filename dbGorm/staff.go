package dbGorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Staff struct {
	gorm.Model
	FullName    string
	BirthDate   time.Time
	Address     string
	Phone       string
	RoleID      int
	Gender      int
	StaffTypeID int
	Account     Account
}

type Role struct {
	gorm.Model
	Name   string
	Staffs []Staff
}

type StaffType struct {
	gorm.Model
	Name   string
	Staffs []Staff
}

type Account struct {
	gorm.Model
	StaffID  int
	Username string
	Password string
}
