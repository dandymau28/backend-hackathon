package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    int    `json:"statusCode"`
	Error   string `json:"error"`
	Message string `json:"statusMessage"`
}

type Data struct {
	Code    int    `json:"statusCode"`
	Message string `json:"statusMessage"`
	Data    any    `json:"data"`
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
	return c.JSON(statusCode, Data{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func MessageResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Data{
		Code:    statusCode,
		Message: message,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}
