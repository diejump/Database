package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Username string
	Account  string
	Password string
}

var db *sql.DB

func InitDB() {
	var err error

	dsn := "root:123456@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return
}

func AddUser(Account, Username, Password string) {
	sqlstr := "insert into user (username,account,password) values (?,?,?)"
	_, err := db.Exec(sqlstr, Username, Account, Password)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	log.Println("insert success")
}

func SelectUser(Account string) bool {
	sqlStr := "select password from user where account =?"
	var password string
	db.QueryRow(sqlStr, Account).Scan(&password)
	if password != "" {
		return true
	}
	return false
}

func SelectPasswordFromAccount(Account string) string {
	sqlstr := "select password from user where account=?"
	var password string
	db.QueryRow(sqlstr, Account).Scan(&password)
	return password
}
