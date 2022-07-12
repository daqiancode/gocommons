package stringutils

import (
	"strconv"

	"strings"
)

func Ints(ss []string) []int {
	r := make([]int, len(ss))
	for i, v := range ss {
		r[i], _ = strconv.Atoi(v)
	}
	return r
}

func IntsDefault(s string, dv ...int) int {
	v, err := strconv.Atoi(s)
	if nil != err {
		if len(dv) > 0 {
			return dv[0]
		}
		return 0
	}
	return v
}
func Strings(ints []int) []string {
	r := make([]string, len(ints))
	for i, v := range ints {
		r[i] = strconv.Itoa(v)
	}
	return r
}

//FileNameAppend hello.jpg,_1 -> hello_1.jpg
func FileNameAppend(filename, subname string) string {
	i := strings.LastIndex(filename, ".")
	if i == -1 {
		return filename + subname
	}
	return filename[0:i] + subname + filename[i:]
}
