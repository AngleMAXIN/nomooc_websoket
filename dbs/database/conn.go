package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	sqlStr = `UPDATE contest_partner SET is_disabled='1' WHERE user_id = ? and contest_id = ?;`
)

var (
	dbConn *sql.DB
	err    error
	stmUp  *sql.Stmt
)

func init() {
	if dbConn, err = sql.Open("mysql", "root:maxinz@tcp(localhost:3306)/oj_database?charset=utf8"); err != nil {
		panic(err.Error())
	}

	if stmUp, err = dbConn.Prepare(sqlStr); err != nil {
		log.Printf("Prepare update SQl Failed: %s\n", err)
		return
	}

}
