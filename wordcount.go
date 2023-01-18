package main

import (
	"fmt"
	"os"
)

func main() {
	line := os.Args[1:]
	wordsCount := len(line)
	fmt.Println(wordsCount)
}
