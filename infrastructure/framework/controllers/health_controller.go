package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/laironacosta/ms-echo-go-layout/domain/users/dto"
	"net/http"
)

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.Response{
		Message: "ok",
	})
}
