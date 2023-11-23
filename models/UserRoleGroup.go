package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name"`
	Desc		string `json:"desc"`
}

type Permission struct {
	gorm.Model
	Name 		string `json:"name"`
	Desc		string `json:"desc"`
}

type RolePermission struct {
	RoleID       int `json:"role_id"`
	PermissionID int `json:"permission_id"`
}

type Group struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GroupRole struct {
	gorm.Model
	GroupID int `json:"group_id"`
	RoleID  int `json:"role_id"`
}

// type User struct {
// 	ID           int    `json:"id"`
// 	Username     string `json:"username"`
// 	PasswordHash string `json:"password_hash"`
// }
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Roles        []Role `gorm:"many2many:roles;joinForeignKey:ID"`
}

type UserRole struct {
	gorm.Model
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

type UserGroup struct {
	gorm.Model
	UserID  int `json:"user_id"`
	GroupID int `json:"group_id"`
}