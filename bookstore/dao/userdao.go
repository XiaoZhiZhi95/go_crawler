package dao

import (
	"book/testSQL/model"
	"book/testSQL/utils"
	"fmt"
)

// CheckUserNameAndPassword
/* @Description: 根据用户名和密码从数据库中查询一条记录
*  @param username
*  @param password
*  @return *model.User
*  @return error
*/
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	sqlStr := "select * from users where username = ? and password = ?"

	row := utils.Db.QueryRow(sqlStr, username, password)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	if err != nil {
		fmt.Println("查询用户名和密码失败:", err)
		return nil, nil
	}

	return user, nil
}

// CheckUserName
/* @Description: 根据用户名查询
*  @param username
*  @return *model.User
*  @return error
*/
func CheckUserName(username string) (*model.User, error) {
	sqlStr := "select * from users where username = ?"

	row := utils.Db.QueryRow(sqlStr, username)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	if err != nil {
		fmt.Println("查询用户名失败: ", err)
		return nil, err
	}

	return user, nil
}

// SaveUser
/* @Description: 新增用户
*  @param username
*  @param password
*  @param email
*  @return error
*/
func SaveUser(username string, password string, email string) error {
	sqlStr := "insert into users(username, password, email) values(?,?,?)"

	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		fmt.Println("新增用户失败: ", err)
		return err
	}

	return nil
}
