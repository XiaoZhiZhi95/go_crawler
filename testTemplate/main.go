package main

import (
	"html/template"
	"net/http"
)

// testTemplate
/* @Description: 模版
*  @param w
*  @param r
*/
func testTemplate(w http.ResponseWriter, r *http.Request)  {
	//解析模版
	//t, _ := template.ParseFiles("./index.html")
	//执行模版
	//t.Execute(w, "Hello lili")

	t := template.Must(template.ParseFiles("./index.html"))
	t.Execute(w, "hello lili")
}

func main() {
	http.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8080", nil)
}
