package goreloaded

import "strconv"

func IsValidHex(arr []string, i int) bool {
	
	if i-1 >= 0 {
	
		_, err := strconv.ParseInt(arr[i-1], 16, 64)
		return err == nil
	}
	return false
}
