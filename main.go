package main

import (
	"gym-backend/db"
	"gym-backend/server"
)

func main() {
	// db := db.onnect()
	db := db.Connect()
	// a := dbGorm.Member{}
	// fmt.Print(a)
	server.StartRouter(db)
}
