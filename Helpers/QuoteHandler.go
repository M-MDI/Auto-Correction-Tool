package goreloaded

func QuoteHandler(text string) string {

	
		count := 1
		r := []rune(text)
		for i := 0; i < len(r); i++ {
			if count%2 != 0 && i < len(r)-1 {
				if r[i] == '\'' {
					if i < len(r)-1 && r[i+1] == ' ' {
						count++
					}
					if i+1 < len(r) && r[i+1] == ' ' {
						// count++
						r = append(r[:i+1], r[i+2:]...)
						continue
					} else {
						// count++
						continue
					}
				}
			}
			if count%2 == 0 && i > 0 && i < len(r) {
				if r[i] == '\'' {
					if i < len(r)-1 && r[i+1] == ' ' {
						count++
					}
					if i > 0 && r[i-1] == ' ' {
						// count++
						r = append(r[:i-1], r[i:]...)
					}
				}
			}
		}
		return string(r)
	
}
