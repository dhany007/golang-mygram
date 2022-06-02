package models

import "time"

type Comment struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
