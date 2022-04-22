package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306/go_test)"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	// 连接成功 设置最大连接数
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return nil
}

// 创建用户的函数

func createUser(username, password, email string) error {
	sqlstr := "insert into user(username,password,email) values(?,?,?)"
	_, err := db.Exec(sqlstr, username, password, email)
	if err != nil {
		fmt.Println("插入user数据失败")
		return err
	}
	return nil
}
