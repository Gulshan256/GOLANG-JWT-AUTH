package models

import "gorm.io/gorm"



type AuthUser struct {
	gorm.Model
	Id       int    `json:"ID" grom:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" grom:"unique"`
	Password string `json:"password"`
}