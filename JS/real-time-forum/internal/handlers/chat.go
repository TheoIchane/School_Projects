package handlers

import (
	"net/http"
	"real-time-forum/internal/database"
	"real-time-forum/internal/variables"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Content  any    `json:"content"`
	Date     string `json:"date"`
	Type     string `json:"type"`
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	database.DeleteExpiredSessions()
	var users []*variables.User
	var online []*variables.User
	var offline []*variables.User
	user := database.GetCurrentUser(r)
	if user == nil {
		RespondJSON(w, http.StatusUnauthorized, map[string]any{
			"error": "Unauthorized",
		})
		return
	}
	users = database.GetOtherUsers(user)
	for _, u := range users {
		if database.IsUserConnected(u) {
			online = append(online, u)
		} else {
			offline = append(offline, u)
		}
	}
	RespondJSON(w, http.StatusOK, map[string]any{
		"message": "Users Successfuly Queried",
		"online":  online,
		"offline": offline,
	})
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {

	// if savedSocketReader == nil {
	// 	savedSocketReader = make([]*socketReader, 0)
	// }

}
