package users

import (
	"book_crud_ca/bussinesses/users"
	"book_crud_ca/controllers"
	"book_crud_ca/controllers/users/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userUseCase users.UseCase
}

func NewUserController(uc users.UseCase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (ctrl *UserController) Insert(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}

	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if err := ctrl.userUseCase.Register(ctx, req.ToDomain()); err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusOK, "Successfully inserted")
}
