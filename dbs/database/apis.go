package database

import (
	"log"
)

// MarkUserStatus 用来更新用户的竞赛权限
func MarkUserStatus(uID, contestID int) error {
	var (
		err error
	)
	if _, err = stmUp.Exec(uID, contestID); err != nil {
		log.Printf("Exec update userid: %d and contest_id: %d Failed\n", uID, contestID)
		return err
	}
	log.Printf("Update user %d, contest %d\n", uID, contestID)
	return nil
}
