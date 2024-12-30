package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

type wordstr struct {
	Word string
	freq int
}

func onlstr(str string) string {
	var result []rune
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

func Top10(a string) []string {
	a = strings.ToLower(a)
	a = onlstr(a)

	words := strings.Fields(a)
	wordCount := make(map[string]int)

	for i := 0; i < len(words); i++ {
		if _, ok := wordCount[words[i]]; !ok {
			wordCount[words[i]] = 1
		} else {
			wordCount[words[i]]++
		}
	}

	wordstruct := []wordstr{}

	for wrd, num := range wordCount {
		neword := wordstr{Word: wrd, freq: num}
		wordstruct = append(wordstruct, neword)
	}

	sort.Slice(wordstruct, func(i, j int) bool {
		if wordstruct[i].freq != wordstruct[j].freq {
			return wordstruct[i].freq > wordstruct[j].freq
		}

		return wordstruct[i].Word < wordstruct[j].Word
	})

	otv := []string{}

	for i := 0; i < randeotv(len(wordstruct)); i++ {
		otv = append(otv, wordstruct[i].Word)
	}
	return (otv)
}

func randeotv(a int) int {
	if a > 10 {
		return 10
	}
	return a
}
