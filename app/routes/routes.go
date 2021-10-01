package routes

import (
	"book_crud_ca/controllers/users"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(echo *echo.Echo) {
	users := echo.Group("users")
	users.POST("/register", cl.UserController.Insert)
}
