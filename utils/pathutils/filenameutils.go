package pathutils

import "strings"

//
func TagFilename(file, tag string) string {
	if tag == "" {
		return file
	}
	i := strings.LastIndex(file, ".")
	return file[:i] + "_" + tag + file[i:]
}
