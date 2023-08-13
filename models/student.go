package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	Document string `json:"document"`
	UserName string `json:"username"`
}
