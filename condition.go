package utils

// If 自定的三元计算
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// ErrorCheck 错误检测
func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
	return
}
