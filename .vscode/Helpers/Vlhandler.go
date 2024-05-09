package goreloaded

//Vowel handler
func VlHandler(arr []string) []string {
	if len(arr) < 1 {
		return arr
	}
	var ResArr []string
	for _, v := range arr {
		if len(v) == 0 {
			continue
		}
		ResArr = append(ResArr, v)
	}
	arr = ResArr
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == "a" || arr[i] == "A" {
			if (arr[i+1][0]) == 'a' || (arr[i+1][0]) == 'e' || (arr[i+1][0]) == 'i' || (arr[i+1][0]) == 'o' || (arr[i+1][0]) == 'u' || (arr[i+1][0]) == 'h' || (arr[i+1][0]) == 'A' || (arr[i+1][0]) == 'E' || (arr[i+1][0]) == 'I' || (arr[i+1][0]) == 'O' || (arr[i+1][0]) == 'U' {
				arr[i] = arr[i] + "n"
			}
		}
	}
	return arr
}
