package dbGorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Staff struct {
	gorm.Model
	FullName  string
	BirthDate time.Time
	Address   string
	Phone     string
	Role      Role
	Gender    int
	Type      StaffType
}

type Role struct {
	gorm.Model
	Name string
}

type StaffType struct {
	gorm.Model
	Name string
}
