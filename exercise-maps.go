package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	cmap := make(map[string]int)
	for _, v := range strings.Fields(s) {
		cmap[v]++
	}
	return cmap
}

func main() {
	fmt.Println(WordCount("I am learngin GO! I am a Japanese"))
}
