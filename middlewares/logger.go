package middlewares

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)
func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Log incoming request
        log.Printf("Incoming Request: %s %s", c.Request().Method, c.Request().URL.String())

        // Call next handler
        if err := next(c); err != nil {
            c.Error(err)
        }

        // Log outgoing response
        log.Printf("Outgoing Response: %d %s", c.Response().Status, http.StatusText(c.Response().Status))

        return nil
    }
}



