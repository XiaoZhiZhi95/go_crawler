package main

import (
	"book/bookstore/controller"
	"html/template"
	"net/http"
)

// handlerIndex
/* @Description: 渲染index.html
*  @param w
*  @param r
*/
func handlerIndex(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("views/myIndex.html"))
	t.Execute(w, "")
}

func main() {
	//设置静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	//index.html页面
	http.HandleFunc("/bookstore", handlerIndex)

	//去登陆
	http.HandleFunc("/login", controller.Login)

	//去注册
	http.HandleFunc("/register", controller.Register)

	//注册中判断是否存在相同的用户名
	http.HandleFunc("/FindUserByName", controller.FindUserByName)

	//监听端口
	http.ListenAndServe(":8080", nil)
}
