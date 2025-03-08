package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		log.Println(r.Method)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Executado com sucesso"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Hello hello"))
	})

	log.Println("Running on port 8080")
	http.ListenAndServe(":8080", nil)
}
