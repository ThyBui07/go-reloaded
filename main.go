package main

import (
	"fmt"
	"regexp"
)

func main() {
	args := argsCheck()
	input := args[0]
	a1 := readFile(input)

	//regex for fixing whitespace inside (up, 2)
	reg1 := `\((up|low|cap)(, \s*(\d*))\)`
	a1 = regexp.MustCompile(reg1).ReplaceAllStringFunc(a1, replaceWhiteSpace)
	// regex for punctuation
	//reg2 := `\s(\.|\,|\!|\?\:|\;)\s*`

	fmt.Println(a1)
	// arrayInput := strings.Split(a1, " ")
	// fmt.Println(arrayInput)
	// var arrayOutput []string
	// for i, e := range arrayInput {

	// 	if strings.Contains(e, "(up") {
	// 		fmt.Println("contain up at i", i)
	// 		arrayOutput[i-1] = strings.ToUpper(arrayOutput[i-1])
	// 	} else {
	// 		arrayOutput = append(arrayOutput, e)
	// 		fmt.Println("output rn = ", arrayOutput)
	// 	}
	// }

	//fmt.Println("array output =", arrayOutput)
	// output := args[1]
	//writeFile()
}

// matched, err := regexp.MatchString(`foo.*`, "seafood")
// fmt.Println(matched, err)

// matched, err := regexp.MatchString(`(up.*)`, "(up, 2)")
// fmt.Println(matched, err)
