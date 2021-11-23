package dao

import (
	"book/bookstore/model"
	"book/testSQL/utils"
	"fmt"
	"strconv"
)

// GetBooks
/* @Description: 获取所有的图书信息
*  @return []*model.Book
*  @return error
*/
func GetBooks() ([]*model.Book, error) {
	sqlStr := "select * from books"

	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		b := model.Book{}
		rows.Scan(&b.ID, &b.Title, &b.Author, &b.Price, &b.Sales, &b.Stock, &b.ImgPath)

		books = append(books, &b)
	}

	return books, nil
}

// GetPageBooks
/* @Description: 获取分页数据
*  @param pageno
*  @param pagesize
*  @return *model.Page
*  @return error
*/
func GetPageBooks(pageno string, pagesize string) (*model.Page, error) {
	var page model.Page

	sqlStr := "select count(1) from books"
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&page.Count)

	//设置每页的数量
	if pagesize == "" {
		pagesize = "4"
	}
	page.PageSize, _ = strconv.ParseInt(pagesize, 10, 0)

	//当前页码
	page.IndexPage, _ = strconv.ParseInt(pageno, 10, 0)

	//计算总页数
	temp := page.Count/page.PageSize
	if page.Count%page.PageSize == 0{
		page.Pages = temp
	}else{
		page.Pages = temp+1
	}

	//获取分页的数据
	sqlStr = "select * from books limit ?, ?"
	rows, _ := utils.Db.Query(sqlStr, (page.IndexPage - 1)*page.PageSize, page.PageSize)

	var books []*model.Book
	for rows.Next() {
		b := &model.Book{}
		rows.Scan(&b.ID, &b.Title, &b.Author, &b.Price, &b.Sales, &b.Stock, &b.ImgPath)

		books = append(books, b)
	}

	page.Books = books

	return &page, nil
}

// AddBook
/* @Description: 添加图书
*  @param b
*  @return error
*/
func AddBook(b *model.Book) error {
	sqlStr := "insert into books(title, author, price, sales, stock, img_path) values(?,?,?,?,?,?)"

	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		fmt.Println("添加图书失败，错误是：", err)
		return err
	}

	return nil
}

// DelBook
/* @Description: 删除图书
*  @param id
*  @return error
*/
func DelBook(id int) error  {
	sqlStr := "delete from books where id = ?"

	_, err := utils.Db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除图书失败，ID = {id}", err)
		return err
	}

	return nil
}

// GetBookById
/* @Description: 根据id获取书的信息
*  @param id
*  @return *model.Book
*  @return error
*/
func GetBookById(id int) (*model.Book, error) {
	sqlStr := "select * from books where id = ?"

	row := utils.Db.QueryRow(sqlStr, id)
	var b model.Book
	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.Price, &b.Sales, &b.Stock, &b.ImgPath)
	if err != nil {
		fmt.Println("获取信息失败: ", err)
		return nil, err
	}

	return &b, nil
}

// UpdateBookById
/* @Description: 更新操作
*  @param book
*  @return error
*/
func UpdateBookById(book *model.Book) error  {
	sqlStr := "update books set title=? , author = ? , price = ? , sales = ? , stock = ? where id = ?"

	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	if err != nil {
		fmt.Println("更新books表失败，id = ", book.ID, err)
	}

	return err
}