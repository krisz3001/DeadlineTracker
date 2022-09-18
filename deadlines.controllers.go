package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Deadline struct {
	Id       int    `json:"id"`
	Subject  string `json:"subject"`
	Deadline string `json:"deadline"`
	Type     string `json:"type"`
	Topic    string `json:"topic"`
	Comments string `json:"comments"`
}

type NewDeadline struct {
	SubjectId int    `json:"subjectid"`
	Deadline  string `json:"deadline"`
	TypeId    int    `json:"typeid"`
	Topic     string `json:"topic"`
	Comments  string `json:"comments"`
}

func Controller_Deadlines(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")
	if token == "" {
		token = CreateToken()
	}
	switch r.Method {
	case http.MethodGet:
		result, err := GetDeadlines()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, result, token, "data")
	case http.MethodPost:
		var request NewDeadline
		if !DecodeRequest(w, r, &request) {
			return
		}
		err := CreateDeadline(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{}, token)
	}
}

func Controller_Deadlines_Id(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	if token == "" {
		token = CreateToken()
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		result, err := GetDeadline(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, result, token, "data")
	case http.MethodPatch:
		var request NewDeadline
		if !DecodeRequest(w, r, &request) {
			return
		}
		err := UpdateDeadline(request, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{}, token)
	case http.MethodDelete:
		err := DeleteDeadline(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{}, token)
	}
}
