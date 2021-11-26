package model

// CartItem
/*  @Description: 购物项
*/
type CartItem struct {
	CardItemId int
	Book *Book		//图书信息
	Count int		//图书数量
	Amount float64	//金额小记，通过计算得到
	CartId string	//购物车id(uuid)
}

func (cartItem *CartItem)GetAmount() float64 {
	//获取当前购物项中的价格
	price := cartItem.Book.Price
	return float64(cartItem.Count)*price
}