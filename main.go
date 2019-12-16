package main

import (
	"gym-backend/db"
	"gym-backend/server"
)

func main() {
	db := db.Connect()
	server.StartRouter(db)
}
