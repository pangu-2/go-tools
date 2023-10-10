package pagePg

import (
	"gorm.io/gorm"
)

// 分页
type PaginatorPg[T any] struct {
	PageNum   int64                  `json:"pageNum"`   //当前页
	PageSize  int64                  `json:"pageSize"`  //每页条数
	TotalPage int64                  `json:"totalPage"` //总页码
	Total     int64                  `json:"total"`     //总条数
	Data      []T                    `json:"data"`
	MapData   map[string]interface{} `json:"map"`
	Pageable  PageablePg             `json:"pageable"` //分页数据
	OffsetId  string                 `json:"offsetId"` //分页数据 ids
}

type Option[T any] func(c *PaginatorPg[T])

func NewPaginatorPg[T any](options ...Option[T]) PaginatorPg[T] {
	opts := PaginatorPg[T]{
		PageNum:  1,
		PageSize: 20,
		Data:     []T{},
	}

	for _, option := range options {
		option(&opts)
		if opts.PageNum <= 0 {
			opts.PageNum = 1
		}
	}
	// 如果没有分页，那么进行初始化配置
	pg := NewPageablePg(opts.Total, opts.PageNum, opts.PageSize)
	opts.SetPageable(pg)
	return opts
}

func (c PaginatorPg[T]) GetPageable() PageablePg {
	return c.Pageable
}

func (c PaginatorPg[T]) SetPageable(p PageablePg) {
	c.Pageable = p
	c.PageNum = p.PageNum
	c.PageSize = p.PageSize
	c.TotalPage = p.TotalPage
	c.Total = p.Total
}

func (c PaginatorPg[T]) SetPageable2(p *PageablePg) {
	//
	c.Pageable = PageablePg{
		Total:       p.Total,
		PageNum:     p.PageNum,
		TotalPage:   p.TotalPage,
		PageSize:    p.PageSize,
		Pages:       p.Pages,
		FirstPage:   p.FirstPage,
		FirstPageIs: p.FirstPageIs,
		LastPageIs:  p.LastPageIs,
		LastPage:    p.LastPage,
		NextPageIs:  p.NextPageIs,
		NextPage:    p.NextPage,
		Offset:      p.Offset,
	}
	c.PageNum = p.PageNum
	c.PageSize = p.PageSize
	c.TotalPage = p.TotalPage
	c.Total = p.Total
}

func (c PaginatorPg[T]) Scopes() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		size := c.PageSize
		offset := int((c.PageNum - 1) * size)
		return db.Offset(offset).Limit(int(size))
	}
}

// Paginate加上T就行
func PaginateScopes[T any](page PaginatorPg[T]) func(db *gorm.DB) *gorm.DB {
	return page.Scopes()
}

func NewWithPaginatorPg[T any](options ...Option[T]) PaginatorPg[T] {
	opts := PaginatorPg[T]{
		PageNum:  1,
		PageSize: 20,
		Data:     []T{},
	}

	for _, option := range options {
		option(&opts)
	}
	// 如果没有分页，那么进行初始化配置
	pg := NewPageablePg(opts.Total, opts.PageNum, opts.PageSize)
	opts.SetPageable(pg)

	return opts
}
