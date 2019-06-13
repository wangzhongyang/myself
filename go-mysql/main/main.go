package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "gateway:pass4u@tcp(db-master.stg.tw.bindo.in:3306)/bindo?parseTime=true")
	if err != nil {
		fmt.Println("this is err:		", err.Error())
	} else {
		fmt.Println("conn success")
	}
	defer db.Close()
	type Res struct {
		Total int
	}
	var res Res
	sql := "SELECT COUNT(*) as total FROM (SELECT DISTINCT categories.* FROM `categories` LEFT JOIN products ON products.category_id = categories.id LEFT JOIN listings ON listings.product_id = products.id WHERE listings.store_id = '4934') as T1"
	if err := db.Model(nil).Raw(sql).Scan(&res).Error; err != nil {
		fmt.Println("db.raw error:		", err.Error())
	} else {
		fmt.Println("total:		", res.Total)
	}
}
