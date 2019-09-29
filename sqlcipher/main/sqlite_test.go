package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkSQLite(b *testing.B) {
	//b.ResetTimer()
	sqlite := new(SQLite)
	fileName := "users_sqlite.db"
	if err := sqlite.OpenFile("users_sqlite.db"); err != nil {
		b.Fatalf("SQLite open error:%s", err.Error())
	}
	defer os.Remove(fileName)
	defer sqlite.Close()

	c := "CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER PRIMARY KEY, `name` char, `password` chart, UNIQUE(`name`));"
	if _, err := sqlite.Exec(c); err != nil {
		b.Fatalf("create SQLite table error:%s", err.Error())
	}

	for i := 0; i < b.N; i++ {
		d := fmt.Sprintf("INSERT INTO `users` (name, password) values('xeodou%d', 123456);", i)
		_, err := sqlite.Exec(d)
		if err != nil {
			b.Fatalf("install SQLite row error:%s", err.Error())
		}
	}
}
