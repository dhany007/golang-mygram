package models

import (
	"final/helpers"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"type:varchar(50);unique"`
	Email     string    `json:"string" gorm:"type:varchar(50);unique"`
	Password  string    `json:"-"`
	Age       int       `json:"age"`
	Photo     *Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:",omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("kesini dulu kan")
	u.Password = helpers.HashPass(u.Password)
	err = nil

	return
}
