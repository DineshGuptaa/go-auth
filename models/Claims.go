 // PATH: go-auth/models/Claims.go

 package models

 import "github.com/dgrijalva/jwt-go"

 type Claims struct {
     UserID       uint `json:"user_id"`
     Role string `json:"role"`
     Group string `json:"group"`
     Username string `json:"name"`
     jwt.StandardClaims
 }
