package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("capturing panic:", r)
		}
	}()

	// read file into bytes
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	// read text bytes and split into words
	words := strings.Fields(string(bytes))

	// prepare output with 4 fields: letters, symbols, numbers, words
	ret := make(map[string]int, 4)
	ret["words"] = len(words)

	for _, word := range words {
		for _, c := range word {
			switch {
			case unicode.IsNumber(c):
				ret["numbers"]++
			case unicode.IsLetter(c):
				ret["letters"]++
			case unicode.IsSymbol(c) || unicode.IsPunct(c):
				// punctuation like ? . ! is not part of symbols
				ret["symbols"]++
			default:
				fmt.Println("nani: ", string(c))
			}
		}
	}

	spew.Dump(ret)
}
