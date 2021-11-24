package main

import (
	"fmt"
	"net/http"
)

// SetCookie
/* @Description: 设置cookie
*  @param w
*  @param r
*/
func SetCookie(w http.ResponseWriter, r *http.Request)  {
	//创建cookie1
	cookie1 := http.Cookie{
		Name: "user",
		Value: "admin",
		HttpOnly: true,
		MaxAge: 60,		//单位为秒
	}
	cookie2 := http.Cookie{
		Name: "zhizhi",
		Value: "lili",
		HttpOnly: true,
	}

	//将cookie发送给浏览器，在响应头设置
	w.Header().Set("Set_Cookie", cookie1.String())
	w.Header().Add("Set_Cookie", cookie2.String())
	//http.SetCookie(w, &cookie2)
	//http.SetCookie(w, &cookie1)
}

func GetCookies(w http.ResponseWriter, r *http.Request)  {
	//方式一，获取所有cookie
	c := r.Header["Cookie"]
	fmt.Println(c)

	//方式二，得到某个cookie
	cookieUser,_ := r.Cookie("user")
	fmt.Println(cookieUser)
}

func main()  {
	http.HandleFunc("/testcookie", SetCookie)
	http.HandleFunc("/getcookies", GetCookies)

	http.ListenAndServe(":8080", nil)
}
