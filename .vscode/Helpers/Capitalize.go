package goreloaded

import (
	"strings"
	"unicode"
)

// to capitalize arr[i-1] if arr[i] = ("cap")
func Capitalize(s string) string {
	for i, v := range s {
		return string(unicode.ToUpper(v)) + strings.ToLower(s[i+1:])
	}
	return s
}
