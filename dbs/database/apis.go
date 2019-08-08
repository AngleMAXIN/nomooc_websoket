package database

import (
	"log"
	"time"
)

// MarkUserStatus 用来更新用户的竞赛权限
func MarkUserStatus(uID, contestID int) (err error) {

	if _, err = stmUpdateUserStatus.Exec(uID, contestID); err != nil {
		log.Printf("Exec update userid: %d and contest_id: %d Failed\n", uID, contestID)
		return
	}
	return
}

// GetcontestEndTime 获取次竞赛的关闭时间
func GetcontestEndTime(contestID int) (endTime time.Time, err error) {
	// var rows sql.Result
	if err = stmSelectEndTiem.QueryRow(contestID).Scan(&endTime); err != nil {
		log.Printf("Exec select contest_id: %d Failed\n", contestID)
		return
	}
	return
}
