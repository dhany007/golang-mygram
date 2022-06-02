package models

import "time"

type User struct {
	ID        uint   `json:"id,omitempty" gorm:"primaryKey"`
	Username  string `json:"username"`
	Email     string `json:"string"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
	Photo     Photo
	Comment   Comment
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
