package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Mail struct {
	Username string
	Password string
}

type Member struct {
	gorm.Model
	FullName      string `gorm:DEFAULT CHARACTER SET utf8`
	BirthDate     time.Time
	Address       string `gorm:DEFAULT CHARACTER SET utf8`
	Phone         string
	IdentityCard  string
	ExpirationDay time.Time

	StaffID        uint
	Staff          Staff
	MemberStatusID uint
	Classes        []Class `gorm:"many2many:class_members;"`
}

type MemberStatus struct {
	gorm.Model
	Name string `gorm:DEFAULT CHARACTER SET utf8`

	Members []Member
}

type Staff struct {
	gorm.Model
	FullName  string `gorm:DEFAULT CHARACTER SET utf8`
	BirthDate time.Time
	Address   string `gorm:DEFAULT CHARACTER SET utf8`
	Phone     string
	Gender    int
	Email     string
	BeginDay  time.Time
	IsNew     bool

	RoleID      uint
	StaffTypeID uint
}

type StaffType struct {
	gorm.Model
	Name string `gorm:DEFAULT CHARACTER SET utf8`

	Staffs []Staff
}

type Account struct {
	gorm.Model
	Username string
	Password string

	StaffID int
}

type Role struct {
	gorm.Model
	Name string `gorm:DEFAULT CHARACTER SET utf8`

	Staffs []Staff
}

type Class struct {
	gorm.Model
	Name           string `gorm:DEFAULT CHARACTER SET utf8`
	Price          float64
	DurationDays   int
	ScheduleString string `gorm:DEFAULT CHARACTER SET utf8`

	ClassTypeID uint
	StaffID     uint
	Members     []Member `gorm:"many2many:class_members;"`
}

type ClassMember struct {
	gorm.Model
	MemberID uint
	ClassID  uint
}

type ClassType struct {
	gorm.Model
	Name string `gorm:DEFAULT CHARACTER SET utf8`
}

type Device struct {
	gorm.Model
	Name      string `gorm:DEFAULT CHARACTER SET utf8`
	InputDate time.Time

	DeviceStatusID uint
	DeviceTypeID   uint
}

type DeviceStatus struct {
	gorm.Model
	Name string `gorm:DEFAULT CHARACTER SET utf8`

	Devices []Device
}

type DeviceType struct {
	gorm.Model
	Name string `gorm:DEFAULT CHARACTER SET utf8`

	Devices []Device
}

type Bill struct {
	gorm.Model
	Amount float64

	BillTypeID uint
}

type BillType struct {
	gorm.Model
	Name string `gorm:DEFAULT CHARACTER SET utf8`
}

type Parameter struct {
	gorm.Model
	Name        string `gorm:DEFAULT CHARACTER SET utf8`
	Value       float64
	Description string `gorm:DEFAULT CHARACTER SET utf8`
}
