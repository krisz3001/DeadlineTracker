package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	rand.Seed(time.Now().UnixNano())
	tpl = template.Must(template.ParseGlob("../*.html"))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewSessionId() string {
	b := make([]rune, 20)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func SendResponse(w http.ResponseWriter, i any, wrapper ...string) {
	data, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(wrapper) > 0 {
		data = append([]byte("{\""+wrapper[0]+"\":"), data...)
		data = append(data, []byte("}")...)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func DecodeRequest(w http.ResponseWriter, r *http.Request, i any) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

func IsAuthorized(r *http.Request) (string, int) {
	token, err := r.Cookie("Token")
	if err == http.ErrNoCookie {
		return "", 0
	}
	result := db.QueryRow("SELECT `Level` FROM SESSIONS LEFT JOIN USERS ON users.UserId = sessions.UserId WHERE `Token`=?", token.Value)
	var level int
	err = result.Scan(&level)
	if err != nil {
		return "", 0
	}
	return token.Value, level
}

func Home(w http.ResponseWriter, r *http.Request) {
	token, level := IsAuthorized(r)
	if token == "" {
		tpl.ExecuteTemplate(w, "login.html", nil)
		return
	}
	_ = token
	_ = level
	tpl.ExecuteTemplate(w, "index.html", nil)
}
func Controller_Assets(prefix string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _ := IsAuthorized(r)
		if token == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		} else {
			p := strings.TrimPrefix(r.URL.Path, prefix)
			rp := strings.TrimPrefix(r.URL.RawPath, prefix)
			if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
				r2 := new(http.Request)
				*r2 = *r
				r2.URL = new(url.URL)
				*r2.URL = *r.URL
				r2.URL.Path = p
				r2.URL.RawPath = rp
				h.ServeHTTP(w, r2)
			} else {
				http.NotFound(w, r)
			}
		}
	})
}

var db *sql.DB

func main() {
	fmt.Println("Welcome to DeadlineTracker v1.0!")
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(35.239.48.58:3306)/DEADLINETRACKER")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected to database.")
	mux := mux.NewRouter()

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-Content-Type-Options"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	mux.HandleFunc("/", Home).Methods("GET")
	mux.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	mux.PathPrefix("/assets/").Handler(Controller_Assets("/assets/", http.FileServer(http.Dir("assets/"))))
	mux.HandleFunc("/signup", Controller_Signup).Methods("POST")
	mux.HandleFunc("/login", Controller_Login).Methods("POST")
	mux.HandleFunc("/logout", Controller_Logout).Methods("GET")
	mux.HandleFunc("/deadlines", Controller_Deadlines).Methods("GET", "POST")
	mux.HandleFunc("/deadlines/{id:[0-9]+}", Controller_Deadlines_Id).Methods("GET", "PATCH", "DELETE")
	mux.HandleFunc("/subjects", Controller_Subjects).Methods("GET", "POST")
	mux.HandleFunc("/subjects/{id:[0-9]+}", Controller_Subjects_Id).Methods("GET", "PATCH", "DELETE")
	mux.HandleFunc("/deadlinetypes", Controller_DeadlineTypes).Methods("GET", "POST")
	mux.HandleFunc("/deadlinetypes/{id:[0-9]+}", Controller_DeadlineTypes_Id).Methods("GET", "PATCH", "DELETE")

	http.ListenAndServe(":80", handlers.CORS(header, methods, origins)(mux))
}
