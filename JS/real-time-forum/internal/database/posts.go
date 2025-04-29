package database

import (
	"fmt"
	"net/http"
	"real-time-forum/internal/variables"
	"time"
)

func GetPostByID(id int) *variables.Post {
	post := &variables.Post{}
	rows, err := variables.A.DB.Query(`SELECT * FROM posts WHERE id = ?`, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		user_id := []byte{}
		var topic_id int
		var date time.Time
		rows.Scan(&post.ID, &post.Title, &post.Content, &user_id, &topic_id, &date)
		post.Date = date.Format("Mon 2 Jan 15:04")
		post.Author = GetUserByID(user_id)
		post.Topic = GetTopicByID(topic_id)
	}
	if post.ID == 0 {
		return nil
	}
	return post
}

func GetPosts() ([]*variables.Post, int) {
	posts := []*variables.Post{}
	rows, err := variables.A.DB.Query(`SELECT id FROM posts ORDER BY created_at DESC`)
	if err != nil {
		fmt.Println(err)
		return nil, http.StatusInternalServerError
	}
	for rows.Next() {
		var post_id int
		rows.Scan(&post_id)
		post := GetPostByID(post_id)
		if err != nil {
			fmt.Println(err)
			return nil, http.StatusInternalServerError
		}
		posts = append(posts, post)
	}
	return posts, http.StatusOK
}

func InsertPost(post *variables.Post) int {
	_, err := variables.A.DB.Exec(`
	INSERT INTO posts (
		"title",
		"content",
		"user_id",
		"topic_id",
        "created_at"
	)
	VALUES(?,?,?,?,?)
	`, post.Title, post.Content, post.Author.ID, post.Topic.ID, time.Now())
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
