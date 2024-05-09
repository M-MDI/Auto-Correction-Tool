package main

import (
	"fmt"
	helper "goreloaded/Helpers"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	green := "\033[1;32m"
	yellow := "\033[1;33m"
	red := "\033[1;31m"
	reset := "\033[0m"

	if len(os.Args) < 3 {
		fmt.Println(yellow + "Warning :" + reset + red + " Put the " + reset + green + " Output && input file " + reset + red + "and try again ")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	} else if string(content) == "" || helper.SkipSpace(string(content)) == " " {
		fmt.Println(yellow + "Warning :" + reset + red + " No text in your file ")
		return
	}

	res := ""
	for _, v := range content {
		if v == '\n' {
			res += " " + string(v) + " "
		} else {
			res += string(v)
		}
	}
	str := res
	arr := strings.Split(str, " ")

	boo := false
	for boo == false {
		boo = true
		arr = helper.EmptySlice(arr)
		for i := 0; i < len(arr); i++ {
			if arr[i] == "(cap)" || (len(arr[i]) > 4 && arr[i][:5] == "(cap)") {
				if i == 0 {
					fmt.Println(yellow + "Warning :" + reset + red + "you can't use" + reset + green + " (cap) " + reset + red + "at the beginning of a sentence ")
					return
				}
				arr[i-1] = helper.Capitalize(arr[i-1])
				if len(arr[i]) > 4 {
					arr[i] = arr[i][5:]
				} else {
					arr[i] = ""
				}
				boo = false
				break // continue
			} else if arr[i] == "(up)" || (len(arr[i]) > 4 && arr[i][:4] == "(up)") {

				if i == 0 {
					fmt.Println(yellow + "Warning :" + reset + red + "you can't use" + reset + green + " (up) " + reset + red + "at the beginning of a sentence ")
					return
				}
				arr[i-1] = strings.ToUpper(arr[i-1])
				if len(arr[i]) > 4 {
					arr[i] = arr[i][4:]
				} else {
					arr[i] = ""
				}
				boo = false
				break
				// continue
			} else if arr[i] == "(low)" || (len(arr[i]) > 4 && arr[i][:5] == "(low)") {
				if i == 0 {
					fmt.Println(yellow + "Warning :" + reset + red + "you can't use" + reset + green + " (low) " + reset + red + "at the beginning of a sentence ")
					return
				}
				arr[i-1] = strings.ToLower(arr[i-1])
				if len(arr[i]) > 4 {
					arr[i] = arr[i][5:]
				} else {
					arr[i] = ""
				}
				boo = false
				break // continue
			} else if arr[i] == "(hex)" || (len(arr[i]) > 4 && arr[i][:5] == "(hex)") {
				if i == 0 {
					fmt.Println(yellow + "Warning :" + reset + red + "you can't use" + reset + green + " (hex) " + reset + red + "at the beginning of a sentence ")
					return
				}
				if !helper.IsValidHex(arr, i) {
					fmt.Println(yellow + "Warning :" + reset + red + " put a valid" + reset + green + " (hex) " + reset + red + " and try again ")
					return
				}
				arr[i-1] = strconv.Itoa(helper.AtoiBase(strings.ToUpper(arr[i-1]), "0123456789ABCDEF"))
				if len(arr[i]) > 4 {
					arr[i] = arr[i][5:]
				} else {
					arr[i] = ""
				}
				boo = false
				break
				// continue
			} else if arr[i] == "(bin)" || (len(arr[i]) > 4 && arr[i][:5] == "(bin)") && helper.IsNumeric(arr[i-1]) {
				if i == 0 {
					fmt.Println(yellow + "Warning :" + reset + red + "you can't use" + reset + green + " (bin) " + reset + red + "at the beginning of a sentence ")
					return
				}
				if !helper.IsvalidBin(arr, i) {
					fmt.Println(yellow + "Warning :" + reset + red + " put a valid" + reset + green + " (bin) " + reset + red + " and try again ")
					return
				}

				arr[i-1] = strconv.Itoa(helper.AtoiBase(arr[i-1], "01"))
				if len(arr[i]) > 4 {
					arr[i] = arr[i][5:]
				} else {
					arr[i] = ""
				}
				boo = false
				break
				// continue
			} else if strings.ContainsAny(arr[i], ".!?:;") && i > 0 {
				prevWord := arr[i-1]
				if len(prevWord) == 0 {
					continue
				}
				if strings.ContainsAny(string(prevWord[len(prevWord)-1]), ".,!?:;") {
					for i := 0; i < len(arr); i++ {
						arr = helper.PunctuationsHandler(arr)
					}
					boo = true
					break

				}
			} else if arr[i] == "." {
				arr[i-1] = "."
				boo = false
				break
			} else if (strings.HasSuffix(arr[i], "(cap,") || strings.HasSuffix(arr[i], "(up,") || strings.HasSuffix(arr[i], "(low,")) && i+1 < len(arr) && helper.IfRightFlag(arr[i+1]) {
				if i == 0 {
					fmt.Println(yellow + "Warning : " + red + "you can't use a " + reset + green + "FLAG " + reset + red + "at the beginning of a sentence")
					boo = true
					break
				} else if string(arr[i+1][len(arr[i+1])-1]) != ")" {
					fmt.Println(yellow + "Warning : " + red + "put a valid" + reset + green + " FLAG " + reset + red + " and try again !¡")
					boo = true
					break

				}

				n := helper.ToCap(arr[i+1])
				operation := strings.TrimSuffix(arr[i], ",")

				for j := n; j > 0; j-- {
					if i-j >= 0 {
						switch operation {
						case "(cap":
							arr[i-j] = helper.Capitalize(arr[i-j])
							arr[i+1] = ""
							arr[i] = ""
						case "(low":
							arr[i-j] = strings.ToLower(arr[i-j])
							arr[i+1] = ""
							arr[i] = ""
						case "(up":
							arr[i-j] = strings.ToUpper(arr[i-j])
							arr[i+1] = ""
							arr[i] = ""
						}
					}
				}
				boo = false
				break
			}
		}
	}

	arr1 := helper.VlHandler(arr)
	arr1 = helper.PunctuationsHandler(arr1)
	tmp := []string{}
	for _, t := range arr1 {
		if len(t) != 0 {
			tmp = append(tmp, t)
		}
	}
	j := strings.Join(tmp, " ")
	j = helper.QuoteHandler(j)
	result := ""

	for i := 0; i < len(j); i++ {
		if i != 0 && j[i-1] == '\n' && j[i] == ' ' {
			continue
		}
		result += string(j[i])
	}

	err = os.WriteFile(outputFile, []byte(result), 0644)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(yellow+"Warning : "+red+"Corrected data successfully saved to result file "+reset+green+"%s to %s !! \n", inputFile, outputFile)
}
