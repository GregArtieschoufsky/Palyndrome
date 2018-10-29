package main

import (
	"fmt"
)

func palyndromeCheck(word string) bool {
	wordCopy := word
	for i, r := range word {
		fmt.Printf("%s", string(r))
		var s string
		s = string(wordCopy[len(wordCopy)-1-i])
		fmt.Printf(",%s\n", s)
		if string(r)!=string(wordCopy[len(wordCopy)-1-i]) {
			return false
		}
	}
	return true
}