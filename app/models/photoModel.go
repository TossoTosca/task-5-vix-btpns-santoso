package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Caption   string         `gorm:"not null" json:"caption"`
	PhotoUrl  string         `gorm:"not null" json:"photo_url"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `gorm:"not null" json:"-"`
	User      User           `json:"-"`
}
