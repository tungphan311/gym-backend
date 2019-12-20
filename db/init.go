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

func init() {
	db := Connect()

	db.AutoMigrate(
		&Staff{}, &StaffType{}, &Role{}, &Account{}, &Role{}, &Member{},
		&Class{}, &ClassMember{}, &ClassType{},
		&Device{}, &DeviceStatus{}, &DeviceType{},
		&Bill{}, &Parameter{}, &Mail{})

	initData(db)
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

var (
	types          = [2]StaffType{StaffType{Name: "fulltime"}, StaffType{Name: "partime"}}
	devicestatuses = [3]DeviceStatus{DeviceStatus{Name: "Mới"}, DeviceStatus{Name: "Cũ"}, DeviceStatus{Name: "Đã hỏng"}}
	roles          = [5]Role{Role{Name: "admin"}, Role{Name: "trainer"}, Role{Name: "receptionist"}, Role{Name: "accountant"}, Role{Name: "equipment manager"}}
	classTypes     = [4]ClassType{
		ClassType{
			Name: "THEO NGÀY",
		},
		ClassType{
			Name: "CƠ BẢN",
		},
		ClassType{
			Name: "NÂNG CAO",
		},
		ClassType{
			Name: "CAO CẤP",
		},
	}
	classes = [19]Class{
		Class{Name: "Gói 1 ngày", Price: 50000, DurationDays: 1, ScheduleString: "Cả ngày", ClassTypeID: 1, Haspt: false},
		Class{Name: "Gói 7 ngày", Price: 250000, DurationDays: 7, ScheduleString: "Cả ngày", ClassTypeID: 1, Haspt: false},
		Class{Name: "Gói 15 ngày", Price: 600000, DurationDays: 15, ScheduleString: "Cả ngày", ClassTypeID: 1, Haspt: false},
		Class{Name: "Gói cơ bản 14 tháng (Không HLV)", Price: 9900000, DurationDays: 420, ScheduleString: "Cả ngày", ClassTypeID: 2, Haspt: false},
		Class{Name: "Gói cơ bản 6 tháng (Không HLV)", Price: 6299000, DurationDays: 180, ScheduleString: "Cả ngày", ClassTypeID: 2, Haspt: false},
		Class{Name: "Gói cơ bản 3 tháng (Không HLV)", Price: 3999000, DurationDays: 90, ScheduleString: "Cả ngày", ClassTypeID: 2, Haspt: false},
		Class{Name: "Gói cơ bản 1 tháng (Không HLV)", Price: 1999000, DurationDays: 30, ScheduleString: "Cả ngày", ClassTypeID: 2, Haspt: false},
		Class{Name: "Gói nâng cao 36 tháng (Không HLV)", Price: 21900000, DurationDays: 1080, ScheduleString: "Cả ngày", ClassTypeID: 3, Haspt: false},
		Class{Name: "Gói nâng cao 18 tháng (Không HLV)", Price: 14299000, DurationDays: 540, ScheduleString: "Cả ngày", ClassTypeID: 3, Haspt: false},
		Class{Name: "Gói cao cấp 60 tháng (Không HLV)", Price: 87500000, DurationDays: 1800, ScheduleString: "Cả ngày", ClassTypeID: 4, Haspt: false},
		Class{Name: "Gói cao cấp 24 tháng (Không HLV)", Price: 49900000, DurationDays: 720, ScheduleString: "Cả ngày", ClassTypeID: 4, Haspt: false},

		Class{Name: "Gói cơ bản 14 tháng", Price: 19900000, DurationDays: 420, ScheduleString: "Cả ngày", ClassTypeID: 2, Haspt: true},
		Class{Name: "Gói cơ bản 6 tháng", Price: 12299000, DurationDays: 180, ScheduleString: "Cả ngày", ClassTypeID: 2, Haspt: true},
		Class{Name: "Gói cơ bản 3 tháng", Price: 7999000, DurationDays: 90, ScheduleString: "Cả ngày", ClassTypeID: 2, Haspt: true},
		Class{Name: "Gói cơ bản 1 tháng", Price: 3999000, DurationDays: 30, ScheduleString: "Cả ngày", ClassTypeID: 2, Haspt: true},
		Class{Name: "Gói nâng cao 36 tháng", Price: 41900000, DurationDays: 1080, ScheduleString: "Cả ngày", ClassTypeID: 3, Haspt: true},
		Class{Name: "Gói nâng cao 18 tháng", Price: 25999000, DurationDays: 540, ScheduleString: "Cả ngày", ClassTypeID: 3, Haspt: true},
		Class{Name: "Gói cao cấp 60 tháng", Price: 150000000, DurationDays: 1800, ScheduleString: "Cả ngày", ClassTypeID: 4, Haspt: true},
		Class{Name: "Gói cao cấp 24 tháng", Price: 99900000, DurationDays: 720, ScheduleString: "Cả ngày", ClassTypeID: 4, Haspt: true},
	}
)

func initData(db *gorm.DB) {
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

	var classType ClassType
	db.Find(&classType).Count(&count)
	if count == 0 {
		for i := 0; i < len(classTypes); i++ {
			newClassType := classTypes[i]
			db.Create(&newClassType)
		}
	}

	var class Class
	db.Find(&class).Count(&count)
	if count == 0 {
		for i := 0; i < len(classes); i++ {
			newClass := classes[i]
			db.Create(&newClass)
		}
	}

	var dvs DeviceStatus
	db.Find(&dvs).Count(&count)
	if count == 0 {
		for i := 0; i < len(devicestatuses); i++ {
			n := devicestatuses[i]
			db.Create(&n)
		}
	}
}
