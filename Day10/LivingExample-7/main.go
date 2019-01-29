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

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:Fnw12345@tcp(10.39.6.202:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	test := &Person{
		Username: "buxunxian",
		Sex:      "main",
		Email:    "buxunxian@test.com",
	}
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", test.Username, test.Sex, test.Email)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}
