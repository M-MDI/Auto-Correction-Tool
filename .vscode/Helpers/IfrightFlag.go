package goreloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IfRightFlag(s string) bool {
	green := "\033[1;32m"
	yellow := "\033[1;33m"
	red := "\033[1;31m"
	reset := "\033[0m"

	if strings.HasSuffix(s, ")") {
		s = strings.TrimSuffix(s, ")")
		if strings.HasSuffix(s, ")") {
			fmt.Println(yellow + "Warning: " + red + "put a valid" + reset + green + " FLAG " + reset + red + " and try again!")
			os.Exit(0)
		}
	}else  {
		fmt.Println(yellow + "Warning: " + red + "put a valid" + reset + green + " FLAG " + reset + red + " and try again!")
		os.Exit(0)
	}

	_, err := strconv.Atoi(s)
	return err == nil	
}
