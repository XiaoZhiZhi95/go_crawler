package model

type Page struct {
	Books []*Book	//每页的图书
	IndexPage int64	//当前页码
	PageSize int64	//每页的条数
	Pages int64		//总页数
	Count int64		//总记录条数
	Username string	//管理员姓名
}

// IsHasPrev
/* @Description: 判断是否有上一页
*  @receiver p
*  @return bool
*/
func (p *Page)IsHasPrev() bool {
	return p.IndexPage > 1
}

// IsHasNext
/* @Description: 判断是否有下一页
*  @receiver p
*  @return bool
*/
func (p *Page)IsHasNext() bool {
	return p.IndexPage < p.Pages
}

// GetPrevPageNo
/* @Description: 返回上一页的页码
*  @receiver p
*  @return int64
*/
func (p *Page)GetPrevPageNo() int64 {
	return p.IndexPage -1
}

// GetNextPageNo
/* @Description: 返回下一页的页码
*  @receiver p
*  @return int64
*/
func (p *Page)GetNextPageNo() int64 {
	return p.IndexPage +1
}