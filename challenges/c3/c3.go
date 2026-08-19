package main

import (
	"fmt"
	"strings"
	"regexp"
	"slices"
	"cmp"
)

func clearText(text string) string {
	reg, err := regexp.Compile(`[^a-zA-Z0-9]`)

	if err != nil {
		panic(err)
	}

	newText := reg.ReplaceAllString(text, " ")

	return newText
}

func analyseText(text string) (countWords map[string]int){
	textClear := clearText(text)
	words := strings.Split(textClear, " ")
	
	countWords = make(map[string]int)

	for _, word := range words {
		if(word == "") {
			continue
		}

		countWords[word]++
	}

	return
}

func getTheFiveCommonsWordsInMap(mapWords map[string]int) (string, int) {
	type keyVal struct {
		k string
		v int
	}

	var listKeyWords []keyVal

	for key, value := range mapWords {
		listKeyWords = append(listKeyWords, keyVal{key, value})
	}

	slices.SortFunc(listKeyWords, func(i, j keyVal) int {
		return cmp.Compare(j.v, i.v)
	})

	lengthList := len(listKeyWords)
	var repeatedWords string
	var numWords int

	for i := 0; i < 5 && i < lengthList; i++ {
		repeatedWords = repeatedWords + listKeyWords[i].k + ", "
		numWords++
	} 

	return repeatedWords, numWords
}

func main(){
	mapNumWordwsInText := analyseText(`Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since 1966, when designers at Letraset and James Mosley, the librarian at St Bride Printing Library in London, took a 1914 Cicero translation and scrambled it to make dummy text for Letraset's Body Type sheets. It has survived not only many decades, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised thanks to these sheets and more recently with desktop publishing software like Aldus PageMaker and Microsoft Word including versions of Lorem Ipsum.`)

	mostCommonWords, numWords := getTheFiveCommonsWordsInMap(mapNumWordwsInText)

	fmt.Printf("As %d palavras mais repetidas são: %s", numWords, mostCommonWords)
}