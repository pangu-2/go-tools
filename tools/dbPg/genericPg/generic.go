package genericPg

type ID interface {
	int | int32 | int64 | string | uint
}
