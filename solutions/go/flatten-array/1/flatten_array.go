package flatten

func Flatten(nested interface{}) []interface{} {
	var result = []interface{}{}
	switch t := nested.(type) {
	case []interface{}:
		for _, v := range t {
			flattened := Flatten(v)
			result = append(result, flattened...)
		}
	default:
		if t != nil {
			result = append(result, t)
		}
	}
	return result
}
