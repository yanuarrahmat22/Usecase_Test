package main

import (
	"usecase_test/db"
	"usecase_test/router"
)

func main() {
	db := db.Connect()
	router := router.SetupRouter(db)
	router.Run()
}
