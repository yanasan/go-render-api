package controller

import (
	"net/http"

	"github.com/Poul-george/go-render-ap/go-render-api/response"
	"github.com/labstack/echo/v4"
)

func List(ctx echo.Context) error {
	return response.ListRes(ctx, http.StatusOK, "text")
}
