package mr

import (
	"github.com/daqiancode/gocommons/utils/atoms"
)

func Map[T any](arr []T, fn func(i int, v T) T) []T {
	if arr == nil {
		return nil
	}
	r := make([]T, len(arr))
	for i, v := range arr {
		r[i] = fn(i, v)
	}
	return r
}

func MapTo[T any, E any](arr []T, fn func(i int, v T) E) []E {
	if arr == nil {
		return nil
	}
	r := make([]E, len(arr))
	for i, v := range arr {
		r[i] = fn(i, v)
	}
	return r
}

func Filter[T any](arr []T, fn func(i int, v T) bool) []T {
	if arr == nil {
		return nil
	}
	var r []T
	for i, v := range arr {
		if fn(i, v) {
			r = append(r, v)
		}
	}
	return r
}

func Reduce[T any, E atoms.Basic](arr []T, initValue E, fn func(i int, v T, cum E) E) E {
	cum := initValue
	for i, v := range arr {
		cum = fn(i, v, cum)
	}
	return cum
}

func GroupBy[T any, E atoms.Basic](arr []T, by func(v T) E) map[E][]T {
	r := make(map[E][]T)
	var key E
	for _, v := range arr {
		key = by(v)
		if b, ok := r[key]; ok {
			r[key] = append(b, v)
		} else {
			r[key] = []T{v}
		}
	}
	return r
}

func GroupByOne[T any, E atoms.Basic](arr []T, keepFirst bool, by func(v T) E) map[E]T {
	r := make(map[E]T)
	var key E
	for _, v := range arr {
		key = by(v)
		if _, ok := r[key]; ok {
			if !keepFirst {
				r[key] = v
			}
		}
	}
	return r
}

func IndexOf[T any](arr []T, fn func(v T) bool) int {
	for i, v := range arr {
		if fn(v) {
			return i
		}
	}
	return -1
}

//Any Returns true if arr contain any true value
func Any(arr []bool) bool {
	for _, v := range arr {
		if v {
			return true
		}
	}
	return false
}

//Any Returns true if all values are true
func All(arr []bool) bool {
	for _, v := range arr {
		if !v {
			return false
		}
	}
	return true
}

//Reindex 按indexes设置索引
func Reindex[T any](arr []T, indexes []int) {
	for i, v := range indexes {
		arr[i], arr[v] = arr[v], arr[i]
	}
}

//ValueIndex arr -> [value]index
func ValueIndex[T atoms.Basic](arr []T) map[T]int {
	r := make(map[T]int, len(arr))
	for i, v := range arr {
		r[v] = i
	}
	return r
}

// func MakeMap[K atoms.Basic, V any](kvs ...interface{}) map[K]V {
// 	r := make(map[K]V)
// 	for i := 0; i < len(kvs); i += 2 {
// 		r[kvs[i]] = kvs[i+1]
// 	}
// 	return r
// }

//Sort 排序结构体数组 。arr:[]T,fieldAsc : field true|false,...
// func Sort(arr interface{}, fieldAsc ...interface{}) {
// 	if nil == arr {
// 		return
// 	}
// 	arrV := reflect.ValueOf(arr)
// 	sort.Slice(arr, func(i, j int) bool {
// 		for k := 0; k < len(fieldAsc); k += 2 {
// 			asc := true
// 			if k < len(fieldAsc)-1 {
// 				asc = fieldAsc[k+1].(bool)
// 			}
// 			field := fieldAsc[k].(string)
// 			ei := arrV.Index(i).FieldByName(field).Interface()
// 			ej := arrV.Index(j).FieldByName(field).Interface()
// 			if ei == ej {
// 				continue
// 			}
// 			le, err := LE(ei, ej)
// 			if err != nil {
// 				panic(err)
// 			}
// 			return asc && le || !asc && !le
// 		}
// 		return false
// 	})
// }
// func SortLikeInts(arr interface{}, field string, orderedInts []int) {
// 	if nil == arr {
// 		return
// 	}
// 	im := IntsIndexMap(orderedInts)
// 	arrV := reflect.ValueOf(arr)
// 	sort.Slice(arr, func(i, j int) bool {
// 		ei := arrV.Index(i).FieldByName(field).Interface().(int)
// 		ej := arrV.Index(j).FieldByName(field).Interface().(int)
// 		i1, ok1 := im[ei]
// 		i2, ok2 := im[ej]
// 		if !ok1 {
// 			return false
// 		}
// 		if !ok2 {
// 			return true
// 		}
// 		return i1 < i2
// 	})

// }
