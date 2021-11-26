package main

import (
	"book/bookstore/controller"
	"net/http"
)

func main() {
	//设置静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	//index.html页面,首页
	http.HandleFunc("/main", controller.HandlerMain)

	//登陆预处理，避免重复登陆
	http.HandleFunc("/toLogin", controller.ToLogin)
	//去登陆
	http.HandleFunc("/login", controller.Login)

	//去注销
	http.HandleFunc("/logout", controller.Logout)

	//去注册
	http.HandleFunc("/register", controller.Register)

	//注册中判断是否存在相同的用户名
	http.HandleFunc("/FindUserByName", controller.FindUserByName)

	//获取所有图书
	http.HandleFunc("/getPageBooks", controller.GetBooks)

	//渲染编辑图书页面
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)

	//更新或者添加图书
	http.HandleFunc("/updateOraddBook", controller.UpdateOrAddBook)

	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)

	//获取价格范围内的数据
	http.HandleFunc("/queryPrice", controller.QueryPrice)

	//监听端口
	http.ListenAndServe(":8080", nil)
}
