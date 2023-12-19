package response

import (
	"github.com/labstack/echo/v4"
)

func ListRes(c echo.Context, code int, text interface{}) error {
	return c.JSON(
		code,
		struct {
			Text interface{} `json:"text"`
		}{
			Text: text,
		},
	)
}
