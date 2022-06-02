package main

import (
	"final/configs"
	"final/routers"
)

func main() {
	PORT := ":3000"
	db := configs.StartDB()

	router := routers.StartEngine(db)

	router.Run(PORT)
}
