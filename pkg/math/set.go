package math

import "reflect"

type Equal func(i1 interface{}, i2 interface{}) bool

func Intersect(slices1 interface{}, slices2 interface{}, efunc Equal) (set []interface{}) {
	av := reflect.ValueOf(slices1)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		if Contains(slices2, el, efunc) {
			set = append(set, el)
		}
	}
	return
}

func Contains(a interface{}, e interface{}, efunc Equal) bool {
	v := reflect.ValueOf(a)

	for i := 0; i < v.Len(); i++ {
		if efunc(v.Index(i).Interface(), e) {
			return true
		}
	}
	return false
}
