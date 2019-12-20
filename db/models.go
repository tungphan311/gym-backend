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
	Email         string
	Gender        int

	StaffID  uint
	Staff    Staff
	IsActive bool
	Classes  []Class `gorm:"many2many:class_members;"`
	Active   bool    `gorm:"DEFAULT:true"`
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
	Active      bool `gorm:"DEFAULT:true"`
}

type StaffType struct {
	gorm.Model
	Name string `gorm:DEFAULT CHARACTER SET utf8`

	Staffs []Staff
	Active bool `gorm:"DEFAULT:true"`
}

type Account struct {
	gorm.Model
	Username string
	Password string

	StaffID int
	Active  bool `gorm:"DEFAULT:true"`
}

type Role struct {
	gorm.Model
	Name string `gorm:DEFAULT CHARACTER SET utf8`

	Staffs []Staff
	Active bool `gorm:"DEFAULT:true"`
}

type Class struct {
	gorm.Model
	Name           string `gorm:DEFAULT CHARACTER SET utf8`
	Price          float64
	DurationDays   int
	ScheduleString string `gorm:DEFAULT CHARACTER SET utf8`
	ClassTypeID    uint
	Members        []Member `gorm:"many2many:class_members;"`
	Haspt          bool
	Active         bool `gorm:"DEFAULT:true"`
}

type ClassMember struct {
	gorm.Model
	MemberID uint
	ClassID  uint
	Active   bool `gorm:"DEFAULT:true"`
}

type ClassType struct {
	gorm.Model
	Name   string `gorm:DEFAULT CHARACTER SET utf8`
	Active bool   `gorm:"DEFAULT:true"`
}

type Device struct {
	gorm.Model
	Name      string `gorm:DEFAULT CHARACTER SET utf8`
	InputDate time.Time

	DeviceStatusID uint
	DeviceTypeID   uint
	Active         bool `gorm:"DEFAULT:true"`
}

type DeviceStatus struct {
	gorm.Model
	Name string `gorm:DEFAULT CHARACTER SET utf8`

	Devices []Device
	Active  bool `gorm:"DEFAULT:true"`
}

type DeviceType struct {
	gorm.Model
	Name string `gorm:DEFAULT CHARACTER SET utf8`

	Devices []Device
	Active  bool `gorm:"DEFAULT:true"`
}

type Bill struct {
	gorm.Model
	Amount      float64
	MemberID    uint
	StaffID     uint
	ClassID     uint
	CreatedTime time.Time

	Active bool `gorm:"DEFAULT:true"`
}

type Parameter struct {
	gorm.Model
	Name        string `gorm:DEFAULT CHARACTER SET utf8`
	Value       float64
	Description string `gorm:DEFAULT CHARACTER SET utf8`
	Active      bool   `gorm:"DEFAULT:true"`
}
