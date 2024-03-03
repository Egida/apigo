package model

import (
	"api/db1"

	"time"
)

type Ip_addresses struct {
	ID         string `gorm:"primaryKey"`
	Subnet_id  int    `gorm:"not null"`
	Ip         string `gorm:"size:45;not null"`
	Customer   string `gorm:"size:40"`
	Status     string `gorm:"size:191;not null"`
	Type       string `gorm:"size:20;not null"`
	Zone       string `gorm:"size:30;not null"`
	Zonename   string `gorm:"size:40;not null"`
	Created_at time.Time
	Updated_at time.Time
}

func FindByip(ip string) (Ip_addresses, error) {
	var ipadress Ip_addresses
	err := db1.Db1.Where("ip=?", ip).Find(&ipadress).Error
	if err != nil {
		return ipadress, err
	}
	return ipadress, nil
}
func FindZoneby(id string) (Ip_addresses, error) {
	var ipadress Ip_addresses
	err := db1.Db1.Where("id=?", id).Find(&ipadress).Error
	if err != nil {
		return ipadress, err
	}

	return ipadress, nil
}
func userip(ip string, email string) (Ip_addresses, error) {
	var ipadress Ip_addresses
	err := db1.Db1.Where("ip=?", ip).Where("email=?", email).Find(&ipadress).Error
	if err != nil {
		return ipadress, err
	}
	return ipadress, nil
}
func ListIps() (Ip_addresses, error) {
	var ipadress Ip_addresses
	err := db1.Db1.First(&ipadress).Error
	if err != nil {
		return ipadress, err
	}
	return ipadress, nil
}
