package main

import (
	"Salonisaroha/api"
	db "Salonisaroha/database"
)

func main() {
	db := db.NewDBInstance("root:1020547676@tcp(127.0.0.1:3306)", "expense_tracker")
	db.Connect()

	api.StartServer()
}
