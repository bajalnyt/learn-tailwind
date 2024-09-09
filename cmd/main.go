package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /index", homePageHandler)

	// Serve all static resources
	http.Handle("/static/*",
		http.StripPrefix("/static/",
			http.FileServer(
				http.Dir("./static"),
			),
		),
	)

	// Return all posts
	http.HandleFunc("GET /posts", postsHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/src/index.html")
}
