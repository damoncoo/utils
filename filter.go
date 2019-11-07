package utils

// Filter 筛选数组
func Filter(slice []string, filter func(string) bool) (ret []string) {
	for _, s := range slice {
		if filter(s) {
			ret = append(ret, s)
		}
	}
	return
}
