package main

import (
	"database/sql" //标准库
	"fmt"
	_ "github.com/go-sql-driver/mysql" //我们使用的mysql，需要导入相应驱动包，否则会报错
	"log"
)

var db *sql.DB

type Student struct {
	id   int
	name string
	age  int
}

func initDB() {
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

func queryRowDemo() {
	sqlStr := "select id, name, age from Student where id>?;"

	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u Student
		rows.Scan(&u.id, &u.name, &u.age)
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func Insert() {
	student := [10]Student{
		{0, "a", 18},
		{0, "b", 19},
		{0, "c", 20},
		{0, "d", 21},
		{0, "e", 22},
		{0, "f", 23},
		{0, "g", 24},
		{0, "h", 25},
		{0, "i", 26},
		{0, "j", 27},
	}
	sqlstr := "insert into Student (name,age) values (?,?)"
	for i := 0; i < 10; i++ {
		_, err := db.Exec(sqlstr, student[i].name, student[i].age)
		if err != nil {
			fmt.Printf("insert failed,err:%v\n", err)
			return
		}
	}
	log.Println("insert success")
}

func main() {
	initDB()
	Insert()
	queryRowDemo()
}
