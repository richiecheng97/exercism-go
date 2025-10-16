package series

func All(n int, s string) []string {
	if n == 0 || n > len(s) {
		return nil
	}

	var res []string
	for i := 0; i <= len(s)-n; i++ {
		res = append(res, s[i:i+n])
	}
	return res
}

func UnsafeFirst(n int, s string) string {
	if n == 0 || n > len(s) {
		return ""
	}
	return All(n, s)[0]
}

func First(n int, s string) (first string, ok bool) {
	if n == 0 || n > len(s) {
		return "", false
	}
	return All(n, s)[0], true
}
