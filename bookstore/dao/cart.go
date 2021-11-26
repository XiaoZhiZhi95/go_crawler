package dao

import (
	"book/bookstore/model"
	"book/bookstore/utils"
	"fmt"
)

// AddCart
/* @Description: 创建一个购物车，并创建其中的购物项
*  @param cart
*  @return error
*/
func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts values(?,?,?,?)"

	_,err := utils.Db.Exec(sqlStr, cart.CartId, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserId)
	if err != nil {
		fmt.Println("插入购物项失败：", err)
		return err
	}

	//遍历得到购物车中的每一个购物项，将其插入数据库
	for _, cartItem := range cart.CartItems {
		err = AddCartItem(cartItem)
		if err != nil {
			fmt.Println("创建购物车时，一个购物项添加失败，cartItem: ", cartItem)
		}
	}

	return nil
}

func GetCartByUserId(uId int)   {

}
