package randuser

import (
	"testing"
)

func TestRandUser(t *testing.T) {
	t.Run("More than 3,500 calls per minute", func(t *testing.T) {
		if _, err := GetRandUserInfo(true); err != nil {
			t.Fatalf("More than 3,500 calls per minute failed,error:%s", err.Error())
		}
	})

	t.Run("Less than 3,500 calls per minute", func(t *testing.T) {
		if _, err := GetRandUserInfo(false); err != nil {
			t.Fatalf("Less than 3,500 calls per minute failed,error:%s", err.Error())
		}
	})
}

//mysqldump --u root -p'123456' --databases province_warehouse_test > province_warehouse_test-2019-10-23.sql
