package model

import (
	"api/db1"
	"time"
)

type Accessip struct {
	ID         int    `gorm:"primaryKey"`
	Ip         string `gorm:"size:20"`
	Customer   string `gorm:"size:50"`
	Created_at time.Time
	Updated_at time.Time
}

func FindAccess(ip string) (Accessip, error) {
	var accessip Accessip
	err := db1.Db1.Where("ip=?", ip).Find(&accessip).Error
	if err != nil {
		return accessip, err
	}
	return accessip, nil
}
