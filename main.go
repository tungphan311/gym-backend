package main

import (
	"gym-backend/dbGorm"
	"gym-backend/server"
)

func main() {
	// db := db.onnect()
	db := dbGorm.Connect()
	// a := dbGorm.Member{}
	// fmt.Print(a)
	server.StartRouter(db)
}
