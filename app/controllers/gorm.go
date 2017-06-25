package controllers

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"github.com/shiro16/golang-bbs/app/models"
	"log"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", dbInfoString())
	if err != nil {
		log.Panic("Failed to connect to database: %v\n", err)
	}

	db.DB()
	db.AutoMigrate(&models.Comment{}) //ここでTableの作成を行っている
	DB = db
}

func dbInfoString() string {
	s, b := revel.Config.String("db.info")
	if !b {
		log.Panic("database info not found")
	}

	return s
}