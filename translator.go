package translator

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

type translator struct {
	wordTranslator wordTranslator
}

func NewTranslatorFromLang(lang string) (*translator, error) {
	switch lang {
	case "pig":
		return &translator{
			wordTranslator: &pigLatingWordTranslator{},
		}, nil
	}

	return nil, errors.New("Invalid language")
}

func (t *translator) Translate(text string) string {
	var translated []string
	words := t.getWordsSplitted(text)
	for _, word := range words {
		if t.isAlphanumeric(word) {
			word = t.wordTranslator.translateWord(word)
		}
		translated = append(translated, word)
	}
	return strings.Join(translated, "")
}

func (t *translator) isAlphanumeric(str string) bool {
	return regexp.MustCompile(`^[A-Za-z]+$`).MatchString(str)
}

func (t *translator) getWordsSplitted(text string) []string {
	var splitedWords []string
	var carry strings.Builder

	for _, c := range text {
		if unicode.IsLetter(c) {
			carry.WriteRune(c)
		} else if carry.Len() == 0 {
			splitedWords = append(splitedWords, string(c))
		} else {
			splitedWords = append(splitedWords, carry.String(), string(c))
			carry.Reset()
		}
	}
	splitedWords = append(splitedWords, carry.String())

	return splitedWords
}
