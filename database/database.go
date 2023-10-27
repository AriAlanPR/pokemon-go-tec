package database

import (
	"fmt"
	"gofiber/pokemon"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

//"https://github.com/AriAlanPR/pokemon-go-tec.git"

func Init() {
	var err error
	dsn := "root:1234@tcp(127.0.0.1:3306)/pokemon?charset=utf8mb4&parseTime=True&loc=Local"
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&pokemon.Pokemon{})
	fmt.Println("Database Migrated")
}

func Close() {
	if sqlDB, err := DBConn.DB(); err == nil {
		sqlDB.Close()
	}
}
