package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string          `gorm:"type:varchar(200);UNIQUE"`
	Password string          `gorm:"type:varchar(200);"`
	FullName string          `gorm:"type:varchar(200);"`
	Info     json.RawMessage `gorm:"type:jsonb" json:"user_info"`
}

type UserInfo struct {
	Hobby    string `json:"hobby" validate:"required"`
	Category string `json:"category" validate:"required"`
}
