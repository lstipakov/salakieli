package main

import (
	"log"
	"net/http"

	"github.com/lstipakov/salakieli/renderer"
)

func handler(w http.ResponseWriter, r *http.Request) {
	text := "salakieli"

	if err := r.ParseForm(); err == nil {
		val, ok := r.Form["text"]
		if ok {
			text = val[0]
		}
	}

	renderer.Render(text, w)
}

func main() {
	http.HandleFunc("/salakieli", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
