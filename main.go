package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	args := argsCheck()
	input := args[0]
	stringI := readFile(input)
	//regex for fixing whitespace inside (up/low/cap, 2)

	regQuote := `\s\'\s*(.+?)\s*\'`
	stringI = regexp.MustCompile(regQuote).ReplaceAllString(stringI, " '$1'")
	//fmt.Println(stringI)

	//regPunc := `\s*(\.|\,|\!|\?\:|\;)([[:punct:]])*\s*`
	//stringI = regexp.MustCompile(regPunc).ReplaceAllStringFunc(stringI, addSpace)
	//fmt.Println(stringI)

	regVowelLower := `(a)\s([aeyiuoh]+)`
	stringI = regexp.MustCompile(regVowelLower).ReplaceAllString(stringI, "an $2")
	//fmt.Println(stringI)

	regVowelUpper := `(A)\s([aeyiuoh]+)`
	stringI = regexp.MustCompile(regVowelUpper).ReplaceAllString(stringI, "An $2")
	//fmt.Println(stringI)

	regIns := `\((up|low|cap)(,\s*(\d*))\)`
	stringI = regexp.MustCompile(regIns).ReplaceAllStringFunc(stringI, replaceWhiteSpace)
	//split the string into array

	arrayInput := strings.Split(stringI, " ")
	var arrayOutput []string
	for i, e := range arrayInput {
		if strings.Contains(e, "(hex)") {
			hexStr, err := strconv.ParseInt(arrayInput[i-1], 16, 64)
			if err != nil {
				os.Exit(1)
			}
			arrayInput[i-1] = strconv.Itoa(int(hexStr))
		}
		if strings.Contains(e, "(bin)") {
			binStr, err := strconv.ParseInt(arrayInput[i-1], 2, 64)
			if err != nil {
				os.Exit(1)
			}
			arrayInput[i-1] = strconv.Itoa(int(binStr))
		}

		var numOfInstance int
		var instance string
		reg1 := `\((up|low|cap)(\,)*(\d)*\)`
		if regexp.MustCompile(reg1).MatchString(e) {
			numOfInstance, instance = findNumOfInstance(reg1, e)
			switch {
			case instance == "up":
				if numOfInstance == 0 {
					arrayInput[i-1] = strings.ToUpper(arrayInput[i-1])
				} else {
					if numOfInstance > len(arrayInput[:i]) {
						fmt.Print("too big number")
						os.Exit(0)
					} else {
						for x := numOfInstance; x >= 1; x-- {
							arrayInput[i-x] = strings.ToUpper(arrayInput[i-x])
						}
					}
				}
			case instance == "low":
				if numOfInstance == 0 {
					arrayInput[i-1] = strings.ToLower(arrayInput[i-1])
				} else {
					if numOfInstance > len(arrayInput[:i]) {
						fmt.Print("too big number")
						os.Exit(0)
					} else {
						for x := numOfInstance; x >= 1; x-- {
							arrayInput[i-x] = strings.ToLower(arrayInput[i-x])
						}
					}
				}
			case instance == "cap":
				if numOfInstance == 0 {
					arrayInput[i-1] = strings.Title(arrayInput[i-1])
				} else {
					if numOfInstance > len(arrayInput[:i]) {
						fmt.Print("too big number")
						os.Exit(0)
					} else {
						for x := numOfInstance; x >= 1; x-- {
							arrayInput[i-x] = strings.Title(arrayInput[i-x])
						}
					}
				}
			}
		}

	}

	fmt.Println("array output =", arrayOutput)

	strOutput := strings.Join(arrayInput, " ")

	fmt.Println(strOutput)

	regBracket := `\(.+?\)`
	strOutput = regexp.MustCompile(regBracket).ReplaceAllString(strOutput, "")

	regPunc := `\s*(\.|\,|\!|\?\:|\;)([[:punct:]])*\s*`
	strOutput = regexp.MustCompile(regPunc).ReplaceAllStringFunc(strOutput, addSpace)

	fmt.Println(strOutput)
}

// matched, err := regexp.MatchString(`foo.*`, "seafood")
// fmt.Println(matched, err)

// matched, err := regexp.MatchString(`(up.*)`, "(up, 2)")
// fmt.Println(matched, err)
