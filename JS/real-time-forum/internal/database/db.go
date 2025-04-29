package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	database, err := sql.Open("sqlite3", "internal/data/forum.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		"uuid" BLOB NOT NULL PRIMARY KEY UNIQUE,
		"username" TEXT NOT NULL UNIQUE,
		"lastname" TEXT NOT NULL,
		"firstname" TEXT NOT NULL,
		"age" INTEGER NOT NULL,
		"gender" TEXT NOT NULL,
		"password" BLOB NOT NULL,
		"email" TEXT NOT NULL UNIQUE
	);`
	_, err = database.Exec(createUserTable)
	if err != nil {
		log.Fatal(err)
	}

	createSessionTable := `
	CREATE TABLE IF NOT EXISTS sessions (
		"session_id" TXT NOT NULL PRIMARY KEY UNIQUE,
		"user_UUID" BLOB NOT NULL,
		"expires_at" TIMESTAMP,
		FOREIGN KEY("user_uuid") REFERENCES users("uuid") ON DELETE CASCADE
	);`

	_, err = database.Exec(createSessionTable)
	if err != nil {
		log.Fatal(err)
	}

	createImageTable := `
	CREATE TABLE IF NOT EXISTS images (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	imageURL TEXT UNIQUE
	)`
	_, err = database.Exec(createImageTable)
	if err != nil {
		log.Fatal(err)
	}

	// Create the comments table
	createTopicTable := `
    CREATE TABLE IF NOT EXISTS topics (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
       	title TEXT UNIQUE,
		content TEXT
    );`
	_, err = database.Exec(createTopicTable)
	if err != nil {
		log.Fatal(err)
	}

	// Create the posts table
	createPostTable := `
    CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        user_id BLOB,
		topic_id TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(topic_id) REFERENCES topic(id),
        FOREIGN KEY(user_id) REFERENCES users(uuid)
    );`
	_, err = database.Exec(createPostTable)
	if err != nil {
		log.Fatal(err)
	}

	// Create the comments table
	createCommentTable := `
    CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL,
        post_id INTEGER,
        user_id BLOB,
		likes BLOB,
		dislikes BLOB,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY(post_id) REFERENCES posts(id),
        FOREIGN KEY(user_id) REFERENCES users(uuid)
    );`
	_, err = database.Exec(createCommentTable)
	if err != nil {
		log.Fatal(err)
	}

	createMessageTable := `
	CREATE TABLE IF NOT EXISTS messages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	sender_id BLOB NOT NULL,
	receiver_id BLOB NOT NULL,
	at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(sender_id) REFERENCES users(uuid),
	FOREIGN KEY(receiver_id) REFERENCES users(uuid) 
	);`

	_, err = database.Exec(createMessageTable)
	if err != nil {
		log.Fatal(err)
	}

	return database
}
