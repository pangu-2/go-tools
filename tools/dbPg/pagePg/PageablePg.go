package pagePg

import (
	"fmt"
	"math"
)

// PageablePg 分页
// @Description:
type PageablePg struct {
	Total       int64   `json:"total"`     //总条数
	PageNum     int64   `json:"pageNum"`   //当前页
	TotalPage   int64   `json:"totalPage"` //总页码
	PageSize    int64   `json:"pageSize"`  //每页条数
	Pages       []int64 `json:"pages"`     //页码数组
	FirstPage   int64   `json:"first"`     //上一页
	FirstPageIs bool    `json:"firstPage"`
	LastPageIs  bool    `json:"lastPage"`    //
	LastPage    int64   `json:"last"`        //尾页
	NextPageIs  bool    `json:"hasNextPage"` //
	NextPage    int64   `json:"nextPage"`    //下一页
	Offset      int64
}

func NewPageablePg(total int64, page int64, pageSize int64) PageablePg {
	pg := PageablePg{Total: total, PageNum: page, PageSize: pageSize}
	pg.MakePageable()
	return pg
}

// 分页  总数，当前页，每页条数
func (c PageablePg) MakePageable() PageablePg {
	if c.Total == 0 {
		c.TotalPage = 1
	} else {
		c.TotalPage = int64(math.Ceil(float64(c.Total) / float64(c.PageSize))) //page总数
	}

	//if count % pageSize > 0 {
	//	PageNum.TotalPage = count / pageSize + 1
	//}
	page := c.PageNum
	if page > c.TotalPage {
		page = c.TotalPage
	}
	if page <= 0 {
		page = 1
	}
	c.PageNum = page
	//当前页
	c.FirstPageIs = page != 1
	c.NextPageIs = page != c.TotalPage
	if page == 1 {
		c.NextPageIs = false
		if c.TotalPage > 5 {
			c.LastPageIs = true
		}
	} else {
		c.LastPageIs = c.NextPageIs
		if c.TotalPage > 5 && page != c.TotalPage {
			c.LastPageIs = true
		}
	}
	c.LastPage = c.TotalPage
	//读取起始条数
	c.Offset = (page - 1) * c.PageSize
	if c.TotalPage == 0 {
		c.TotalPage = 1
	}
	fmt.Println("PageNum.TotalPage", c.TotalPage)
	pages := make([]int64, 5)
	switch {
	case page > c.TotalPage-5 && c.TotalPage > 5: //最后5页
		c.Pages = make([]int64, 5)
		start := c.TotalPage - 5 + 1
		c.FirstPage = page - 1
		c.NextPage = int64(math.Min(float64(c.TotalPage), float64(page+1)))
		for i, _ := range pages {
			c.Pages[i] = start + int64(i)
		}
	case page >= 3 && c.TotalPage > 5:
		c.Pages = make([]int64, 5)
		start := page - 3 + 1
		c.FirstPage = page - 3
		for i, _ := range pages {
			c.Pages[i] = start + int64(i)
		}
		c.FirstPage = page - 1
		c.NextPage = page + 1
	default:
		if c.TotalPage > 1 {
			num := int64(math.Min(float64(5), float64(c.TotalPage)))
			c.Pages = make([]int64, num)
			for i := int64(0); i < num; i++ {
				c.Pages[i] = i + 1
			}
			c.FirstPage = int64(math.Max(float64(1), float64(page-1)))
			if page < c.TotalPage {
				c.NextPage = page + 1
			}
		}
	}
	return c
}
