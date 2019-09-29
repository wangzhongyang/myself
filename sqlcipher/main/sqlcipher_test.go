package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func BenchmarkSQLCipher(b *testing.B) {
	fileName := "users_sqlcipher.db"
	cipher := new(SQLCipher)
	if err := cipher.OpenFile(fileName, "_key=123456"); err != nil {
		b.Fatalf("SQLCipher open error:%s", err.Error())
	}
	defer os.Remove(fileName)
	defer cipher.Close()

	c := "CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER PRIMARY KEY, `name` char, `password` chart, `created_at` datetime, `updated_at` datetime,`deleted_at` datetime, UNIQUE(`name`));"
	_, err := cipher.Exec(c)
	if err != nil {
		b.Fatalf("create table error:%s", err.Error())
	}
	for i := 0; i < b.N; i++ {
		d := fmt.Sprintf("INSERT INTO `users` (name, password) values('xeodou%d', 123456);", i)
		_, err = cipher.Exec(d)
		if err != nil {
			b.Fatalf("install row error:%s", err.Error())
		}
	}
}

func BenchmarkSQLCipherOrm(b *testing.B) {
	fileName := "users_orm.db"
	orm := new(SQLCipherOrm)
	if err := orm.OpenFile(fileName, "_key=123456"); err != nil {
		b.Fatalf("SQLCipher open error:%s", err.Error())
	}
	defer os.Remove(fileName)
	defer orm.Close()

	c := "CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER PRIMARY KEY, `name` char, `password` chart, `created_at` datetime, `updated_at` datetime,`deleted_at` datetime, UNIQUE(`name`));"
	if err := orm.Exec(c).Error; err != nil {
		b.Fatalf("create table error:%s", err.Error())
	}
	for i := 0; i < b.N; i++ {
		if err := orm.Create(&User{
			Name:     fmt.Sprintf("xeodou%d", i),
			Password: "123456",
			ID:       i,
		}).Error; err != nil {
			b.Fatalf("install row error:%s", err.Error())
		}
	}
}

func BenchmarkSQLCipherOrmOptimize(b *testing.B) {
	fileName := "users_orm.db"
	orm := new(SQLCipherOrm)
	if err := orm.OpenFile(fileName, "_key=123456"); err != nil {
		b.Fatalf("SQLCipher open error:%s", err.Error())
	}
	defer os.Remove(fileName)
	defer orm.Close()

	c := "CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER PRIMARY KEY, `name` char, `password` chart, `created_at` datetime, `updated_at` datetime,`deleted_at` datetime, UNIQUE(`name`));"
	if err := orm.Exec(c).Error; err != nil {
		b.Fatalf("create table error:%s", err.Error())
	}
	orm.DB = orm.Exec("PRAGMA synchronous = OFF").Begin()
	isCommit := true
	for i := 0; i < b.N; i++ {
		if err := orm.Create(&User{
			Name:     fmt.Sprintf("xeodou%d", i),
			Password: "123456",
			ID:       i,
		}).Error; err != nil {
			b.Fatalf("install row error:%s", err.Error())
			isCommit = false
			break
		}
	}
	t := time.Now()
	if isCommit {
		orm.Commit()
	} else {
		orm.Rollback()
	}
	fmt.Println("commit time:", time.Since(t))
}
