package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		log.Println(r.Method)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Executado com sucesso"))
	})

	log.Println("Running on port 8080")
	http.ListenAndServe(":8080", nil)
}
