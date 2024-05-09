package goreloaded

//removefirstrune If the length is 1 or less, return an empty string
func RemoveFirstRune(s string) string {
	if len(s) <= 1 {
		return ""
	}
	return s[1:]
}
