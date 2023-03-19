package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Blog struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}

type Blogs struct {
	Blogs []Blog
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, "site/index.html")
}

func ReturnBlogs(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/json")

	b, _ := json.Marshal(Blogs{
		Blogs: []Blog{
			{
				Title:   "Title",
				Content: "Hello",
				Date:    time.Now(),
			},
		},
	})

	w.Write(b)
}

func main() {
	// Home Page
	http.HandleFunc("/", Homepage)

	http.HandleFunc("/api/blogs", ReturnBlogs)

	// Server static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start listening for HTTP requests on port 8080
	log.Println("Listening on :8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
