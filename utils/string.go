package utils

import "unicode/utf8"

func Strlen(s string) int {
	return utf8.RuneCountInString(s)
}
