package models

import "time"

type SocialMedia struct {
	ID             uint      `json:"id,omitempty" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"not null,varchar(50)"`
	SocialMediaUrl string    `json:"social_media_url" gorm:"text"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
