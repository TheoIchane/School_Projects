package main

import (
	"real-time-forum/internal/database"
	"real-time-forum/internal/handlers"
	"real-time-forum/internal/variables"
)

func main() {
	variables.A.DB = database.InitDB()
	defer variables.A.DB.Close()
	handlers.Server()
}