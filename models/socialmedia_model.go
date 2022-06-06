package models

import "time"

type SocialMedia struct {
	ID             uint       `json:"id,omitempty" gorm:"primaryKey"`
	Name           string     `json:"name" gorm:"not null,varchar(50)"`
	SocialMediaUrl string     `json:"social_media_url" gorm:"text"`
	UserID         uint       `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	User           *User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:",omitempty"`
}
