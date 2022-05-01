package math

import "sort"

func Intersect[T comparable](slices1 []T, slices2 []T) (set []T) {

	for _, v := range slices1 {
		if Contains(slices2, v) {
			set = append(set, v)
		}
	}
	return
}

func Contains[T comparable](a []T, e T) bool {
	for _, v := range a {
		if v == e {
			return true
		}
	}
	return false
}

type UserLess[T comparable] func(t1 T, t2 T) bool
type MapKey[T comparable, TT comparable] func(t1 T) TT

type SortByT[T comparable] struct {
	data []T
	less UserLess[T]
}

func (s SortByT[T]) Len() int           { return len(s.data) }
func (s SortByT[T]) Swap(i, j int)      { s.data[i], s.data[j] = s.data[j], s.data[i] }
func (s SortByT[T]) Less(i, j int) bool { return s.less(s.data[i], s.data[j]) }

func Sort[T comparable](slices []T, less UserLess[T]) {
	sort.Sort(SortByT[T]{slices, less})
}

func GroupBy[T comparable, TT comparable](slices []T, less UserLess[T], mapkey MapKey[T, TT]) (groups map[TT][]T) {
	groups = make(map[TT][]T)
	Sort(slices, less)

	i := 0
	var j int
	for {
		if i >= len(slices) {
			break
		}
		for j = j + 1; j < len(slices) && !less(slices[i], slices[j]); j++ {

		}
		// groups = append(groups, slices[i:j])
		groups[mapkey(slices[i])] = slices[i:j]
		i = j
	}

	return
}
