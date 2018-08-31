package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := []Article{
			{
				Title: "API Listening",
				Body:  "API Server ver: 0.1",
			},
		}
		render.JSON(w, r, response)
	})
	fmt.Println("Starting server on port 3000")
	http.ListenAndServe(":3000", r)
}
