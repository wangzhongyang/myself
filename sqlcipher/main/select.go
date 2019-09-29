package main

import (
	_ "github.com/xeodou/go-sqlcipher"
)

//func main() {
//	db, err := sql.Open("sqlite3", "./users.db?_key=123456")
//	if err != nil {
//		fmt.Println("open:", err)
//	}
//	defer db.Close()
//	fmt.Println("db:", db.Ping())
//
//	e := "select name, password from users where name='xeodou';"
//	rows, err := db.Query(e)
//	if err != nil {
//		fmt.Println("query:", err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var name string
//		var password string
//		rows.Scan(&name, &password)
//		fmt.Print("{\"name\":\"" + name + "\", \"password\": \"" + password + "\"}")
//	}
//}
