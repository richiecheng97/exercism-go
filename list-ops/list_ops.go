package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	s = s.Reverse()
	for _, v := range s {
		initial = fn(v, initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var res = make([]int, 0)
	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	var res = make([]int, 0)
	for _, v := range s {
		res = append(res, fn(v))
	}
	return res
}

func (s IntList) Reverse() IntList {
	ll := s.Length()
	for i := 0; i < ll/2; i++ {
		s[i], s[ll-1-i] = s[ll-1-i], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	var res = s
	for _, v := range lists {
		res = res.Append(v)
	}
	return res
}
