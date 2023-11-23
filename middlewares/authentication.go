package middlewares

import (
	"go-auth/conf"
	"go-auth/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"

	//"go-auth/conf/properties.json"

	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties"
)


func init(){
    log.SetOutput(os.Stdout)
    log.SetFlags(log.Llongfile | log.LstdFlags)
    dir, err := filepath.Abs(filepath.Dir("conf"))
    if err != nil {
            log.Fatal(err)
    }
    conf.Props = properties.MustLoadFile(dir+"/conf/messages.properties", properties.UTF8)
}
func IsAuthorized(next echo.HandlerFunc) echo.HandlerFunc{
    return func(c echo.Context) error {
        skip := []string{"/login", "/signup"}

        // Skip authentication for specific routes
		if utils.Contains(skip, c.Path()) {
			return next(c)
		}

        cookie, err := c.Cookie("token")

        if err != nil {
            //c.Abort()
            return c.JSON(401, map[string]string{"error": "unauthorized"})
        }

        claims, err := utils.ParseToken(cookie.Value)

        if err != nil {
            
            //c.Abort()
            return c.JSON(401, map[string]string{"error": "unauthorized"})
        }

        c.Set("role", claims.Role)
        return next(c)
    }
}

// Middleware function for user authentication.
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
        // Log a custom message using Echo's logger.

	    log.Printf(conf.Props.MustGet("LogMSM"), c.Path())
        skip := []string{"/login", "/signup"}

        // Skip authentication for specific routes
		if utils.Contains(skip, c.Path()) {
            log.Printf(conf.Props.MustGet("LogMAS"), c.Path())
			return next(c)
		}

		// Extract the JWT token from the request header.
		authHeader := c.Request().Header.Get("X-Auth-ID")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, conf.Props.MustGet("ErrMAH"))
		}

		tokenString := authHeader

		// Validate and parse the JWT token.
		claims, err := utils.ParseToken(tokenString)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, conf.Props.MustGet("ErrIT"))
		}

		if claims.Role != "user" && claims.Role != "admin" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": conf.Props.MustGet("ErrUA")})
        }
		
		// Set the user in the context for the next middleware or handler.
		c.Set("user", claims)

		return next(c)
	}
}