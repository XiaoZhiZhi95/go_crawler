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

// GetUserById
/* @Description:
*  @receiver user
*  @return *User
*  @return error
*/
func (user *User) GetUserById() (*User, error)  {
	//1、写sql语句
	sqlStr := "select id, username, password,email from users where id = ?"
	//2、执行，获取一条数据
	row := utils.Db.QueryRow(sqlStr, user.ID)
	//3、声明变量
	u := &User{}
	//4、将查询结果保存至指定参数
	err := row.Scan(&u.ID, &u.UserName, &u.PassWord, &u.Email)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUsers
/* @Description: 获取数据库多条信息
*  @receiver user
*  @return []
*  @return error
*/
func (user *User) GetUsers() ([] *User, error) {
	//1、sql语句
	sqlStr := "select * from users"
	//2、执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("获取User信息失败：", err)
		return nil, err
	}
	//3、循环获取数据
	var users []*User
	for rows.Next() {
		//将查询结果保存至指定参数
		u := &User{}
		err := rows.Scan(&u.ID, &u.UserName, &u.PassWord, &u.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}