package stringutils

import (
	"errors"
	"strconv"
	"strings"
)

//ParseSizes 600x800,200x300, -> [[600,800],[200,300]],200->[[200,200]]
func ParseSizes(sizesStr string) ([][]int, error) {
	if sizesStr == "" {
		return nil, errors.New("ParseSize error: size string is empty")
	}
	sizesStrParts := strings.Split(sizesStr, ",")
	r := make([][]int, len(sizesStrParts))
	var err error
	for i, v := range sizesStrParts {
		r[i], err = ParseSize(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

//ParseSize 200x300 -> [200,300],200->[200,200]
func ParseSize(sizeStr string) ([]int, error) {
	if sizeStr == "" {
		return nil, errors.New("ParseSize error: size string is empty")
	}
	parts := strings.Split(strings.ToLower(sizeStr), "x")
	if len(parts) == 0 {
		return nil, nil
	}
	if len(parts) == 1 {
		s, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		return []int{s, s}, nil
	}

	w, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	h, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	return []int{w, h}, nil

}
