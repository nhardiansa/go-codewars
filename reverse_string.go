package main

import (
	"fmt"
	"strings"
)

func reverseStr(s string) string {
	var str = strings.Split(s, "")
	var result = make([]string, len(s))

	for i := (len(str) - 1); i >= 0; i-- {
		result = append(result, str[i])
	}

	return strings.Join(result, "")
}

func main() {
	var word = reverseStr("world")
	var word2 = reverseStr("word")
	fmt.Println(word)
	fmt.Println(word2)
}
