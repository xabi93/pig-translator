package translator

import (
	"math"
	"strings"
	"unicode"
)

type wordTranslator interface {
	translateWord(string) string
}

const (
	sufix      string = "ay"
	vowelSufix string = "w" + sufix
)

type pigLatingWordTranslator struct {
}

func (t *pigLatingWordTranslator) translateWord(word string) string {
	firstLetter := []rune(word)[0]

	if t.isVowel(firstLetter) {
		return word + vowelSufix
	}

	firstVowel := t.getFirstVowelPosition(word)

	if firstLetter == unicode.ToUpper(firstLetter) {
		return strings.Title(word[firstVowel:]) + strings.ToLower(word[:firstVowel]) + sufix
	}
	return word[firstVowel:] + word[:firstVowel] + sufix
}

func (t *pigLatingWordTranslator) getFirstVowelPosition(word string) int {
	lowerStr := strings.ToLower(word)
	isOdd := len(lowerStr)%2 != 0
	middle := int(math.Ceil(float64(len(lowerStr)) / 2))
	startsQ := lowerStr[0] == 'q'

	for i, c := range lowerStr {
		if (t.isVowel(c) || ('y' == c && isOdd && i+1 == middle)) && !('u' == c && i == 1 && startsQ) {
			return i
		}
	}

	return 0
}

func (t *pigLatingWordTranslator) isVowel(character rune) bool {
	switch character {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}

	return false
}
