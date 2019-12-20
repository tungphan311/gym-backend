package service

import (
	dbGorm "gym-backend/db"
	"net/http"
	"sort"

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
		class      dbGorm.Class
	)

	db.Where(&dbGorm.Bill{Active: true}).Find(&allBills)
	db.Where(&dbGorm.Class{Active: true}).Find(&allClasses)

	topClasses = make([]TopClassTotalMoney, 0)

	mapClassMoney := make(map[uint]float64)
	for _, b := range allBills {
		mapClassMoney[b.ClassID] += b.Amount
	}

	for k, v := range mapClassMoney {
		db.Where("id = ?", k).First(&class)

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
