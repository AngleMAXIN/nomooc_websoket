package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	if dbConn, err = sql.Open("mysql", "root:maxinz@tcp(localhost:3306)/oj_database?charset=utf8"); err != nil {
		panic(err.Error())
	}

}
