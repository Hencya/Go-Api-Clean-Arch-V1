package controller

import (
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"rc"`
		Message  string   `json:"message"`
		Messages []string `json:"messages"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewErrorResponse(ctx echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Something gone wrong"
	response.Meta.Messages = []string{err.Error()}

	return ctx.JSON(status, response)
}

func NewSuccessResponse(ctx echo.Context, status int, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Success"
	response.Data = param

	return ctx.JSON(status, response)

}
