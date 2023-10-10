package pagePg

import "gorm.io/gorm"

// PageCondition 分页查询条件
// @Description:
type PageCondition[T any] struct {
	Condition  *gorm.DB  //自定义条件
	PageOption Option[T] // 分页参数
}

// OptionPageCondition 分页查询条件
type OptionPageCondition[T any] func(c *PageCondition[T])

// NewOptionPageCondition 分页查询条件
func NewOptionPageCondition[T any](options ...OptionPageCondition[T]) (PaginatorPg[T], *gorm.DB) {
	opts := PageCondition[T]{}
	for _, option := range options {
		option(&opts)
	}
	pg := NewPaginatorPg[T](opts.PageOption)
	return pg, opts.Condition
}
