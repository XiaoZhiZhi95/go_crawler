package dao

import (
	"fmt"
	"testing"
)

/*	//对用户信息的测试
func TestCheckUserNameAndPassword(t *testing.T) {
	userName := "lili"
	passwd := "123456"
	user,err := CheckUserNameAndPassword(userName, passwd)
	fmt.Println(user, err)
}

func TestCheckUserName(t *testing.T) {
	username := "lili"
	user, err := CheckUserName(username)
	fmt.Println(user, err)
}

func TestSaveUser(t *testing.T) {
	uname := "zhizhi"
	passwd := "fz..555FZ"
	em := "zhizhi@lili.com"
	err := SaveUser(uname, passwd, em)
	fmt.Println("保存", err)
}
 */

//// TestGetBooks
///* @Description: 测试获取图书
//*  @param t
//*/
//func TestGetBooks(t *testing.T) {
//	books, _ := GetBooks()
//	for _, v := range books {
//		fmt.Println(v)
//	}
//}

//func TestAddBook(t *testing.T) {
//	b := &model.Book{
//		Title: "剑指Offer",
//		Author: "何海淘",
//		Price: 80.00,
//		Sales: 1000,
//		Stock: 2000,
//		ImgPath: "static/img/default.jpg",
//	}
//
//	err := AddBook(b)
//	fmt.Println(err)
//}

//func TestDelBook(t *testing.T) {
//	id := 45
//	DelBook(id)
//}

//func TestGetBookById(t *testing.T) {
//	id := 3
//	b,_ := GetBookById(id)
//	fmt.Println(b)
//}

//func TestUpdateBookById(t *testing.T) {
//	b := &model.Book{
//		3,
//		"中国哲学史1",
//		"冯友兰1",
//		441,
//		1011,
//		991,
//		"",
//	}
//	UpdateBookById(b)
//}

//func TestGetPageBooks(t *testing.T) {
//	pageno := "2"
//	pagesize := "5"
//	page,_ := GetPageBooks(pageno, pagesize)
//	for _, v := range page.Books{
//		fmt.Println(v)
//	}
//	fmt.Println(page)
//}

/*func TestGetPageBooksByPrice(t *testing.T) {
	pageno := "2"
	pagesize := "5"
	minprice := "100"
	maxPrice := "110"
	page,_ := GetPageBooksByPrice(pageno, pagesize, minprice, maxPrice)
	for _, v := range page.Books{
		fmt.Println(v)
	}
	fmt.Println(page)
}*/

/*
func TestAddSession(t *testing.T) {
	s := model.Session{
		"asdsd",
		"zhizhi",
		17,
	}
	err := AddSession(&s)
	fmt.Println(err)
}*/

/*func TestDelBook(t *testing.T) {
	sessionId := "asdsd"
	err := DelSessionBySessionId(sessionId)
	fmt.Println(err)
}
*/

func TestGetSessionBySessionId(t *testing.T) {
	sid := "756ebf9c-4cb8-4214-4eb5-e1b751fec581"
	session, err := GetSessionBySessionId(sid)
	fmt.Println("session :" , session)
	fmt.Println("err: ", err)
}