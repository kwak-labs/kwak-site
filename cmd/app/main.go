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
	Picture string    `json:"picture"`
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
				Title:   "Why you should invest into sifchain",
				Content: "Sifchain is good investment",
				Picture: "https://pbs.twimg.com/profile_banners/1291413585189433344/1633944414/1500x500",

				Date: time.Date(2023, time.March,
					18, 21, 34, 01, 0, time.UTC),
			},
			{
				Title:   "Why Osmosis sucks",
				Content: "Hello",
				Picture: "https://pbs.twimg.com/profile_banners/1357133940109103104/1648318625/1500x500",

				Date: time.Now(),
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
