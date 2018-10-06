package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	translator "github.com/xabi93/pig-translator"
)

func main() {
	api := Api{}

	api.Initialize()
	api.Run(os.Getenv("API_ADDR"))
}

type Api struct {
	Router *mux.Router
}

func (api *Api) Initialize() {
	api.Router = mux.NewRouter()
	api.initRoutes()
}

func (api *Api) initRoutes() {
	api.Router.HandleFunc("/translate", api.translateFromEnglish).Methods("POST")
}

func (api *Api) Run(addr string) {
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, api.Router))
}

func (api *Api) translateFromEnglish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request translateRequest
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &request)

	translator, err := translator.NewTranslatorFromLang(request.Lang)

	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	responseJSON(w, http.StatusOK, map[string]string{"translation": translator.Translate(request.Text)})
}

func responseJSON(w http.ResponseWriter, httpStatus int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(response)
}

type translateRequest struct {
	Lang string `json:lang`
	Text string `json:text`
}
