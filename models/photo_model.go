package models

import "time"

type Photo struct {
	ID        uint   `json:"id,omitempty" gorm:"primaryKey"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserID    uint   `json:"user_id"`
	Comment   Comment
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
