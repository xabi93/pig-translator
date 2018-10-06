package translator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

type MockedWordTranslator struct {
	mock.Mock
}

func (m *MockedWordTranslator) translateWord(s string) string {
	args := m.Called(s)
	return args.String(0)
}

func TestGivenAlfaNumericTextTranslates(t *testing.T) {
	testWordTranslator := new(MockedWordTranslator)
	str := "test"
	exptStr := "esttay"

	testWordTranslator.On("translateWord", str).Return(exptStr)

	trans := translator{wordTranslator: testWordTranslator}

	assert.Equal(t, exptStr, trans.Translate(str))

	testWordTranslator.AssertCalled(t, "translateWord", str)
}

func TestGivenNonAlfaNumericTextNotTranslates(t *testing.T) {
	testWordTranslator := new(MockedWordTranslator)
	str := "-"
	exptStr := "-"

	trans := translator{wordTranslator: testWordTranslator}

	assert.Equal(t, exptStr, trans.Translate(str))

	testWordTranslator.AssertNotCalled(t, "translateWord")
}

func TestGivenTextShouldTranslateCorrectly(t *testing.T) {
	testWordTranslator := new(MockedWordTranslator)
	str := "hello, word!"
	exptStr := "ellohay, ordway!"

	testWordTranslator.On("translateWord", "hello").Return("ellohay")
	testWordTranslator.On("translateWord", "word").Return("ordway")

	trans := translator{wordTranslator: testWordTranslator}

	assert.Equal(t, exptStr, trans.Translate(str))

	testWordTranslator.AssertNumberOfCalls(t, "translateWord", 2)
}

func TestNewTranslatorGivenValidLangReturnsTranslator(t *testing.T) {
	trans, err := NewTranslatorFromLang("pig")

	assert.Nil(t, err)
	assert.IsType(t, &pigLatingWordTranslator{}, trans.wordTranslator)

}

func TestNewTranslatorGivenInvalidLangReturnsError(t *testing.T) {
	trans, err := NewTranslatorFromLang("es")

	assert.Nil(t, trans)
	assert.EqualError(t, err, "Invalid language")

}
