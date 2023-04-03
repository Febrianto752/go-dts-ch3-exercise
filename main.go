package main

import (
	"latihan/database"
	"latihan/routers"
)

func main() {
	database.StartDB()
	router := routers.StartApp()
	router.Run(":8080")
}
