package database

import "database/sql"

const (
	updateSQLStr = `UPDATE contest_partner SET is_disabled='1' WHERE user_id = ? and contest_id = ?;`
	selectSQLStr = `SELECT end_time FROM contest WHERE id = ?;`
)

var (
	dbConn                                *sql.DB
	stmUpdateUserStatus, stmSelectEndTiem *sql.Stmt
)
