package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "jaqen:123pass/Yogurt?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d

}

func GetDb() *gorm.DB {
	return db
}
