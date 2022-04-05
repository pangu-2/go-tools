package mapPg

// IsEmpty 验证是否为空
func IsEmptyMap(val map[string]interface{}) bool {
	if nil == val {
		return false
	}
	return 0 == len(val)
}

// IsNotEmpty 验证 不为空
func IsNotEmptyMap(val map[string]interface{}) bool {
	return !IsEmptyMap(val)
}

// IsNil 验证是 否为 nil
func IsNilMap(val map[string]interface{}) bool {
	return nil == val
}

// IsNotNil 验证 不为 nil
func IsNotNilMap(val map[string]interface{}) bool {
	return nil != val
}
