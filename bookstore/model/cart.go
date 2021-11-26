package model

type Cart struct {
	CartId string
	CartItems []*CartItem	//购物车中所有的购物项
	TotalCount int			//购物车中图书的总数量
	TotalAmount float64		//购物车中图书的总金额
	UserId int				//当前购物车所属的用户
}

// GetTotalCount
/* @Description: 获取购物车中商品总数
*  @receiver cart
*  @return int
*/
func (cart *Cart)GetTotalCount() int {
	var totalCount int
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

// GetTotalAmount
/* @Description: 获取购物车的总金额
*  @receiver cart
*  @return float64
*/
func (cart *Cart)GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}
