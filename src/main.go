package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type URLShortener struct {
	redisClient *redis.Client
}

type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

var ctx = context.Background()

func main() {
	rand.Seed(time.Now().UnixNano())

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		log.Fatalf("REDIS_ADDR environment variable is not set")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	shortener := &URLShortener{redisClient: redisClient}

	r := mux.NewRouter()
	r.HandleFunc("/shorten", shortener.shortenURL).Methods("POST")
	r.HandleFunc("/{shortURL}", shortener.redirect).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func (s *URLShortener) shortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL()
	err := s.redisClient.Set(ctx, shortURL, req.LongURL, 0).Err()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	res := ShortenResponse{ShortURL: shortURL}
	json.NewEncoder(w).Encode(res)
}

func (s *URLShortener) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	longURL, err := s.redisClient.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func generateShortURL() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
