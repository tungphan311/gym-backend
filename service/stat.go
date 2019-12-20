package service

import (
	"fmt"
	dbGorm "gym-backend/db"
	"net/http"
	"sort"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type StatsReponse struct {
	totalMoney float64
	topclasses string
}

func GetStatsRecentMonth(c echo.Context, db *gorm.DB) error {
	var (
		allMembers []dbGorm.Member
		allBills   []dbGorm.Bill
		allDevices []dbGorm.Device
	)
	db.Where(&dbGorm.Device{Active: true}).Find(&allDevices)
	db.Where(&dbGorm.Bill{Active: true}).Find(&allBills)
	db.Where(&dbGorm.Member{Active: true}).Find(&allMembers)

	mapClassMoney := make(map[uint]float64)
	for _, b := range allBills {
		mapClassMoney[b.ClassID] += b.Amount
	}

	// return c.JSON(http.StatusOK, map[string]string{})
	return c.JSON(http.StatusOK, mapClassMoney)
}

type StatsDashboardResponse struct {
	TodayMember                   int
	MonthMember                   int
	TotalMember                   int
	TodayBillCount                int
	MonthBillsCount               int
	TodayMoney                    float32
	MonthMoney                    float32
	IncreasePercentMonthBillCount float32
	IncreasePercentMonthMoney     float32
	IncreaseMonthBillCount        float32
	IncreaseMonthMoney            float32
	TotalMoney                    float32
}

func GetStatsForDashboard(c echo.Context, db *gorm.DB) error {
	var (
		allMembers []dbGorm.Member
		allBills   []dbGorm.Bill
		allDevices []dbGorm.Device
		result     StatsDashboardResponse
	)
	db.Where(&dbGorm.Device{Active: true}).Find(&allDevices)
	db.Where(&dbGorm.Bill{Active: true}).Find(&allBills)
	db.Where(&dbGorm.Member{Active: true}).Find(&allMembers)

	// Get today bills
	var (
		todayBills []dbGorm.Bill
		todayMoney float64 = 0
	)
	for _, b := range allBills {
		if b.CreatedTime == time.Now() {
			_ = append(todayBills, b)
			todayMoney += b.Amount
		}
	}

	var todayMembers []dbGorm.Member
	for _, b := range allMembers {
		if b.CreatedAt == time.Now() {
			_ = append(todayMembers, b)
		}
	}

	// Get this month bills
	var (
		monthBills []dbGorm.Bill
		monthMoney float64 = 0
	)
	for _, b := range allBills {
		if b.CreatedTime.Month() == time.Now().Month() &&
			b.CreatedTime.Year() == time.Now().Year() {
			_ = append(monthBills, b)
			monthMoney += b.Amount
		}
	}

	// Get this month members
	var monthMembers []dbGorm.Member
	for _, b := range allMembers {
		if b.CreatedAt.Month() == time.Now().Month() &&
			b.CreatedAt.Year() == time.Now().Year() {
			_ = append(monthMembers, b)
		}
	}

	// Get last month bills
	var (
		lastMonthBills []dbGorm.Bill
		lastMonthMoney float64 = 0
	)

	lastMonth := 12
	if int(time.Now().Month()) != 1 {
		lastMonth = int(time.Now().Month()) - 1
	}
	lastYear := int(time.Now().Year())
	if lastMonth == 12 {
		lastYear -= 1
	}
	for _, b := range allBills {
		if lastMonth == int(b.CreatedAt.Month()) &&
			lastYear == int(b.CreatedAt.Year()) {
			_ = append(lastMonthBills, b)
			lastMonthMoney += b.Amount
		}
	}

	var lastMembers []dbGorm.Member
	for _, b := range allMembers {
		if lastMonth == int(b.CreatedAt.Month()) &&
			lastYear == int(b.CreatedAt.Year()) {
			_ = append(lastMembers, b)
		}
	}

	// Total revenue
	var (
		totalMoney float64 = 0
	)
	for _, b := range allBills {
		totalMoney += b.Amount
	}

	// Result
	result.TodayBillCount = len(todayBills)
	result.MonthBillsCount = len(monthBills)
	result.MonthMoney = float32(monthMoney)
	result.TodayMoney = float32(todayMoney)
	result.TotalMoney = float32(totalMoney)
	increasedMoney := float32(monthMoney - lastMonthMoney)
	increasedBills := float32(len(monthBills) - len(lastMonthBills))
	result.TodayMember = len(todayMembers)
	result.MonthMember = len(monthMembers)
	result.TotalMember = len(allMembers)

	if increasedBills >= 0 {
		result.IncreaseMonthBillCount = increasedBills
	}

	if increasedMoney >= 0 {
		result.IncreaseMonthMoney = increasedMoney
	}

	if len(lastMonthBills) != 0 {
		result.IncreasePercentMonthBillCount = (float32(len(monthBills) - len(lastMonthBills))) / float32(len(lastMonthBills))
	}

	if lastMonthMoney != 0 {
		result.IncreasePercentMonthMoney = float32(((monthMoney - lastMonthMoney) / lastMonthMoney))
	}

	return c.JSON(http.StatusOK, result)
}

// Get Top Money Classes
type TopClassTotalMoney struct {
	Class      dbGorm.Class
	TotalMoney float64
}

func GetTopMoneyClasses(c echo.Context, db *gorm.DB) error {
	var (
		allClasses []dbGorm.Class
		allBills   []dbGorm.Bill
		topClasses []TopClassTotalMoney
		results    [5]TopClassTotalMoney
		topClass   TopClassTotalMoney
	)

	db.Where(&dbGorm.Bill{Active: true}).Find(&allBills)
	db.Where(&dbGorm.Class{Active: true}).Find(&allClasses)

	topClasses = make([]TopClassTotalMoney, 0)

	mapClassMoney := make(map[uint]float64)
	for _, b := range allBills {
		mapClassMoney[b.ClassID] += b.Amount
	}

	for k, v := range mapClassMoney {
		class := dbGorm.Class{}
		db.Where("id = ?", k).First(&class)

		fmt.Println("id = ", k)
		fmt.Println(class)
		fmt.Println("----")

		if class.ID != 0 {
			topClass.Class = class
			topClass.TotalMoney = v
			topClasses = append(topClasses, topClass)
		}
	}

	sort.Slice(topClasses, func(i, j int) bool {
		return topClasses[i].TotalMoney > topClasses[j].TotalMoney
	})

	for index, b := range topClasses {
		if index >= 5 {
			break
		}
		results[index] = b
	}

	//return c.JSON(http.StatusOK, results)
	return c.JSON(http.StatusOK, results)
}
