package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"twitter/server"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	s := server.Server{
		TweetsRepository: &server.TweetsMemoryRepository{},
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	rateLimit := httprate.LimitByIP(10, 1*time.Minute)
	r.With(rateLimit).Post("/tweets", s.AddTweet)

	r.Get("/tweets", s.ListTweets)

	go addTweets()

	log.Fatal(http.ListenAndServe(":8080", r))
}

func addTweets() {
	for {
		time.Sleep(time.Millisecond * 100)

		addTweetPayload := server.Tweet{
			Message:  "ass",
			Location: "ass",
		}

		b, err := json.Marshal(addTweetPayload)
		if err != nil {
			panic(err)
		}

		resp, err := http.Post(""+"http://localhost:8080/tweets", "application/json", bytes.NewBuffer(b))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		respBody := server.Response{}

		if err := json.Unmarshal(body, &respBody); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("Added tweet %d\n", respBody.ID)
		}
	}
}
