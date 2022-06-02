package models

import "time"

type User struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"varchar(50),unique"`
	Email     string    `json:"string" gorm:"varchar(50),unique"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	Photo     Photo     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comment   Comment   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
