package model

import (
	
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"fmt"
  )

  var DB *gorm.DB

  func Connect(){
	dsn := "root:@tcp(127.0.0.1:3306)/fiber-orm?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	fmt.Println("connected database ",db)

	db.AutoMigrate(&RisetPenelitian{})
	db.AutoMigrate(&Peneliti{})
	db.AutoMigrate(&Usertype{})
	db.AutoMigrate(&Users{})
	
	DB = db
  }