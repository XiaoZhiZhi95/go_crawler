package dao

import (
	"book/bookstore/model"
	"book/bookstore/utils"
	"fmt"
)

// AddCartItem
/* @Description: 添加新的购物项
*  @param cartitem
*  @return error
*/
func AddCartItem(cartitem *model.CartItem) error {
	sqlStr := "insert into cart_items(COUNT, amount, book_id, cart_id) values(?,?,?,?)"

	_, err := utils.Db.Exec(sqlStr, cartitem.Count, cartitem.GetAmount(), cartitem.Book.ID, cartitem.CartId)
	if err != nil {
		fmt.Println("插入购物项失败：", err)
	}

	return err
}

// GetCartItemByBookId
/* @Description: 根据书的id获取购物项
*  @param bookId
*  @return *model.CartItem
*  @return error
*/
func GetCartItemByBookIdAndCartId(bookId int, cartId string) (*model.CartItem, error) {
	sqlStr := "select * from cart_items where book_id = ? and cart_id = ?"
	row := utils.Db.QueryRow(sqlStr, bookId, cartId)


	row.
	if err != nil {
		fmt.Println("查询图书是否存在购物项失败，bookId, cartId, err :", bookId, cartId, err)
		return nil, err
	}
}

func GetCartItemsByCartId()  {

}
