package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username" db:"Username"`
	Password string `json:"password" db:"Password"`
	Secret   string
}

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

func NewSession(id int) (string, error) {
	Token := NewSessionId()
	if _, err := db.Query("INSERT INTO SESSIONS(`Token`,`UserId`) VALUES (?, ?)", Token, id); err != nil {
		return "", err
	}
	return Token, nil
}

func Authenticate(w http.ResponseWriter, creds Credentials) (int, string) {
	result := db.QueryRow("SELECT `UserId` FROM USERS WHERE `Username` = ?", creds.Username)
	var userid int
	err := result.Scan(&userid)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return 0, ""
		}
	}
	Token, err := NewSession(userid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 0, ""
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   Token,
		Expires: time.Now().AddDate(2, 0, 0),
	})
	return userid, Token
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
		if creds.Secret != "kecske" {
			http.Error(w, "invalid secret", http.StatusUnauthorized)
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
		userid, token := Authenticate(w, *creds)
		if userid == 0 {
			return
		}
		h, m, s := time.Now().Clock()
		fmt.Println("[" + fmt.Sprint(h) + ":" + fmt.Sprint(m) + ":" + fmt.Sprint(s) + "] " + creds.Username + " signed up. (UserId: " + fmt.Sprint(userid) + ", Token: " + token + ")")
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
		result := db.QueryRow("SELECT `Password` FROM USERS WHERE `USERNAME`=?", creds.Username)
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
			return
		}
		userid, token := Authenticate(w, *creds)
		if userid == 0 {
			return
		}
		h, m, s := time.Now().Clock()
		fmt.Println("[" + fmt.Sprint(h) + ":" + fmt.Sprint(m) + ":" + fmt.Sprint(s) + "] " + creds.Username + " logged in. (UserId: " + fmt.Sprint(userid) + ", Token: " + token + ")")
	}
}

func Controller_Logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		token, _ := IsAuthorized(r)
		result, err := db.Exec("DELETE FROM `SESSIONS` WHERE `Token`=?", token)
		if err != nil {
			return
		}
		rows, _ := result.RowsAffected()
		if rows == 0 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}
