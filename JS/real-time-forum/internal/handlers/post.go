package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"real-time-forum/internal/database"
	"real-time-forum/internal/variables"
	"strconv"
)

type PostResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	TopicID int    `json:"topic_id"`
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	data, status := database.GetPosts()
	if status == 200 {
		RespondJSON(w, status, data)
	} else {
		RespondJSON(w, status, map[string]any{
			"error": "Internal Server Error",
		})
	}
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	id_str := path.Base(r.URL.Path)
	id, _ := strconv.Atoi(id_str)
	data := database.GetPostByID(id)
	if data == nil {
		RespondJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
	} else {
		RespondJSON(w, 200, data)
	}
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	user := database.GetCurrentUser(r)
	if user == nil {
		RespondJSON(w, http.StatusUnauthorized, map[string]any{
			"error":   "Unauthorized",
			"message": "User Not Connected",
		})
		return
	}
	resp := PostResponse{}
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		fmt.Println(err)
		RespondJSON(w, http.StatusBadRequest, map[string]any{
			"error": "Bad Request",
		})
		return
	}
	post := &variables.Post{
		Title:   resp.Title,
		Content: resp.Content,
		Author:  user,
		Topic:   database.GetTopicByID(resp.TopicID),
	}
	status := database.InsertPost(post)
	if status == 200 {

		RespondJSON(w, status, map[string]any{
			"message": "POST SUCESSFULY INSERTED",
		})
	} else {
		RespondJSON(w, status, map[string]any{
			"error": "Internal Server Error",
		})
	}
}

func GetTopicsHandler(w http.ResponseWriter, r *http.Request) {
	data, status := database.GetTopics()
	if status == 200 {
		RespondJSON(w, status, data)
	} else {
		RespondJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
	}
}
