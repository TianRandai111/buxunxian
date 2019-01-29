package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	//	database, err := sqlx.Open("mysql", "root:Fnw12345@tcp(10.39.6.202:3306)/test")
	database, err := sqlx.Open("mysql", "root:Fnw12345@tcp(10.39.6.202:3306)/test")
	if err != nil {
		fmt.Println("数据库链接失败:", err)
		return
	}
	DB = database

}
func main() {
	_, err := DB.Exec("delete from person where user_id =?", 1)
	if err != nil {
		fmt.Println("删除失败:", err)
		return
	}
	fmt.Println("删除成功")
}
