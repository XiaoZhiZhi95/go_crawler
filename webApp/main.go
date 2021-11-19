package main

import (
	"fmt"
	"net/http"
)

//创建处理器处理函数
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "hello zhizhi", r.URL.Path)
}

//自定义一个处理器
type MyHandler struct {}
func (m *MyHandler)  ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的处理器处理请求！")
}

func main() {
	/* 由HandleFunc转换函数
	http.HandleFunc("/test1", handler)
	//创建路由
	//http.ListenAndServe(":8080", nil)
	 */

	/* 使用自己定义的处理器
	myHandler := MyHandler{}
	http.Handle("/testMyHandler", &myHandler)
	//http.ListenAndServe(":8080", nil)

	//使用自己设置的服务器配置Server
	server := http.Server{
		Addr: ":8080",
		Handler: &myHandler,
		ReadTimeout: 2*time.Second,
	}
	server.ListenAndServe()
	 */

	//自己创建多路复用器处理请求
	mux := http.NewServeMux()
	mux.HandleFunc("/testNewMux", handler)
	http.ListenAndServe(":8080", mux)
}