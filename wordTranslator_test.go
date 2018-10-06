package translator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var translateWordToPigLatingTests = []struct {
	testTitle    string
	text         string
	expectedText string
}{
	{"starts with volwel", "is", "isway"},
	{"starts with consonant", "rose", "oseray"},
	{"starts with consonant 'y' in middle", "style", "ylestay"},
	{"starts with consonant 'u' as consonant", "question", "estionquay"},
	{"starts with consonant and capitalized", "Question", "Estionquay"},
	{"starts with consonant and no vowel letter", "chty", "chtyay"},
}

func TestPigLatinTranslatorGivenTextTranslatesIt(t *testing.T) {
	translator := pigLatingWordTranslator{}
	for _, test := range translateWordToPigLatingTests {
		t.Run(test.testTitle, func(t *testing.T) {
			translation := translator.translateWord(test.text)
			assert.Equal(t, test.expectedText, translation)
		})
	}
}

func TestGivenWordStartedWithVowel(t *testing.T) {
	translator := pigLatingWordTranslator{}

	assert.Equal(t, "isway", translator.translateWord("is"))
}

func TestGivenWordStartedWithConsonant(t *testing.T) {
	translator := pigLatingWordTranslator{}

	assert.Equal(t, "isway", translator.translateWord("is"))
}
