package collections

import "github.com/daqiancode/gocommons/utils/atoms"

func Unique[T atoms.JSON](vs []T) []T {
	m := make(map[T]bool, len(vs))
	var r []T
	for _, v := range vs {
		if m[v] {
			continue
		}
		r = append(r, v)
		m[v] = true

	}
	return r
}

func AsMap[T atoms.JSON](vs []T) map[T]bool {
	m := make(map[T]bool, len(vs))
	for _, v := range vs {
		m[v] = true
	}
	return m
}

func Subtract[T atoms.JSON](a []T, b []T) []T {
	if len(a) == 0 {
		return nil
	}
	m := AsMap(b)
	a = Unique(a)
	var r []T
	for _, v := range a {
		if !m[v] {
			r = append(r, v)
		}
	}
	return r
}

func Union[T atoms.JSON](vs ...[]T) []T {
	var r []T
	for _, v := range vs {
		r = append(r, v...)
	}
	return Unique(r)
}

func Intersection[T atoms.JSON](a, b []T) []T {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	m := AsMap(b)
	a = Unique(a)
	var r []T
	for _, v := range a {
		if m[v] {
			r = append(r, v)
		}
	}
	return r
}
