package controller

import (
	"book/bookstore/dao"
	"book/bookstore/model"
	"book/bookstore/utils"
	"fmt"
	"html/template"
	"net/http"
)

// ToLogin
/* @Description: 检验用户是否登陆，防止重复登陆
*  @param w
*  @param r
*/
func ToLogin(w http.ResponseWriter, r *http.Request)  {
	session := utils.IsLogin(r, utils.CookieName)
	if session != nil {
		//重定向
		http.Redirect(w, r, "/main", 302)
	}else{
		t := template.Must(template.ParseFiles("views/pages/user/logining.html"))
		t.Execute(w, "")
	}
}

// Login
/* @Description: 处理用户登陆
*  @param w
*  @param r
*/
func Login(w http.ResponseWriter, r *http.Request)  {
	//获取用户名和密码
	uname := r.PostFormValue("username")
	passwd := r.PostFormValue("password")

	user, err := dao.CheckUserNameAndPassword(uname, passwd)
	if err != nil {
		fmt.Fprintln(w, "获取信息失败，请稍后重试", err)
	}

	if user != nil {
		//用户名正确，则解析登陆成功的页面
		fmt.Println("登陆成功")

		//创建会话
		uuid := utils.CreateUUID()
		session := &model.Session{
			uuid,
			uname,
			user.ID,
		}
		//保存cookie
		errAddSession := dao.AddSession(session)
		if errAddSession != nil {
			//创建session失败
			fmt.Println("创建session失败啦")
		}else {
			//创建cookie
			cookie := http.Cookie{
				Name: "user",
				Value: uuid,
				HttpOnly: true,
			}

			//发送cookie
			http.SetCookie(w, &cookie)
		}

		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, uname)
	}else {
		fmt.Println("返回登陆页面")
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

// Logout
/* @Description: 注销
*  @param w
*  @param r
*/
func Logout(w http.ResponseWriter, r *http.Request)  {
	//获取cookie
	cookie,_ := r.Cookie("user")
	if cookie != nil {
		sessionId := cookie.Value

		//删除数据库中对应的session
		err := dao.DelSessionBySessionId(sessionId)
		if err != nil {	//存在session，表示已登陆
			fmt.Println("删除session失败：", err)
		}

		//设置cookie失效
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}

	//去首页
	GetBooks(w, r)
}

// Register
/* @Description: 注册
*  @param w
*  @param r
*/
func Register(w http.ResponseWriter, r *http.Request)  {
	//获取用户名和密码
	uname := r.PostFormValue("username")
	passwd := r.PostFormValue("password")
	reps := r.PostFormValue("repwd")
	email := r.PostFormValue("email")

	//校验密码与确认的密码(前端有校验)
	if passwd != reps {
		fmt.Println("密码与确认密码不一致")
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "密码与确认密码不一致")
	}

	//查询是否存在相同的用户名和密码
	user, _ := dao.CheckUserNameAndPassword(uname, passwd)

	if user != nil {
		//用户存在，则跳转当前注册页面
		fmt.Println("用户名存在，请修改后注册")
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	}else {
		//用户不存在，添加用户
		fmt.Println("添加用户信息")
		err := dao.SaveUser(uname, passwd, email)
		if err != nil {
			fmt.Println("添加用户信息失败，请稍后重试")
			fmt.Fprintln(w, "添加用户信息失败，请稍后重试")
		}

		//添加用户成功，则跳转至注册成功界面
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

// FindUserByName
/* @Description: 查询是否存在相同的用户名
*  @param w
*  @param r
*/
func FindUserByName(w http.ResponseWriter, r *http.Request)  {
	//获取用户名
	uname := r.PostFormValue("username")

	user, _ := dao.CheckUserName(uname)

	if user != nil {
		//用户名存在
		//fmt.Fprintln(w,"用户名存在")
		w.Write([]byte("用户名存在!"))
	}else {
		//用户名不存在
		//fmt.Fprintln(w, "")

		w.Write([]byte("<front style='color:green'>用户名可用！</front>"))
	}
}