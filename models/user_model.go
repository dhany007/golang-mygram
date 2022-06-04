package models

import (
	"final/helpers"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint       `json:"id,omitempty" gorm:"primaryKey"`
	Username  string     `json:"username" gorm:"type:varchar(50);unique"`
	Email     string     `json:"email" gorm:"type:varchar(50);unique"`
	Password  string     `json:"-"`
	Age       int        `json:"age,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPass(u.Password)
	err = nil

	return
}
