package test

import (
	"fmt"
	"os"

	"github.com/go-testfixtures/testfixtures/v3"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ImportFixture() {
	DBMS := "mysql"
	HOST := os.Getenv("DATABASE_HOST")
	PORT := os.Getenv("DATABASE_PORT")
	USER := os.Getenv("DATABASE_USER")
	PASS := os.Getenv("DATABASE_PASSWORD")
	DBNAME := os.Getenv("DATABASE_SCHEMA")

	PROTOCOL := "tcp(" + HOST + ":" + PORT + ")"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"

	db, err := sql.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	datapath := fmt.Sprintf("%s/src/github.com/yongwoon/echo-blog/test/fixtures", os.Getenv("GOPATH"))
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect(DBMS),
		testfixtures.Directory(datapath),
	)

	if err != nil {
		panic(err)
	}

	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}
