package simple_util

import "strings"

// try to convert string to given width
func FormatWidth(width int, str string, fill byte) string {
	for len(str) < width {
		str = str + string(fill)
	}
	for len(str) > width {
		strs := strings.Split(str, "")
		if len(strs) > width {
			str = strings.Join(strs[0:width-1], "")
		} else {
			str = strings.Join(strs[0:len(strs)-1], "")
		}
	}
	return str
}
