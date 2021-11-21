package main

import (
	"book/testSQL/model"
	json2 "encoding/json"
	"fmt"
	"net/http"
)

// handler1
/* @Description: 请求行与请求头内容获取
*  @param w
*  @param r
*/
func handler1(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "发送的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "发送的请求地址的查询信息是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头信息：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息：", r.Header["Accept-Encoding"])	//得到的是一个切片
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息：", r.Header.Get("Accept-Encoding"))	//得到一个string

	//获取请求体的内容
	/*
	//1、获取请求体的内容长度
	lenOfBody := r.ContentLength
	//2、创建byte切片
	body := make([]byte, lenOfBody)
	//3、将请求体body的内容读取到body中
	r.Body.Read(body)
	fmt.Fprintln(w, "请求体的内容为：", string(body))
	 */


	//获取请求参数form字段
	/*
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(w, "错误", err)
	}

	fmt.Fprintln(w, "所有请求参数：", r.Form)
	fmt.Fprintln(w, "post请求参数:", r.PostForm)
	 */

	fmt.Fprintln(w, "请求参数", r.FormValue("lili"))
	fmt.Fprintln(w, "请求参数", r.PostFormValue("username"))
}

// testJsonRes
/* @Description: 返回json格式的数据
*  @param w
*  @param r
*/
func testJsonRes(w http.ResponseWriter, r *http.Request)  {
	//设置响应内容的类型
	w.Header().Set("Content-Type", "application/json")

	//创建User
	user := model.User{
		1,
		"admin",
		"12345",
		"admin@lili.com",
	}
	//将user转换成json格式
	json,_ := json2.Marshal(user)
	w.Write(json)
}

// testRedirect
/* @Description: 测试重定向
*  @param w
*  @param r
*/
func testRedirect(w http.ResponseWriter, r *http.Request){
	//设置响应头中的Location属性
	w.Header().Set("Location", "https://www.baidu.com")
	//设置响应的状态码
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/query", handler1)
	http.HandleFunc("/testJson", testJsonRes)
	http.HandleFunc("/testRedirect", testRedirect)

	http.ListenAndServe(":8080", nil)
}
