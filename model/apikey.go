package model

import (
	"api/database"
	"crypto/rand"
	"encoding/hex"

	"gorm.io/gorm"
)

type APIKey struct {
	gorm.Model
	Token  string `gorm:"size:255;not null;unique"`
	UserID uint
	User   User
}

func CreateAPIKey(User User) (*APIKey, error) {
	token := GenerateSecureToken(8)

	apiKey := &APIKey{
		Token:  token,
		UserID: User.ID,
	}

	err := database.Database.Create(&apiKey).Error
	if err != nil {
		return nil, err
	}
	return apiKey, nil
}

func (key *APIKey) Delete() error {
	err := database.Database.Unscoped().Delete(&key).Error
	if err != nil {
		return err
	}
	return nil
}

func FindAPIKey(token string) (*APIKey, error) {
	var apiKey APIKey
	err := database.Database.Preload("User").Where("token=?", token).First(&apiKey).Error
	if err != nil {
		return nil, err
	}
	return &apiKey, nil
}
func FindUserKey(userid any) (*APIKey, error) {

	var apiKey APIKey
	err := database.Database.Preload("User").Where("user_id=?", userid).First(&apiKey).Error
	if err != nil {
		return nil, err
	}
	return &apiKey, nil
}
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return "dnic" + hex.EncodeToString(b)
}
