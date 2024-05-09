package goreloaded

//punctuation handler

func PunctuationsHandler(arr []string) []string {
	i := 1
	for i != len(arr) {
		if arr[i] != "" {
			if string(arr[i][0]) == "," || string(arr[i][0]) == "." || string(arr[i][0]) == ":" || string(arr[i][0]) == "!" || string(arr[i][0]) == "?" || string(arr[i][0]) == ";" {
				arr[i-1] = arr[i-1] + string(arr[i][0])
				arr[i] = RemoveFirstRune(arr[i])
				continue
			}
		}
		i++
	}
	return arr
}
