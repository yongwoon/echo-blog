package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yongwoon/echo-blog/model"
)

var db *gorm.DB
var err error

// Init gorm db
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

	if os.Getenv("ENVIRONMENT") == "test" {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Post{})
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)

	return db
}

// DbManager get db connector
func DbManager() *gorm.DB {
	return db
}
