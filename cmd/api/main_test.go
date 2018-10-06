package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const payload = `
{
	"lang":"%s",
	"text":"%s"
}`

var api = Api{}

func init() {
	api.Initialize()
}

var translateWordToPigLatingTests = []struct {
	text         string
	expectedText string
}{
	{"quiet", "ietquay"},
	{"yellow", "ellowyay"},
	{"style", "ylestay"},
	{"Challenge the status quo", "Allengechay ethay atusstay oquay"},
	{"Roses are red, violets are blue", "Osesray areway edray, ioletsvay areway ueblay"},
	{"He is 2 years old", "Ehay isway 2 earsyay oldway"},
	{"The design is state-of-the-art", "Ethay esignday isway atestay-ofway-ethay-artway"},
	{"An off-campus apartment", "Anway offway-ampuscay apartmentway"},
}

func TestTranslateTextToPig(t *testing.T) {
	for _, test := range translateWordToPigLatingTests {
		t.Run(test.text, func(t *testing.T) {
			bytePayload := []byte(fmt.Sprintf(payload, "pig", test.text))
			req, _ := http.NewRequest("POST", "/translate", bytes.NewBuffer(bytePayload))
			response := executeRequest(req)

			assert.Equal(t, http.StatusOK, response.Code)
			var translation translateResponse
			b, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal(b, &translation)

			assert.Equal(t, test.expectedText, translation.Translation)
		})
	}
}

func TestInvalidTranslationReturnsBadRequest(t *testing.T) {
	bytePayload := []byte(fmt.Sprintf(payload, "es", "test"))
	req, _ := http.NewRequest("POST", "/translate", bytes.NewBuffer(bytePayload))
	response := executeRequest(req)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	var errorResponse ErrorResponse
	b, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(b, &errorResponse)

	assert.Equal(t, "Invalid language", errorResponse.Error)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	api.Router.ServeHTTP(recorder, req)

	return recorder
}

type translateResponse struct {
	Translation string `json:translation`
}
type ErrorResponse struct {
	Error string `json:error`
}
