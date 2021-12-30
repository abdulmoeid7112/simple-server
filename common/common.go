package common

import (
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	Value string
	Count int
}

// WordCount - count the occurrence of words.
func WordCount(str string) interface{} {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}

	countStruct := make([]wordCount, 0)
	for key, val := range counts {
		countStruct = append(countStruct, wordCount{
			Value: key,
			Count: val,
		})
	}

	sort.SliceStable(countStruct, func(i, j int) bool {
		return countStruct[i].Count > countStruct[j].Count
	})

	if len(counts) > 10 {
		return countStruct[:10]
	}

	return countStruct
}

// RemovePunctuations - remove punctuations from string
func RemovePunctuations(textString string) (string, error) {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		return "", err
	}

	return re.ReplaceAllString(textString, " "), nil
}
