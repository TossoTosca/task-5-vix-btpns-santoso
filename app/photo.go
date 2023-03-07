package app

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Photo struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Title    string `gorm:"not null"`
	Caption  string
	PhotoUrl string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	User     User
}

func (p *Photo) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (p *Photo) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
package app

import "time"

type Photo struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	URL       string    `gorm:"not null" json:"url"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	UserID    uint      `gorm:"not null" json:"-"`
}

// CreatePhoto creates a new photo
func (a *App) CreatePhoto(photo *Photo) error {
	err := a.DB.Create(photo).Error
	if err != nil {
		return err
	}
	return nil
}

// GetPhotosByUserID gets all photos by user ID
func (a *App) GetPhotosByUserID(userID uint) ([]Photo, error) {
	var photos []Photo
	err := a.DB.Where("user_id = ?", userID).Find(&photos).Error
	if err != nil {
		return nil, err
	}
	return photos, nil
}

// GetPhotoByID gets a photo by ID
func (a *App) GetPhotoByID(id uint) (*Photo, error) {
	var photo Photo
	err := a.DB.First(&photo, id).Error
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

// UpdatePhoto updates a photo
func (a *App) UpdatePhoto(photo *Photo) error {
	err := a.DB.Save(photo).Error
	if err != nil {
		return err
	}
	return nil
}

// DeletePhoto deletes a photo
func (a *App) DeletePhoto(photo *Photo) error {
	err := a.DB.Delete(photo).Error
	if err != nil {
		return err
	}
	return nil
}
