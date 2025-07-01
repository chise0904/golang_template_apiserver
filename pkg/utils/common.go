package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommonResponse struct {
	Status string `json:"status"`
	// Error  string      `json:"error,omitempty"`
	// Detail interface{} `json:"detail,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func MakeResponse(ctx echo.Context, statusCode int, data interface{}) error {

	return ctx.JSON(statusCode, &CommonResponse{
		Status: http.StatusText(statusCode),
		Data:   data,
	})
}
