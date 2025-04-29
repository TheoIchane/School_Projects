package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"real-time-forum/internal/variables"

	"github.com/gorilla/mux"
)

const port = ":8080"

func Server() {
	defer variables.A.DB.Close()
	variables.A.Router = mux.NewRouter().StrictSlash(true)
	variables.A.Router.PathPrefix("/internal/src/static/").Handler(http.StripPrefix("/internal/src/static/", http.FileServer(http.Dir("./internal/src/static/"))))
	variables.A.Router.HandleFunc("/", servHTML)
	variables.A.Router.HandleFunc("/api/home", HomeHandler)
	variables.A.Router.HandleFunc("/api/topics/", GetTopicsHandler).Methods("GET")
	variables.A.Router.HandleFunc("/api/post/", GetPostsHandler).Methods("GET")
	variables.A.Router.HandleFunc("/api/post/{id}", GetPostHandler).Methods("GET")
	variables.A.Router.HandleFunc("/api/post/", CreatePostHandler).Methods("POST")
	variables.A.Router.HandleFunc("/api/register", RegisterHandler).Methods("POST")
	variables.A.Router.HandleFunc("/api/login", LoginHandler).Methods("POST")
	variables.A.Router.HandleFunc("/api/logout", LogoutHandler).Methods("GET")
	variables.A.Router.HandleFunc("/api/user", GetCurrentUserHandler).Methods("GET")
	variables.A.Router.HandleFunc("/api/users", GetUsersHandler).Methods("GET")
	// variables.A.Router.HandleFunc("/post/{id}",EditPostHandler).Methods("PUT")
	// variables.A.Router.HandleFunc("/post/{id}",DeletePostHandler).Methods("DELETE")

	variables.A.Router.HandleFunc("/socket", WebSocketHandler)
	fmt.Printf("Server Started on http://localhost%s/\n", port)
	http.ListenAndServe(port, variables.A.Router)
}

func servHTML(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("./internal/src/index.html"))
	temp.Execute(w, nil)
}
