package user

import (
	"net/http"
	"strconv"
	"user-management/models"
	"user-management/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.IUserService
}

func NewUserHandler(userService *service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := (*h.userService).GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := (*h.userService).GetUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if user.Name == "" || user.Age == 0 {
		return c.JSON(http.StatusBadRequest, "Name and Age cannot be empty")
	}
	if err := (*h.userService).CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "User created successfully")
}

func (h *UserHandler) UpdateUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if user.Name == "" && user.Age == 0 {
		return c.JSON(http.StatusBadRequest, "Name or Age cannot be empty")
	}
	if err := (*h.userService).UpdateUserByID(uint(id), user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "User updated successfully")
}

func (h *UserHandler) DeleteByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := (*h.userService).DeleteByID(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "User deleted successfully")
}
