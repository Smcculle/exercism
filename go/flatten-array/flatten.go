package flatten

func Flatten(xs []interface{}) []interface{} {
	result := []interface{}{} // required for nil test
	for _, x := range xs {
		if i, ok := x.([]interface{}); ok {
			result = append(result, Flatten(i)...)
		} else if x != nil {
			result = append(result, x)
		}
	}
	return result
}
