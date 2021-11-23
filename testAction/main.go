package main

import (
	"html/template"
	"net/http"
)

// testIf
/* @Description: 测试条件动作
*  @param w
*  @param r
*/
func testIf(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("if.html"))
	age := 17
	t.Execute(w, age > 18)
}

type A struct {
	Id int
	Name string
}
// testRange
/* @Description: 测试迭代动作
*  @param w
*  @param r
*/
func testRange(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("range.html"))
	//a := []int{1,2,3,4}
	var a  []A
	a = append(a, A{1, "zhizhi"})
	a = append(a, A{2,"lili"})

	t.Execute(w, a)
}

// testWith
/* @Description: 测试重置参数
*  @param w
*  @param r
*/
func testWith(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("with.html"))
	str := "狸猫猫"
	t.Execute(w, str)
}

// testTemp
/* @Description: 测试模版的嵌套
*  @param w
*  @param r
*/
func testTemp(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("temp1.html", "temp2.html"))
	str := "狸猫猫"
	t.Execute(w, str)
}

// testDefine
/* @Description: 测试自定义模版
*  @param w
*  @param r
*/
func testDefine(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("define.html"))
	t.ExecuteTemplate(w, "zhizhi", "")
}

// testBlock
/* @Description: 测试块动作
*  @param w
*  @param r
*/
func testBlock(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("block.html"))
	t.ExecuteTemplate(w,"zhizhi", "")
}

func main() {
	http.HandleFunc("/testif", testIf)
	http.HandleFunc("/testrange", testRange)
	http.HandleFunc("/testwith", testWith)
	http.HandleFunc("/testtemp", testTemp)
	http.HandleFunc("/testdefine", testDefine)
	http.HandleFunc("/testblock", testBlock)

	http.ListenAndServe(":8080", nil)
}
