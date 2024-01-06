package router

import (
	"user-management/api/user"
	"user-management/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	// User routes
	userService := service.NewUserService(db)
	userHandler := user.NewUserHandler(&userService)
	user.RegisterRoutes(e, userHandler)
	// Other routes

}
