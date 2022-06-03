package models

import "time"

type Photo struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"varchar(254)"`
	Caption   string    `json:"caption" gorm:"varchar(254)"`
	PhotoUrl  string    `json:"photo_url" gorm:"text"`
	UserID    uint      `json:"user_id"`
	Comment   *Comment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:",omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
