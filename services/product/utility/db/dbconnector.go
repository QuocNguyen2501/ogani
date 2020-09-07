package db

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func ConnectDbHandler() *gorm.DB{
	db, err := gorm.Open(cast.ToString(viper.Get("dbDialect")), viper.Get(cast.ToString("connStr")))
	if err != nil{
		panic(err)
	}

	return db
}