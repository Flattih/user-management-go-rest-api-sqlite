package main

import (
	"user-management/config"
	"user-management/database"
	"user-management/router"

	"github.com/labstack/echo/v4"
)

func main() {
	println("Hello, World!")

	db := database.ConnectDB(config.ProdDB)
	e := echo.New()
	router.SetupRoutes(e, db)
	e.Start(":8080")

}
