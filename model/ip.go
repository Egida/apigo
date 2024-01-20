package model

import (
	"gorm.io/gorm"

	"api/database"
)

type IP struct {
	gorm.Model
	Address string `gorm:"size:255;not null"`
	UserID  uint
	User    User
}

func CreateIP(user User, address string) (*IP, error) {
	ip := &IP{
		Address: address,
		UserID:  user.ID,
	}
	err := database.Database.Create(&ip).Error
	if err != nil {
		return nil, err
	}
	return ip, nil
}

func IPAllowed(userID uint, address string) bool {
	var ip IP
	err := database.Database.
		Where("user_id=?", userID).
		Where("address=?", address).
		First(&ip).
		Error
	return nil == err
}

func (ip *IP) Delete() error {
	return database.Database.
		Unscoped().
		Delete(&ip).
		Error
}
