package app

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"not null"`
	Email    string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Photos   []Photo
}

type Photo struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Title    string `gorm:"not null"`
	Caption  string
	PhotoUrl string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	User     User
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
