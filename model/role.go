package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name  string  `json:"name"`
	Users []*User `json:"users" gorm:"many2many:users_roles"`
}
