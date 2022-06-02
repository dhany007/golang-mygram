package models

import "time"

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
