package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init() {
	//验证数据库参数，并不创建数据库的连接
	Db, err = sql.Open("mysql", "root:root1234@tcp(127.0.0.1:13306)/goTest")
	if err != nil {
		fmt.Println("连接出现异常", err)
		panic(err.Error())	//打印异常信息
	}
}
