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
	FullName      string
	BirthDate     time.Time
	Address       string
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
	Name string

	Members []Member
}

type Staff struct {
	gorm.Model
	FullName  string
	BirthDate time.Time
	Address   string
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
	Name string

	Staffs []Staff
}

type Account struct {
	gorm.Model
	Username string
	Password string

	StaffID int
}

type Permission struct {
	gorm.Model
	Name string

	Roles []Role `gorm: "many2many:roles_permissions;"`
}

type Role struct {
	gorm.Model
	Name string

	Staffs      []Staff
	Permissions []Permission `gorm: "many2many:roles_permissions;"`
}

type RolePermission struct {
	gorm.Model

	RoleID       uint
	PermissionID uint
}

type Class struct {
	gorm.Model
	Name           string
	Price          float64
	DurationDays   int
	ScheduleString string

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
	Name string
}

type Device struct {
	gorm.Model
	Name      string
	InputDate time.Time

	DeviceStatusID uint
	DeviceTypeID   uint
}

type DeviceStatus struct {
	gorm.Model
	Name string

	Devices []Device
}

type DeviceType struct {
	gorm.Model
	Name string

	Devices []Device
}

type Bill struct {
	gorm.Model
	Amount float64

	BillTypeID uint
}

type BillType struct {
	gorm.Model
	Name string
}

type Parameter struct {
	gorm.Model
	Name        string
	Value       float64
	Description string
}
