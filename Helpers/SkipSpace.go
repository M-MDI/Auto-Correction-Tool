package goreloaded

func SkipSpace(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		res += string(str[i])
		if str[i] == ' ' {
			for j := i; j < len(str); j++ {
				if str[i] == ' ' {
					i++
					continue
				} else {
					res += string(str[i])
					break
				}
			}
		}
	}
	return res
}
