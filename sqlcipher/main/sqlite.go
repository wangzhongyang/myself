package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	*sql.DB
}

func (s *SQLite) OpenFile(fileName string) error {
	err := errors.New("")
	s.DB, err = sql.Open("sqlite3", fileName)
	return err
}

func (s *SQLite) Exec(query string) (sql.Result, error) {
	return s.DB.Exec(query)
}

func main() {
	_ = os.Remove("user.db")
	defer os.Remove("user.db")
	sqlite := new(SQLite)
	if err := sqlite.OpenFile("user.db"); err != nil {
		fmt.Println("open file error:", err.Error())
		return
	}
	defer sqlite.Close()
	c := "CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER PRIMARY KEY, `name` char, `password` chart, UNIQUE(`name`));"
	if _, err := sqlite.Exec(c); err != nil {
		fmt.Printf("create SQLite table error:%s\n", err.Error())
		return
	}
	d := fmt.Sprintf("INSERT INTO `users` (name, password) values('xeodou%d', 123456);", 1)
	if _, err := sqlite.Exec(d); err != nil {
		fmt.Printf("install SQLite row error:%s\n", err.Error())
		return
	}
	e := "select name, password from users;"
	rows, err := sqlite.Query(e)
	if err != nil {
		fmt.Println("query:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var password string
		rows.Scan(&name, &password)
		fmt.Print("{\"name\":\"" + name + "\", \"password\": \"" + password + "\"}")
	}
}
