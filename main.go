package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewSessionId() string {
	b := make([]rune, 10)
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

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Middleware ", r.URL)
			next.ServeHTTP(w, r)
		})
}

var db *sql.DB

func main() {
	fmt.Println("Welcome to DeadlineTracker v1.0!")
	var err error
	db, err = sql.Open("mysql", "root:admin@tcp(localhost:3306)/DEADLINETRACKER")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected to database.")
	mux := mux.NewRouter()

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-Content-Type-Options"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	mux.HandleFunc("/signup", Controller_Signup).Methods("POST")
	mux.HandleFunc("/login", Controller_Login).Methods("POST")
	mux.HandleFunc("/deadlines", Controller_Deadlines).Methods("GET", "POST")
	mux.HandleFunc("/deadlines/{id:[0-9]+}", Controller_Deadlines_Id).Methods("GET", "PATCH", "DELETE")
	mux.HandleFunc("/subjects", Controller_Subjects).Methods("GET", "POST")
	mux.HandleFunc("/subjects/{id:[0-9]+}", Controller_Subjects_Id).Methods("GET", "PATCH", "DELETE")
	mux.HandleFunc("/deadlinetypes", Controller_DeadlineTypes).Methods("GET", "POST")
	mux.HandleFunc("/deadlinetypes/{id:[0-9]+}", Controller_DeadlineTypes_Id).Methods("GET", "PATCH", "DELETE")

	http.ListenAndServe(":3556", handlers.CORS(header, methods, origins)(mux))
}
