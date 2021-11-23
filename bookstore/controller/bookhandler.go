package controller

import (
	"book/bookstore/dao"
	"book/bookstore/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// GetBooks
/* @Description: 获取所有图书，图书管理页面
*  @param w
*  @param r
*/
func GetBooks(w http.ResponseWriter, r *http.Request)  {
	//解析当前页
	requestPageNo := r.FormValue("PageNo")
	if requestPageNo == ""{
		requestPageNo = "1"
	}

	//获取分页数据
	page, err :=dao.GetPageBooks(requestPageNo, "")
	if err != nil {
		fmt.Println("获取图书信息失败：", err)
	}

	//解析模版文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, page)
}

// ToUpdateBookPage
/* @Description: 渲染编辑书本的页面
*  @param w
*  @param r
*/
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request)  {
	//根据bookId参数判断是新增还是修改
	updateFlag := true
	if r.FormValue("bookId") == "" {	//如果不存在bookId参数，则是新增
		updateFlag = false
	}

	var book *model.Book
	if updateFlag {
		requestId, _ := strconv.ParseInt(r.FormValue("bookId"), 10, 0)
		fmt.Println("requestId", requestId)
		id := int(requestId)

		//获取图书信息
		var err error
		book, err = dao.GetBookById(id)
		if err != nil {
			fmt.Println("图书信息为空，err: ", err)
		}
	}else {
		book = nil
	}

	t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
	t.Execute(w, book)
}

// UpdateOrAddBook
/* @Description: 修改或新增图书
*  @param w
*  @param r
*/
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request)  {
	//解析参数
	b := &model.Book{}
	b.Title = r.PostFormValue("title")
	b.Author = r.PostFormValue("author")
	//string类型转化
	b.Price, _ = strconv.ParseFloat(r.PostFormValue("price"), 64)

	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
	stock, _ := strconv.ParseInt(r.PostFormValue("stock"),10,0)

	b.Sales = int(sales)
	b.Stock = int(stock)
	b.ImgPath = "/static/img/default.jpg"

	//根据bookId参数判断是新增还是修改
	updateFlag := true
	if r.FormValue("bookId") == "" {	//如果不存在bookId参数，则是新增
		updateFlag = false
	}

	if updateFlag {	//更新
		//解析ID参数
		requestId, _ := strconv.ParseInt(r.FormValue("bookId"), 10, 0)
		b.ID = int(requestId)

		//更新
		dao.UpdateBookById(b)
	}else {	//新增
		if b.Title != "" {	//防止刷新页面的时候添加空白的行
			//调用添加图书的函数
			dao.AddBook(b)
		}
	}

	//再查一次
	GetBooks(w, r)
}

// DeleteBook
/* @Description: 删除图书
*  @param w
*  @param r
*/
func DeleteBook(w http.ResponseWriter, r *http.Request)  {
	bookId,_ := strconv.ParseInt( r.FormValue("bookId"), 10, 0)
	id := int(bookId)

	//删除图书
	err := dao.DelBook(id)
	if err != nil {
		fmt.Println("删除失败！")
	}

	GetBooks(w, r)
}