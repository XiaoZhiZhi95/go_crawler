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

	//index.html页面,首页
	http.HandleFunc("/main", handlerIndex)

	//去登陆
	http.HandleFunc("/login", controller.Login)

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

	//监听端口
	http.ListenAndServe(":8080", nil)
}
