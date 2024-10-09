package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IdCard struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	Id   string `gorm:"column:id" json:"id"`
}

func (table *IdCard) TableName() string {
	return "id_card"
}

var DB = Init()

func Init() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/gin_gorm_oj?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&IdCard{})
	return db
}
