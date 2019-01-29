package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var DB *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:Fnw12345@tcp(10.39.6.202:3306)/test")
	if err != nil {
		fmt.Println("conntion mysql failed, ", err)
		return
	}
	DB = database
}

func main() {
	r, err := DB.Exec("update person set username=? where user_id=?", "stu0001", 1)
	if err != nil {
		fmt.Println("update data failed,", err)
		return
	}
	_, err = r.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId data failed,", err)
		return
	}
}
