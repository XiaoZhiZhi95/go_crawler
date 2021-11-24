package utils

import (
	"book/bookstore/dao"
	"book/bookstore/model"
	"fmt"
	"net/http"
)

func IsLogin(r *http.Request, cookieName string) *model.Session  {
	//获取cookie
	cookie,_ := r.Cookie(cookieName)
	if cookie == nil {
		return nil
	}

	//去数据库查询对应的session
	sessionId := cookie.Value
	session,_ := dao.GetSessionBySessionId(sessionId)
	if session != nil {	//存在session，表示已登陆
		fmt.Println("session存在：", session)
		return session
	}else {
		fmt.Println("session不存在，请登录")
		return nil
	}
}
