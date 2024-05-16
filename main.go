package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"
)

type Server struct {
	redis *redis.Client
}

func (s *Server) Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			name, err := s.redis.Get(context.Background(), "name").Result()
			if err != nil {
				log.Println("Error:", err)
			}

			w.Write([]byte(name))

		case http.MethodPost:
			err := s.redis.Set(context.Background(), "name", r.URL.Query().Get("name"), 0).Err()
			if err != nil {
				log.Println("Error:", err)
			}
		}
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil))
}

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	s := Server{redisClient}
	s.Run()
}
