package pathutils

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/daqiancode/gocommons/utils/randutils"
)

func MakeDateRandPath(ext string) string {
	now := time.Now()
	monthStr := fmt.Sprintf("%04d%02d", now.Year(), now.Month())
	return filepath.Join(monthStr, MakeRandPath(ext))
}

func MakeIdPath(id, ext string) string {
	if len(ext) > 0 && ext[0] != '.' {
		ext = "." + ext
	}
	if len(id) <= 3 {
		id = fmt.Sprintf("%04s", id)
	}
	return filepath.Join(id[:2], id[2:4], id+ext)
}

func MakeRandPath(ext string) string {
	randStr := randutils.RandomAlphabetNumberLower(4)
	name := randutils.RandomAlphabetNumberLower(20)
	if len(ext) > 0 && ext[0] != '.' {
		ext = "." + ext
	}
	return filepath.Join(randStr[:2], randStr[2:], name+ext)
}
