package service

import (
	"net/http"
	"strconv"
	"time"

	dbGorm "gym-backend/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type BillRequest struct {
	ID          int     `json:"id"`
	MemberID    uint    `json:"memberid"`
	StaffID     uint    `json:"staffid"`
	ClassID     uint    `json:"classid"`
	Amount      float64 `json:"amount"`
	CreatedTime string  `json:"createdtime"`
}

type BuyClassRequest struct {
	ClassID  uint `json:"classid"`
	MemberID uint `json:"memberid"`
	StaffID  uint `json:"staffid"`
}

type BillResponse struct {
	ID          uint
	MemberID    uint
	StaffID     uint
	ClassID     uint
	Amount      float64
	CreatedTime string
	MemberName  string
	StaffName   string
	ClassName   string
}

func CreateBill(c echo.Context, db *gorm.DB) error {
	r := new(BillRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	format := "02/01/2006"
	createdTime, _ := time.Parse(format, r.CreatedTime)

	n := dbGorm.Bill{
		Amount:      r.Amount,
		CreatedTime: createdTime,
		MemberID:    r.MemberID,
		StaffID:     r.StaffID,
		ClassID:     r.ClassID,
	}
	db.Create(&n)
	return c.JSON(http.StatusCreated, "Thêm hoá đơn mới thành công")
}

func GetBillWithId(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	q := dbGorm.Bill{}
	db.Where("id = ?", id).First(&q)
	if q.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy dữ liệu.",
		})
	}

	return c.JSON(http.StatusOK, q)
}

func GetAllBill(c echo.Context, db *gorm.DB) error {
	a := []dbGorm.Bill{}

	classid, _ := strconv.ParseUint(c.QueryParam("classid"), 10, 64)
	memberid, _ := strconv.ParseUint(c.QueryParam("memberid"), 10, 64)
	staffid, _ := strconv.ParseUint(c.QueryParam("staffid"), 10, 64)
	db.Where(&dbGorm.Bill{
		Active:   true,
		ClassID:  uint(classid),
		MemberID: uint(memberid),
		StaffID:  uint(staffid),
	}).Find(&a)

	var br BillResponse
	var brs []BillResponse
	for _, v := range a {
		br.ID = v.ID
		br.MemberID = v.MemberID
		br.StaffID = v.StaffID
		br.ClassID = v.ClassID
		br.Amount = v.Amount
		br.CreatedTime = v.CreatedTime.String()

		member := &dbGorm.Member{}
		class := &dbGorm.Class{}
		staff := &dbGorm.Staff{}
		db.Where("id = ?", v.MemberID).First(&member)
		db.Where("id = ?", v.StaffID).First(&staff)
		db.Where("id = ?", v.ClassID).First(&class)

		br.MemberName = member.FullName
		br.StaffName = staff.FullName
		br.ClassName = class.Name

		brs = append(brs, br)
	}

	return c.JSON(http.StatusOK, brs)
}

func UpdateBill(c echo.Context, db *gorm.DB) error {
	r := new(BillRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	format := "02/01/2006"
	createdTime, _ := time.Parse(format, r.CreatedTime)

	var q dbGorm.Bill
	db.Where("id = ?", r.ID).First(&q)
	q.Amount = r.Amount
	q.CreatedTime = createdTime
	q.StaffID = r.StaffID
	q.MemberID = r.MemberID
	q.ClassID = r.ClassID
	db.Save(&q)

	return c.JSON(http.StatusOK, "OK")
}

func DeactiveBill(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	n := dbGorm.Bill{}
	db.Where("id = ?", id).First(&n)
	if n.ID == 0 {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Không tìm thấy hoá đơn.",
		})
	}

	n.Active = false
	db.Save(&n)

	return c.JSON(http.StatusOK, "Ok")
}

func BuyClass(c echo.Context, db *gorm.DB) error {
	r := new(BuyClassRequest)

	if err := c.Bind(r); err != nil {
		return err
	}

	newBill := dbGorm.Bill{}
	class := dbGorm.Class{}
	db.Where("id = ?", r.ClassID).First(&class)
	newBill.MemberID = r.MemberID
	newBill.StaffID = r.StaffID
	newBill.Amount = class.Price
	newBill.ClassID = r.ClassID
	newBill.CreatedTime = time.Now()
	db.Create(&newBill)

	return c.JSON(http.StatusOK, newBill)
}
