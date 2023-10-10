package slicePg

// ToSlice 把结构体对象 转换成想要的数组
func ToSlice[T any, K comparable](array []T, iteratee func(T) K) []K {
	result := make([]K, len(array))
	for _, item := range array {
		k := iteratee(item)
		result = append(result, k)
	}

	return result
}

// ToList 把结构体对象 转换成想要的数组
func ToList[T any, K comparable](array []T, iteratee func(T) K) []K {
	return ToSlice(array, iteratee)
}

// ToMap 把结构体 转换为 map
func ToMap[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K]V {
	result := make(map[K]V, len(array))
	for _, item := range array {
		k, v := iteratee(item)
		result[k] = v
	}

	return result
}

func Unique[T comparable](slice []T) []T {
	result := []T{}

	for i := 0; i < len(slice); i++ {
		v := slice[i]
		skip := true
		for j := range result {
			if v == result[j] {
				skip = false
				break
			}
		}
		if skip {
			result = append(result, v)
		}
	}

	return result
}

func Merge[T any](slices ...[]T) []T {
	result := make([]T, 0)

	for _, v := range slices {
		result = append(result, v...)
	}

	return result
}

// ToMapArray 把结构体 转换为 map 数组
func ToMapArray[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K][]V {
	result := make(map[K][]V, len(array))
	for _, item := range array {
		k, v := iteratee(item)
		result[k] = append(result[k], v)
	}
	return result
}
