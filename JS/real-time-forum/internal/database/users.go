package database

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"real-time-forum/internal/variables"
	"time"
)

func GetUserByID(id []byte) *variables.User {
	user := &variables.User{}
	rows, err := variables.A.DB.Query(`SELECT * FROM users WHERE uuid = ?`, id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		password := []byte{}
		rows.Scan(&user.ID, &user.UserName, &user.LastName, &user.FirstName, &user.Age, &user.Gender, &password, &user.Email)
	}
	return user
}

func GetUserByEmail(email string) *variables.User {
	user := &variables.User{}
	rows, _ := variables.A.DB.Query(`SELECT * FROM users WHERE email = ?`, email)
	for rows.Next() {
		rows.Scan(&user.ID, &user.UserName, &user.LastName, &user.FirstName, &user.Age, &user.Gender, &user.PassWord, &user.Email)
	}
	return user
}

func GetUserByUserName(userName string) *variables.User {
	user := &variables.User{}
	rows, _ := variables.A.DB.Query(`SELECT * FROM users WHERE username = ?`, userName)
	for rows.Next() {
		rows.Scan(&user.ID, &user.UserName, &user.LastName, &user.FirstName, &user.Age, &user.Gender, &user.PassWord, &user.Email)
	}
	return user
}

func InsertUser(user *variables.User) int {
	_, err := variables.A.DB.Exec(`
	INSERT INTO users
	VALUES(?,?,?,?,?,?,?,?)
	`, user.ID, user.UserName, user.LastName, user.FirstName, user.Age, user.Gender, user.PassWord, user.Email)
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func InsertSession(session_token []byte, user_id []byte, expiration time.Time) int {
	rows, _ := variables.A.DB.Query(`
	DELETE FROM sessions WHERE user_UUID = ?`, user_id)
	for rows.Next() {
		var session_id []byte
		rows.Scan(session_id)
		DeleteSession(session_id)
	}
	_, err := variables.A.DB.Exec(`
		INSERT INTO sessions
		VALUES(?,?,?)
	`, session_token, user_id, expiration)
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func DeleteSession(session_token []byte) {
	_, _ = variables.A.DB.Exec(`
		DELETE FROM sessions 
		WHERE session_id = ?;
	`, session_token)
}

func DeleteExpiredSessions() {
	_, _ = variables.A.DB.Exec(`
		DELETE FROM sessions 
		WHERE expires_at <= ?;
	`, time.Now())
}

func GetCurrentUser(r *http.Request) *variables.User {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil
	}
	decodedSession, _ := base64.StdEncoding.DecodeString(cookie.Value)
	var user *variables.User
	rows, _ := variables.A.DB.Query(`
	SELECT user_UUID FROM sessions WHERE session_id = ?`,
		decodedSession,
	)

	for rows.Next() {
		var userUUID []byte
		err := rows.Scan(&userUUID)
		if err != nil {
			return nil
		}
		user = GetUserByID(userUUID)
	}

	return user
}

func IsUserConnected(user *variables.User) bool {
	var user_id []byte
	err := variables.A.DB.QueryRow(`SELECT user_UUID from sessions WHERE user_UUID = ? and expires_at > ? `, user.ID, time.Now()).Scan(&user_id)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func GetOtherUsers(user *variables.User) []*variables.User {
	var users []*variables.User
	rows, err := variables.A.DB.Query(`SELECT uuid FROM users WHERE uuid != ?`, user.ID)
	if err != nil {
		return nil
	}
	for rows.Next() {
		var user_id []byte
		err := rows.Scan(&user_id)
		if err != nil {
			return nil
		}
		users = append(users, GetUserByID(user_id))
	}
	return users
}
