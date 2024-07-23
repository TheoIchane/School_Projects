package main

import (
	"fmt"
	"groupie/handlers"
	"log"
	"net/http"
	"time"
)

const port = ":8080"

func main() {
	handlers.Data()
	http.Handle("/templates/stylesheets/", http.StripPrefix("/templates/stylesheets/", http.FileServer(http.Dir("./templates/stylesheets/"))))
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/artist/", handlers.Artist)
	srv := &http.Server{
		Addr:              port,              //adresse du server (le port choisi est à titre d'exemple)
		Handler:           nil,               // listes des handlers
		ReadHeaderTimeout: 10 * time.Second,  // temps autorisé pour lire les headers
		WriteTimeout:      10 * time.Second,  // temps maximum d'écriture de la réponse
		IdleTimeout:       120 * time.Second, // temps maximum entre deux rêquetes
		MaxHeaderBytes:    1 << 20,           // 1 MB // maxinmum de bytes que le serveur va lire
	}
	fmt.Println("http://localhost:8080/")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
