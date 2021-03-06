package main

import (
	"fmt"
	"net/http"
	"os"
)

const dest = "https://github.com/gobridge/about-us/blob/master/README.md"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println(
		http.ListenAndServe(
			":"+port,
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					http.Redirect(w, r, dest, http.StatusMovedPermanently)
				},
			),
		),
	)
}
