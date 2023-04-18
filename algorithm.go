package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	myMap := make(map[string]int)
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")
	splt := strings.Split(text, " ")

	for _, word := range splt {
		myMap[word]++
	}
	fmt.Println(myMap)

	return myMap

}
