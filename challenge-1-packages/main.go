package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	fmt.Println(getRandomProverb())
}

func getRandomProverb() string {
	return proverbs.Random().Saying
}
