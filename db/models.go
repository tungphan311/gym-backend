package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Member struct {
	gorm.Model
	Status        MemberStatus
	FullName      string
	BirthDate     time.Time
	Address       string
	Phone         string
	IdentityCard  string
	ExpirationDay time.Time
}

type MemberStatus struct {
	gorm.Model
	Name string
}

type Staff struct {
	gorm.Model
	FullName    string
	BirthDate   time.Time
	Address     string
	Phone       string
	RoleID      int
	Gender      int
	Email       string
	BeginDay    time.Time
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
