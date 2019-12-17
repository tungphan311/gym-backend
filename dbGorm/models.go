package dbGorm

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
