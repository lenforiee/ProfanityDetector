package profanitydetector

import (
	"fmt"
	"strings"

	blacklistedWords "github.com/lenforiee/ProfanityDetector/blacklisted_words"
)

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

type ProfanityFilter struct {
	profanityWordList []string
}

func NewProfanityFilter() (filter ProfanityFilter) {
	filter.profanityWordList = blacklistedWords.WordList
	return filter
}

func NewCleanProfanityFilter() (filter ProfanityFilter) {
	filter.profanityWordList = []string{}
	return filter
}

func (f *ProfanityFilter) ClearList() {
	f.profanityWordList = []string{}
}

func (f *ProfanityFilter) AddList(list []string) {
	for _, word := range list {
		f.AddOne(word)
	}
}

func (f *ProfanityFilter) AddOne(word string) {
	idx := indexOf(word, f.profanityWordList)

	if idx != -1 {
		fmt.Println("ProfanityFilter:", fmt.Sprintf("Word %s is already in list!", word))
		return
	}

	f.profanityWordList = append(f.profanityWordList, word)
}

func (f *ProfanityFilter) RemoveList(list []string) {
	for _, word := range list {
		f.RemoveOne(word)
	}
}

func (f *ProfanityFilter) RemoveOne(word string) {
	idx := indexOf(word, f.profanityWordList)

	if idx == -1 {
		fmt.Println("ProfanityFilter:", fmt.Sprintf("Word %s not found in list!", word))
		return
	}

	f.profanityWordList = append(f.profanityWordList[:idx], f.profanityWordList[idx+1:]...)
}

func (f *ProfanityFilter) IsProfanity(str string, substring ...bool) bool {

	for _, word := range f.profanityWordList {
		strLower := strings.ToLower(str)
		wordLower := strings.ToLower(word)

		substringAllowed := len(substring) > 0 && substring[0] == true

		if strLower == wordLower {
			return true
		} else if substringAllowed && strings.Contains(strLower, wordLower) {
			return true
		} else {
			continue
		}
	}
	return false
}
