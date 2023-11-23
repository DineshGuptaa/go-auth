package controllers

import (
	"fmt"
	"net/http"
	"time"

	"go-auth/models"

	"go-auth/conf"
	"go-auth/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte("my_secret_key")

func Login(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	var existingUser models.User

	models.DB.Preload("Roles").Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		return c.JSON(400, map[string]string{"error": "user does not exist"})
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		return c.JSON(400, map[string]string{"error": "invalid password"})
	}
	
	// If authentication is successful, generate a JWT token
    token := jwt.New(jwt.SigningMethodHS256)
    claims := &models.Claims{
		UserID:   existingUser.ID,
        Username: existingUser.Name,
        StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
        },
    }
    //token.Claims = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token.Claims = claims

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "could not generate token"})
	}
	// cookie := new(http.Cookie)
	// cookie.Name = "token"
	// cookie.Value = tokenString
	// cookie.Domain = "localhost"
	// cookie.Path = "/"
	now := time.Now()
	// Print the time in IST
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	// cookie.Expires = now.Add(48 * time.Hour)
	// c.SetCookie(cookie)
	return c.JSON(200, map[string]string{"success": "user logged in", "token":tokenString})
}

func Signup(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		return c.JSON(400, map[string]string{"error": "user already exists"})
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		return c.JSON(500, map[string]string{"error": "could not generate password hash"})
	}

	models.DB.Create(&user)

	return c.JSON(200, map[string]string{"success": "user created"})
}


func Home(c echo.Context) error {
	var iclaims interface{} = c.Get("user")
	// Convert interface to struct using type assertion
	claims, ok := iclaims.(*models.Claims)
	if !ok{
		return echo.NewHTTPError(http.StatusInternalServerError, conf.Props.MustGet("LogMSM"))
	}

	return c.JSON(200, map[string]string{"success": "home page", "role": claims.Role})
}

func Premium(c echo.Context) error {
	var iclaims interface{} = c.Get("user")
	// Convert interface to struct using type assertion
	claims, ok := iclaims.(*models.Claims)
	if !ok{
		return echo.NewHTTPError(http.StatusInternalServerError, conf.Props.MustGet("LogMSM"))
	}

	return c.JSON(200, map[string]string{"success": "premium page", "role": claims.Role})
}

func Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	c.SetCookie(cookie)
	return c.JSON(200, map[string]string{"success": "user logged out"})
}
