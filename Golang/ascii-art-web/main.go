package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"web/handlers"
)

const port = ":8080"

func main() {
	http.Handle("/templates/stylesheets/", http.StripPrefix("/templates/stylesheets/", http.FileServer(http.Dir("./templates/stylesheets/"))))
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/download", handlers.Download)
	http.HandleFunc("/", handlers.Error404)
	// fmt.Println("Listening... (http://localhost:8080/)")
	// http.ListenAndServe(port, nil)
	srv := &http.Server{
		Addr:              port,              //adresse du server (le port choisi est à titre d'exemple)
		Handler:           nil,               // listes des handlers
		ReadHeaderTimeout: 10 * time.Second,  // temps autorisé pour lire les headers
		WriteTimeout:      10 * time.Second,  // temps maximum d'écriture de la réponse
		IdleTimeout:       120 * time.Second, // temps maximum entre deux rêquetes
		MaxHeaderBytes:    1 << 20,           // 1 MB // maxinmum de bytes que le serveur va lire
	}
	fmt.Println("http://localhost:8080/home")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
