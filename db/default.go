package db

import (
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func DBConnect() *gorm.DB {
	var db *gorm.DB
	dsn := "root:123456@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DBConnect err: ", err)
	}
	return db
}