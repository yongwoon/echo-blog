package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() *gorm.DB {
	DBMS := "mysql"
	HOST := os.Getenv("DATABASE_HOST")
	PORT := os.Getenv("DATABASE_PORT")
	USER := os.Getenv("DATABASE_USER")
	PASS := os.Getenv("DATABASE_PASSWORD")
	DBNAME := os.Getenv("DATABASE_SCHEMA")

	PROTOCOL := "tcp(" + HOST + ":" + PORT + ")"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"

	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)
	// db.AutoMigrate(&model.Post{})
	return db
}

func DbManager() *gorm.DB {
	return db
}
