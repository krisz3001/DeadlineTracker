package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string `json:"password" db:"password"`
	Username string `json:"username" db:"username"`
}

type Session struct {
	Username string    `json:"username" db:"username"`
	Expiry   time.Time `json:"expiry" db:"expiry"`
}

var sessions = map[string]Session{}

/* func (s Session) isExpired() bool {
	return s.Expiry.Before(time.Now())
} */

func DoesUsernameExist(c Credentials) (bool, error) {
	rows, err := db.Query("SELECT * FROM `USERS` WHERE `Username`=?", c.Username)
	if err != nil {
		return true, err
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

func Controller_Signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		creds := &Credentials{}
		err := json.NewDecoder(r.Body).Decode(creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if exists, err := DoesUsernameExist(*creds); exists {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Error(w, "user already exists with that username", http.StatusConflict)
			return
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

		if _, err = db.Query("INSERT INTO USERS(`Username`,`Password`) VALUES (?, ?)", creds.Username, string(hashedPassword)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func Controller_Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		creds := &Credentials{}
		err := json.NewDecoder(r.Body).Decode(creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result := db.QueryRow("SELECT `PASSWORD` FROM USERS WHERE `USERNAME`=?", creds.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		storedCreds := &Credentials{}
		err = result.Scan(&storedCreds.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}
		sessionToken := uuid.NewString()
		expiresAt := time.Now().Add(120 * time.Second)
		sessions[sessionToken] = Session{
			Username: creds.Username,
			Expiry:   expiresAt,
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: expiresAt,
		})
	}
}
