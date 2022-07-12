package mr_test

import (
	"fmt"
	"testing"

	"github.com/daqiancode/gocommons/utils/mr"
)

type MyInt int

type structA struct {
	Age   int
	Sex   int
	Class MyInt
}

func makeArr() []structA {
	return []structA{{1, 0, 1}, {3, 0, 2}, {3, 1, 3}, {4, 1, 4}}
}
func New[T any]() *T {
	t := new(T)
	return t
}
func TestNew(t *testing.T) {
	fmt.Println(New[structA]())
}
func TestReduce(t *testing.T) {
	arr := makeArr()
	r := mr.Reduce(arr, 10, func(i int, v structA, cum int) int {
		return v.Age + cum
	})
	fmt.Println(r)
}

func TestGroupBy(t *testing.T) {
	arr := makeArr()
	r := mr.GroupBy(arr, func(v structA) MyInt { return v.Class })
	fmt.Println(r)
}

func TestMapTo(t *testing.T) {
	arr := makeArr()
	r := mr.MapTo(arr, func(i int, v structA) int { return v.Sex })
	fmt.Println(r)
}
