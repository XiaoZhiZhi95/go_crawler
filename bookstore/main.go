package main

import (
	"html/template"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, "")
}

func main() {
	//设置静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.HandleFunc("/bookstore", handlerIndex)
	http.ListenAndServe(":8080", nil)
}
