package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"real-time-forum/internal/database"
	"real-time-forum/internal/variables"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterResp struct {
	UserName  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	PassWord  string `json:"password"`
}

type LoginResp struct {
	Identifier string `json:"identifier"`
	PassWord   string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	logResp := LoginResp{}
	err := json.NewDecoder(r.Body).Decode(&logResp)
	if err != nil {
		fmt.Println(err)
		RespondJSON(w, http.StatusBadRequest, map[string]any{
			"error": "Bad Request",
		})
		return
	}
	user := &variables.User{}
	if strings.Contains(logResp.Identifier, "@") {
		user = database.GetUserByEmail(logResp.Identifier)
	} else {
		user = database.GetUserByUserName(logResp.Identifier)
	}
	if user.Email == "" {
		RespondJSON(w, http.StatusNotFound, map[string]any{
			"error": "User Not Found",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword(user.PassWord, []byte(logResp.PassWord))
	if err != nil {
		RespondJSON(w, http.StatusUnauthorized, map[string]any{
			"error": "Unauthorized",
		})
		return
	}
	setCookie(w, user)
	RespondJSON(w, 200, map[string]any{
		"message": "USER CONNECTED",
		"status":  200,
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	delCookie(w, r)
	database.DeleteSession([]byte(cookie.Value))
	RespondJSON(w, 200, "USER LOGGED OUT")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	regResp := RegisterResp{}
	err := json.NewDecoder(r.Body).Decode(&regResp)
	if err != nil {
		fmt.Println(err)
		RespondJSON(w, http.StatusBadRequest, map[string]any{
			"error": "Bad Request",
		})
		return
	}
	email, err := mail.ParseAddress(regResp.Email)
	if err != nil {
		fmt.Println(err)
		RespondJSON(w, http.StatusBadRequest, map[string]any{
			"error":   "Bad Request",
			"message": "Invalid Email Address",
		})
		return
	}
	user_id, _ := uuid.NewV6()
	hashed, _ := bcrypt.GenerateFromPassword([]byte(regResp.PassWord), 14)

	user := &variables.User{
		ID:        user_id.Bytes(),
		Email:     email.Address,
		UserName:  regResp.UserName,
		FirstName: regResp.FirstName,
		LastName:  regResp.LastName,
		Gender:    regResp.Gender,
		Age:       regResp.Age,
		PassWord:  hashed,
	}
	status := database.InsertUser(user)
	fmt.Println(status)
	if status == http.StatusOK {
		setCookie(w, user)
		RespondJSON(w, http.StatusOK, map[string]any{
			"message": "User Succesfuly Registered",
		})
	} else {
		RespondJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error",
		})
	}
}

func setCookie(w http.ResponseWriter, user *variables.User) int {
	session_token, _ := uuid.NewV4()
	expiration := time.Now().Add(3600 * time.Second)
	cookie := http.Cookie{
		Name:     "session",
		Value:    base64.StdEncoding.EncodeToString(session_token.Bytes()),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return database.InsertSession(session_token.Bytes(), user.ID, expiration)
}

func delCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	session_token, _ := base64.StdEncoding.DecodeString(cookie.Value)
	cookie.MaxAge = -1
	cookie.Value = ""
	database.DeleteSession(session_token)
	http.SetCookie(w, cookie)
}

func GetCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	user := database.GetCurrentUser(r)
	fmt.Println(user)
}
