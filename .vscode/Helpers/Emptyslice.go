package goreloaded

func EmptySlice(s []string) []string {
	new := []string{}
	for _, v := range s {
		if v != "" {
			new = append(new, v)
		}
	}
	return new
}
