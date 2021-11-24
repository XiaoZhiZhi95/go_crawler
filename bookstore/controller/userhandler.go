package controller

import (
	"book/bookstore/dao"
	"fmt"
	"html/template"
	"net/http"
)

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
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, uname)
	}else {
		fmt.Println("返回登陆页面")
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
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