package stringset

import (
	"sort"
	"strconv"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set []string

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, v := range l {
		s.Add(v)
	}
	return s
}

func (s Set) String() string {
	if len(s) == 0 {
		return "{}"
	}
	cpy := make([]string, len(s))
	copy(cpy, s)
	sort.Strings(cpy)

	quoted := make([]string, len(cpy))
	for i, v := range cpy {
		quoted[i] = strconv.Quote(v)
	}
	return "{" + strings.Join(quoted, ", ") + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}

func (s *Set) Add(elem string) {
	if !s.Has(elem) {
		*s = append(*s, elem)
	}
}

func Subset(s1, s2 Set) bool {
	for _, v := range s1 {
		if !s2.Has(v) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for _, v := range s1 {
		if s2.Has(v) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	res := New()
	for _, v := range s1 {
		if s2.Has(v) {
			res.Add(v)
		}
	}
	return res
}

func Difference(s1, s2 Set) Set {
	res := New()
	for _, v := range s1 {
		if !s2.Has(v) {
			res.Add(v)
		}
	}
	return res
}

func Union(s1, s2 Set) Set {
	res := New()
	for _, v := range s1 {
		res.Add(v)
	}
	for _, v := range s2 {
		res.Add(v)
	}
	return res
}
