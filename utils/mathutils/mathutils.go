package mathutils

import "github.com/daqiancode/gocommons/utils/atoms"

func Max[T atoms.NumStr](a []T) T {
	m := a[0]
	for _, v := range a {
		if m > v {
			m = v
		}
	}
	return m
}
func Min[T atoms.NumStr](a []T) T {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}

func In[T atoms.Basic](n T, ints []T) bool {
	for _, v := range ints {
		if n == v {
			return true
		}
	}
	return false
}

func Unique[T atoms.NumStr](vs []T) []T {
	m := make(map[T]bool)
	var r []T
	for _, v := range vs {
		if !m[v] {
			r = append(r, v)
		}
		m[v] = true
	}
	return r
}

//UnionInts AuB
func Union[T atoms.NumStr](a, b []T) []T {
	if nil == a {
		return nil
	}
	if nil == b {
		return a
	}
	var r []T
	m := make(map[T]bool, len(a)+len(b))
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		m[v] = true
	}
	for v := range m {
		r = append(r, v)
	}
	return r
}

//SubstractInts A-B
func Substract[T atoms.NumStr](a, b []T) []T {
	if nil == a {
		return nil
	}
	if nil == b {
		return a
	}
	var r []T
	m := make(map[T]bool)
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if !m[v] {
			r = append(r, v)
		}
	}
	return r
}

//IntersectInts AnB
func Intersect[T atoms.NumStr](a, b []T) []T {
	if nil == a {
		return nil
	}
	if nil == b {
		return nil
	}
	var r []T
	am := make(map[T]bool)
	bm := make(map[T]bool)
	for _, v := range a {
		am[v] = true
	}
	for _, v := range b {
		bm[v] = true
	}
	small, big := am, bm
	if len(am) > len(bm) {
		small, big = bm, am
	}
	for v := range small {
		if big[v] {
			r = append(r, v)
		}
	}
	return r
}
