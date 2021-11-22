package main

import (
	"html/template"
	"net/http"
)

func testIf(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("if.html"))
	age := 17
	t.Execute(w, age > 18)
}

type A struct {
	Id int
	Name string
}
func testRange(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("range.html"))
	//a := []int{1,2,3,4}
	var a  []A
	a = append(a, A{1, "zhizhi"})
	a = append(a, A{2,"lili"})

	t.Execute(w, a)
}

func main() {
	http.HandleFunc("/testif", testIf)
	http.HandleFunc("/testrange", testRange)

	http.ListenAndServe(":8080", nil)
}
