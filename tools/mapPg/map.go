package mapPg

// MapContains 判断key是否存在
func MapContains(src map[string]interface{}, key string) bool {
	if _, ok := src[key]; ok {
		return true
	}
	return false
}
