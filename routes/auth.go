// PATH: go-auth/routes/auth.go

package routes

import (
    "go-auth/controllers"

    "github.com/labstack/echo/v4"
)

func AuthRoutes(r *echo.Echo) {
    r.POST("/login", controllers.Login)
    r.POST("/signup", controllers.Signup)
    r.GET("/home", controllers.Home)
    r.GET("/premium", controllers.Premium)
    r.GET("/logout", controllers.Logout)
}
