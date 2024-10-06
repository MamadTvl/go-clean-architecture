package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"size:255"`
	LastName  string `json:"last_name" gorm:"size:255"`
	UserName  string `json:"username" gorm:"uniqueIndex"`
	Password  string `json:"password" gorm:"notnull;size:255"`
}

func (*User) TableName() string {
	return "users"
}
