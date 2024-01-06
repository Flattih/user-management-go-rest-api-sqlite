package user

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, userHandler *UserHandler) {
	e.GET("/users", userHandler.GetAllUsers)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.POST("/users", userHandler.CreateUser)
	e.PUT("/users/:id", userHandler.UpdateUserByID)
	e.DELETE("/users/:id", userHandler.DeleteByID)
}
