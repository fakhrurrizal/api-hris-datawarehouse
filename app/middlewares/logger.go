package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger Middleware
func Logger() echo.MiddlewareFunc {
	responses, err := os.Create("public/logs.txt")
	if err != nil {
		responses = os.Stdout
	}

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] - Host: ${host}, Referer: ${referer}, Method: ${method}, Url: ${uri}, Status: ${status}, Latency: ${latency_human} \n",
		Output: responses,
	})
}
