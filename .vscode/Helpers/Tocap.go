package goreloaded

import (
	"fmt"
	"os"
	"strconv"
)

// atoi to convert string into int and tak params in (up, / (low, / (cap,

func ToCap(s string) int {
	
	green := "\033[1;32m"
	yellow := "\033[1;33m"
	red := "\033[1;31m"
	reset := "\033[0m"

	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			str = "-"
		}
		if s[i] >= '0' && s[i] <= '9' {
			str += string(s[i])
		}
	}
	n, err := strconv.Atoi(str)
	if err != nil || n < 0 {
		fmt.Println(yellow + "Warning : " + red + "put a valid" + reset + green + " FLAG " + reset + red + " and try again !¡")
		os.Exit(0)
	}
	return n

}
