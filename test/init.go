package test

import (
	"os"

	"github.com/yongwoon/echo-blog/config/initializer"
	"github.com/yongwoon/echo-blog/db"
)

// Setup configuration before test
func Setup() {
	initializer.InitDotenv()

	db.Init()
}

// Teardown clear after test
func Teardown() {
	DBNAME := os.Getenv("DATABASE_SCHEMA")
	dbM := db.DbManager()

	rows, err := dbM.Raw("show tables").Rows()
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			panic(err.Error())
		}

		dbM.Exec("TRUNCATE TABLE " + DBNAME + "." + table)
	}
}
