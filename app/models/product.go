package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func InitDB() {
	db, err := gorm.Open("mysql", "admin:admin@tcp(localhost:8889)/go_mysql")

	if err != nil {
		panic("faild to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "CANARY11", Price: 1000})
}

func FindAll() Product {
	var product Product
	db.First(&product, 1)
	return product
}
