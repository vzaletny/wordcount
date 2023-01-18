package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	line := os.Args[1]
	wordsCount := len(strings.Fields(line))
	fmt.Println(wordsCount)
}
