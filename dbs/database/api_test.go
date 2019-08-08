package database

import (
	"testing"
)

func TestMarkUserStatus(t *testing.T) {
	uID, contestID := 1, 51
	r := MarkUserStatus(uID, contestID)
	if r != nil {
		t.Errorf("update failed %s\n", r)
	}
}

// func TestAgainMarkUserStatus(t *testing.T) {
// 	uID, contestID := 1, 51
// 	r := MarkUserStatus(uID, contestID)
// 	if r != nil {

// 		t.Errorf("update failed:%s", r)
// 	}
// 	defer stmUp.Close()
// }

func TestGetcontestEndTime(t *testing.T) {
	contestID := 34
	for index := 0; index < 100; index++ {
		err, time := GetcontestEndTime(contestID)
		if err != nil {
			t.Errorf("failed %s", err.Error())
		}
		t.Logf("%v", time)
	}

}
