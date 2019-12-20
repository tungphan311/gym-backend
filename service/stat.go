package service

import (
	dbGorm "gym-backend/db"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetStatsRecentMonth(c echo.Context, db *gorm.DB) error {
	allDevices := []dbGorm.Device{}
	allBills := []dbGorm.Bill{}
	allMembers := []dbGorm.Member{}

	db.Where(&dbGorm.Device{Active: true}).Find(&allDevices)
	db.Where(&dbGorm.Bill{Active: true}).Find(&allBills)
	db.Where(&dbGorm.Member{Active: true}).Find(&allMembers)

	allBills[0].CreatedAt.Month()

	return c.JSON(http.StatusOK, map[string]string{})
}
