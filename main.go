// PATH: go-auth/main.go

package main

import (
	"go-auth/middlewares"
	"go-auth/models"
	"go-auth/routes"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

    // Set the time zone
    time.Local = time.FixedZone("IST", 3600*5.5)
    
    //Echo used
    r := echo.New()

    // Logger configuration
	loggerConfig := middleware.LoggerConfig{
		// Output can be a file or os.Stdout by default
		Output:           os.Stdout,		
		// Customize the message format
		Format:           "${time_rfc3339} ${method} ${uri} ${status} ${error}\n",
	}
   

    // Middleware
    r.Use(middleware.LoggerWithConfig(loggerConfig))
    r.Use(middleware.Recover())
    r.Use(middlewares.AuthMiddleware)
    // Load .env file and Create a new connection to the database
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    config := models.Config{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
    }

    // Initialize DB
    models.InitDB(config)

    // Load the routes
    routes.AuthRoutes(r)

    // Run the server
    r.Start(":8080")
}
