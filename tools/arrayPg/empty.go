package arrayPg

// IsEmpty 验证是否为空
func IsEmptyString(val []string) bool {
	if nil == val {
		return false
	}
	return 0 == len(val)
}

// IsNotEmpty 验证 不为空
func IsNotEmptyString(val []string) bool {
	return !IsEmptyString(val)
}

// IsNil 验证是 否为 nil
func IsNilString(val []string) bool {
	return nil == val
}

// IsNotNil 验证 不为 nil
func IsNotNilString(val []string) bool {
	return nil != val
}

// IsEmptyFloat32 验证是否为空
func IsEmptyFloat32(val []float32) bool {
	if nil == val {
		return false
	}
	return 0 == len(val)
}

// IsNotEmptyFloat32 验证 不为空
func IsNotEmptyFloat32(val []float32) bool {
	return !IsEmptyFloat32(val)
}

// IsNilFloat32 验证是 否为 nil
func IsNilFloat32(val []float32) bool {
	return nil == val
}

// IsNotNilFloat32 验证 不为 nil
func IsNotNilFloat32(val []float32) bool {
	return nil != val
}

// IsEmptyFloat64 验证是否为空
func IsEmptyFloat64(val []float64) bool {
	if nil == val {
		return false
	}
	return 0 == len(val)
}

// IsNotEmptyFloat64 验证 不为空
func IsNotEmptyFloat64(val []float64) bool {
	return !IsEmptyFloat64(val)
}

// IsNilFloat64 验证是 否为 nil
func IsNilFloat64(val []float64) bool {
	return nil == val
}

// IsNotNilFloat64 验证 不为 nil
func IsNotNilFloat64(val []float64) bool {
	return nil != val
}

// IsEmptyInt 验证是否为空
func IsEmptyInt(val []int) bool {
	if nil == val {
		return false
	}
	return 0 == len(val)
}

// IsNotEmptyInt 验证 不为空
func IsNotEmptyInt(val []int) bool {
	return !IsEmptyInt(val)
}

// IsNilInt 验证是 否为 nil
func IsNilInt(val []int) bool {
	return nil == val
}

// IsNotNilInt 验证 不为 nil
func IsNotNilInt(val []int) bool {
	return nil != val
}

// IsEmptyInt64 验证是否为空
func IsEmptyInt64(val []int64) bool {
	if nil == val {
		return false
	}
	return 0 == len(val)
}

// IsNotEmptyInt64 验证 不为空
func IsNotEmptyInt64(val []int64) bool {
	return !IsEmptyInt64(val)
}

// IsNilInt64 验证是 否为 nil
func IsNilInt64(val []int64) bool {
	return nil == val
}

// IsNotNilInt64 验证 不为 nil
func IsNotNilInt64(val []int64) bool {
	return nil != val
}

// IsEmptyInt32 验证是否为空
func IsEmptyInt32(val []int32) bool {
	if nil == val {
		return false
	}
	return 0 == len(val)
}

// IsNotEmptyInt32 验证 不为空
func IsNotEmptyInt32(val []int32) bool {
	return !IsEmptyInt32(val)
}

// IsNilInt32 验证是 否为 nil
func IsNilInt32(val []int32) bool {
	return nil == val
}

// IsNotNilInt32 验证 不为 nil
func IsNotNilInt32(val []int32) bool {
	return nil != val
}
