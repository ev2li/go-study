package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

func main() {
	test501()
}


type Person struct{
	UserId		int		`db:"user_id"`
	UserName 	string 	`db:"username"`
	Sex 		string	`db:"sex"`
	Email 		string 	`db:"email"`
}

type Place struct {
	Country		string 	`db:"country"`
	City		string	`db:"city"`
	TelCode		int		`db:"telcode"`
}

var Db *sqlx.DB

func init(){
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed, err:", err)
		return
	}
	Db = database
}

func test501() {
	r, err := Db.Exec("insert into person(username, sex, email) value(?,?,)", "stu001","man", "stu@qq.com")
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,",err)
		return
	}
	fmt.Println("insert succ", id)
}

func test502(){
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where use_id=?",1)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	fmt.Println("select succ", person)
}

func test503(){
	_, err := Db.Exec("update person set username=? where user_id=?", "stu001",1)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
}

func test504(){
	_, err := Db.Exec("delete from person where user_id=?", "stu001",1)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
}