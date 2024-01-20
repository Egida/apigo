package db1

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db1 *gorm.DB

func Connect() {
	var err error

	dsn := viper.GetString("app.dbdb1")
	Db1, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to db1")
	}
}
