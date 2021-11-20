package model

import (
	"book/testSQL/utils"
	"fmt"
)

// User 表的结构体
type User struct {
	ID int
	UserName string
	PassWord string
	Email string
}

// AddUser1
/**
* AddUser1
* @Description: 插入数据库方法1
* @receiver user
* @return error
*/
func (user *User) AddUser1() error  {
	//1、写SQL语句
	sqlStr := "insert into users(username, password, email) values(?,?,?)"

	//2、预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常", err)
		return err
	}
	//3、执行
	_, err2 := inStmt.Exec("admin", "123456", "admin@zhizhi.com")
	if err2 != nil {
		fmt.Println("执行出现异常：", err2)
		return err2
	}
	return nil
}

// AddUser2
/**
* AddUser2
* @Description: 插入数据库方法2
* @receiver user
* @return error
*/
func (user *User) AddUser2() error  {
	//1、写SQL语句
	sqlStr := "insert into users(username, password, email) values(?,?,?)"

	//3、执行
	_, err := utils.Db.Exec(sqlStr, "lili", "123456", "lili@zhizhi.com")
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}