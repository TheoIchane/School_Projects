package handlers

import (
	"net/http"
	"real-time-forum/internal/database"
	"real-time-forum/internal/variables"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	database.DeleteExpiredSessions()
	var users []*variables.User
	user := database.GetCurrentUser(r)

	posts, status := database.GetPosts()
	if status != 200 {
		RespondJSON(w, status, map[string]any{
			"error": "Internal Server Error",
		})
		return
	}
	data := map[string]any{
		"Title": "Welcome to RTF!!!",
		"Posts": posts,
	}
	data_user := []map[string]any{}
	if user != nil {
		users = database.GetOtherUsers(user)
		for _, u := range users {
			data_user = append(data_user, map[string]any{
				"User":        u,
				"IsConnected": database.IsUserConnected(u),
			})
		}
	}
	page := variables.Page{
		Title: "Home",
		User:  user,
		Data:  data,
	}
	RespondJSON(w, 200, page)
}
