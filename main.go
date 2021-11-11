package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type NumberHandler struct {
	numbers []int
}

const JSONEncodingError = `{'error': 'json encoding error'}`

func (h NumberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	number := h.numbers[rand.Intn(len(h.numbers))]
	if err := json.NewEncoder(w).Encode(map[string]int{"number": number}); err != nil {
		http.Error(w, JSONEncodingError, http.StatusInternalServerError)
	}
}

type LetterHandler struct {
	letters []byte
}

func (h LetterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	letter := h.letters[rand.Intn(len(h.letters))]
	if err := json.NewEncoder(w).Encode(map[string]string{"letter": string([]byte{letter})}); err != nil {
		http.Error(w, JSONEncodingError, http.StatusInternalServerError)
	}
}

type WordHandler struct {
	words []string
}

func (h WordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	word := h.words[rand.Intn(len(h.words))]
	if err := json.NewEncoder(w).Encode(map[string]string{"word": word}); err != nil {
		http.Error(w, JSONEncodingError, http.StatusInternalServerError)
	}
}

func main() {
	nHandler := &NumberHandler{[]int{1, 2, 3, 4}}
	lHandler := &LetterHandler{[]byte("abcdef")}
	wHandler := &WordHandler{[]string{"hello", "world", "me"}}
	http.Handle("/api/number", nHandler)
	http.Handle("/api/letter", lHandler)
	http.Handle("/api/word", wHandler)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	log.Printf("Started listening at %v:%v", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
