package arrayPg

// IsEmpty 验证是否为空
func IsEmpty(val []string) bool {
	if nil == val {
		return false
	}
	return 0 == len(val)
}

// IsNotEmpty 验证 不为空
func IsNotEmpty(val []string) bool {
	return !IsEmpty(val)
}

// IsNil 验证是 否为 nil
func IsNil(val []string) bool {
	return nil == val
}

// IsNotNil 验证 不为 nil
func IsNotNil(val []string) bool {
	return nil != val
}
