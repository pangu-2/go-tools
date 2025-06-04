package genericPg

type ID interface {
	int | int32 | uint64 | string | uint | uint32 | int64
}
