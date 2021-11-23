package model

// Book
/*  @Description: 图书的结构
*/
type Book struct {
	ID int
	Title string	//书名
	Author string	//作者
	Price float64	//价格
	Sales int		//销量
	Stock int		//库存
	ImgPath	string	//图片地址
}