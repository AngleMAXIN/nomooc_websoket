package database

import (
	prcf "Project/websocket/config"
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	var (
		err error
	)
	cf := prcf.ProConfig.GetMysqlConfig()

	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=%s&parseTime=true",
		cf.User, cf.Passwd, cf.Host, cf.Port, cf.DB, url.QueryEscape(cf.Loc))

	if dbConn, err = sql.Open("mysql", uri); err != nil {
		panic(err.Error())
	}

	if stmUpdateUserStatus, err = dbConn.Prepare(updateSQLStr); err != nil {
		log.Printf("Prepare update SQl Failed: %s\n", err)
		return
	}

	if stmSelectEndTiem, err = dbConn.Prepare(selectSQLStr); err != nil {
		log.Printf("Prepare select SQl Failed: %s\n", err)
		return
	}

}
