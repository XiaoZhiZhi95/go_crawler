package dao

import (
	"book/bookstore/model"
	"book/testSQL/utils"
	"fmt"
)

// AddSession
/* @Description: 新增会话
*  @param session
*  @return error
*/
func AddSession(session *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"

	_, err := utils.Db.Exec(sqlStr, session.SessionId, session.Username, session.UserId)
	if err != nil {
		fmt.Println("插入session表失败: ", err)
	}

	return err
}

// GetSessionBySessionId
/* @Description: 获取session的信息
*  @param sid
*  @return *model.Session
*  @return error
*/
func GetSessionBySessionId(sid string) (*model.Session, error)  {
	sqlStr := "select * from sessions where session_id = ?"

	//预编译
	instmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("获取信息失败: ", err)
		return nil, err
	}
	//查询
	row :=instmt.QueryRow(sid)
	session := &model.Session{}

	err = row.Scan(&session.SessionId, &session.Username, &session.UserId)
	if err != nil {
		fmt.Println("读取session信息失败:", err)
		return nil, err
	}

	return session, nil
}

// DelSessionBySessionId
/* @Description: 删除session
*  @param sessionId
*  @return error
*/
func DelSessionBySessionId(sessionId string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr, sessionId)
	if err != nil {
		fmt.Println("删除session表失败: ", err)
	}

	return err
}
