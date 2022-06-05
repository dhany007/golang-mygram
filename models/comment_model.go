package models

import "time"

type Comment struct {
	ID        uint       `json:"id,omitempty" gorm:"primaryKey"`
	Message   string     `json:"message"`
	PhotoID   uint       `json:"photo_id"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	User      *User      `json:",omitempty"`
	Photo     *Photo     `json:",omitempty"`
}
