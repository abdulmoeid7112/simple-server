package main

import (
	"encoding/json"
	"github.com/abdulmoeid7112/count-words/common"
	"io/ioutil"
	"log"
	"net/http"
)

func countWords(w http.ResponseWriter, req *http.Request) {
	log.Println("Counting Handler...")
	contentType := req.Header.Get("content-Type")

	if contentType == "text/plain" {
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)

			return
		}

		textString, err := common.RemovePunctuations(string(body))
		if err != nil {
			log.Printf("Error removing punctuation marks: %v", err)
			http.Error(w, "can't remove punctuations", http.StatusBadRequest)

			return
		}

		responsePayload, err := json.Marshal(common.WordCount(textString))
		if err != nil {
			log.Printf("Error marshaling response body: %v", err)
			http.Error(w, "can't sent response body", http.StatusBadRequest)

			return
		}

		w.Write(responsePayload)
		w.WriteHeader(http.StatusOK)

		return
	}

	jsonPayload, err := json.Marshal(map[string]interface{}{
		"message": "invalid data",
	})
	if err != nil {
		log.Printf("Error marshaling response body: %v", err)
		http.Error(w, "invalid data sent", http.StatusBadRequest)

		return
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(jsonPayload)
}

func home(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Simple Word Count Server...", http.StatusOK)

	return
}

func main() {
	http.HandleFunc("/calculate", countWords)
	http.HandleFunc("/", home)

	log.Println("starting server, localhost:8000")
	http.ListenAndServe(":8000", nil)
}
