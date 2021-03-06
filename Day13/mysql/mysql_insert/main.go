package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserID   int    `db:"user_id"`
	UserName string `db:"user_name"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Piace struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:telcode`
}

var DB *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:@tcp(10.39.6.202:3306)/Log_Agent_Test")
	if err != nil {
		fmt.Println("mysql 链接失败", err)
		return
	}
	DB = database
}

func main() {
	r, err := DB.Exec("insert into person(user_name,sex,email)values(?,?,?)", "stu002", "man", "stu002@qq.com")

	if err != nil {
		fmt.Println("插入数据出错", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("插入失败", err)
		return
	}
	fmt.Println("insert succ", id)
}
