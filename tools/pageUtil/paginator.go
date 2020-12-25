package pageUtil

import (
	"math"
)

//分页
type Paginator struct {
	Page        int                    `json:"pageUtil"`      //当前页
	Pages       []int                  `json:"pages"`     //页码数组
	PageSize    int                    `json:"pageSize"`  //每页条数
	TotalPage   int                    `json:"totalPage"` //总页码
	TotalCount  int                    `json:"count"`     //总条数
	FirstPage   int                    `json:"first"`     //上一页
	FirstPageIs bool                   `json:"firstIs"`
	LastPageIs  bool                   `json:"lastIs"` //
	LastPage    int                    `json:"last"`   //尾页
	NextPageIs  bool                   `json:"nextIs"` //
	NextPage    int                    `json:"next"`   //下一页
	Data        []interface{}          `json:"data"`
	OtherData   map[string]interface{} `json:"dataOther"`
	Offset      int                    `json:"offset"`
}

// 分页  总数，当前页，每页条数
func Pagination(count int, page int, pageSize int) *Paginator {
	Page := new(Paginator)
	Page.PageSize = pageSize
	Page.TotalCount = count
	Page.TotalPage = int(math.Ceil(float64(count) / float64(pageSize))) //page总数
	//if count % pageSize > 0 {
	//	Page.TotalPage = count / pageSize + 1
	//}

	if page > Page.TotalPage {
		page = Page.TotalPage
	}
	if page <= 0 {
		page = 1
	}
	Page.Page = page
	//当前页
	Page.FirstPageIs = page != 1
	Page.NextPageIs = page != Page.TotalPage
	if page == 1 {
		Page.NextPageIs = false
		if Page.TotalPage > 5 {
			Page.LastPageIs = true
		}
	} else {
		Page.LastPageIs = Page.NextPageIs
		if Page.TotalPage > 5 && page != Page.TotalPage {
			Page.LastPageIs = true
		}
	}
	Page.LastPage = Page.TotalPage
	//读取起始条数
	Page.Offset = (page - 1) * pageSize
	//fmt.Println("Page.TotalPage", Page.TotalPage)
	pages := make([]int, 5)
	switch {
	case page > Page.TotalPage-5 && Page.TotalPage > 5: //最后5页
		Page.Pages = make([]int, 5)
		start := Page.TotalPage - 5 + 1
		Page.FirstPage = page - 1
		Page.NextPage = int(math.Min(float64(Page.TotalPage), float64(page+1)))
		for i, _ := range pages {
			Page.Pages[i] = start + i
		}
	case page >= 3 && Page.TotalPage > 5:
		Page.Pages = make([]int, 5)
		start := page - 3 + 1
		Page.FirstPage = page - 3
		for i, _ := range pages {
			Page.Pages[i] = start + i
		}
		Page.FirstPage = page - 1
		Page.NextPage = page + 1
	default:
		if Page.TotalPage > 1 {
			num := int(math.Min(float64(5), float64(Page.TotalPage)))
			Page.Pages = make([]int, num)
			for i := 0; i < num; i++ {
				Page.Pages[i] = i + 1
			}
			Page.FirstPage = int(math.Max(float64(1), float64(page-1)))
			if page < Page.TotalPage {
				Page.NextPage = page + 1
			}
		}
	}
	return Page
}
