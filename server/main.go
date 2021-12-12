package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var RequestCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name:        "total_request_count",
})

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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}


func countingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RequestCount.Inc()
		next.ServeHTTP(w, r)
	})
}

func newHandler(next http.Handler) http.Handler {
	return countingMiddleware(loggingMiddleware(next))
}

func init() {
	prometheus.MustRegister(RequestCount)
}

func main() {
	nHandler := newHandler(NumberHandler{[]int{1, 2, 3, 4}})
	lHandler := newHandler(LetterHandler{[]byte("abcdef")})
	wHandler := newHandler(WordHandler{[]string{"hello", "world", "me"}})
	http.Handle("/api/number", nHandler)
	http.Handle("/api/letter", lHandler)
	http.Handle("/api/word", wHandler)
	http.Handle("/metrics", promhttp.Handler())

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	log.Printf("Started listening at %v:%v", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}