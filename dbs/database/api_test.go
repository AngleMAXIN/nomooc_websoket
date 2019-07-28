package database

import "testing"

func TestMarkUserStatus(t *testing.T) {
	uID, contestID := 1, 51
	r := MarkUserStatus(uID, contestID)
	if r != nil {
		t.Error("update failed")
	}
}

func TestAgainMarkUserStatus(t *testing.T) {
	uID, contestID := 1, 51
	r := MarkUserStatus(uID, contestID)
	if r != nil {

		t.Errorf("update failed:%s", r)
	}
	defer stmUp.Close()
}
