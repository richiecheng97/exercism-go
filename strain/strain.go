package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

type Array []any

func Keep[T int | string | []int](cp []T, f func(T) bool) []T {
	var res []T
	for _, v := range cp {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func Discard[T int | string | []int](cp []T, f func(T) bool) []T {
	var res []T
	for _, v := range cp {
		if !f(v) {
			res = append(res, v)
		}
	}
	return res
}
