package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "gymdb"
)

var (
	types = [2]StaffType{StaffType{Name: "fulltime"}, StaffType{Name: "partime"}}
	roles = [5]Role{Role{Name: "admin"}, Role{Name: "trainer"}, Role{Name: "receptionist"}, Role{Name: "accountant"}, Role{Name: "equipment manager"}}
)

func init() {
	db := Connect()

	db.AutoMigrate(
		&Staff{}, &StaffType{}, &Role{}, &Account{},
		&Permission{}, &Role{}, &RolePermission{},
		&Class{}, &ClassMember{}, &ClassType{},
		&Device{}, &DeviceStatus{}, &DeviceType{},
		&Bill{}, &BillType{}, &Parameter{})

	var staffTypes StaffType
	count := 0
	db.Find(&staffTypes).Count(&count)

	if count == 0 {
		for i := 0; i < len(types); i++ {
			staffType := types[i]
			db.Create(&staffType)
		}
	}

	var role Role
	db.Find(&role).Count(&count)

	if count == 0 {
		for i := 0; i < len(roles); i++ {
			newRole := roles[i]
			db.Create(&newRole)
		}
	}
}

func Connect() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}
