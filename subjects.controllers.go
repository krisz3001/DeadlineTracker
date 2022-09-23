package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Subject struct {
	SubjectKey  int    `json:"subjectkey"`
	SubjectName string `json:"subjectname"`
}

type NewSubject struct {
	SubjectName string `json:"subjectname"`
}

func Controller_Subjects(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result, err := GetSubjects()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, result, "data")
	case http.MethodPost:
		var request NewSubject
		if !DecodeRequest(w, r, &request) {
			return
		}
		if exists, err := DoesSubjectExist(request); exists {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Error(w, "subject already exists with that name", http.StatusConflict)
			return
		}
		err := CreateSubject(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{})
	}
}

func Controller_Subjects_Id(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		result, err := GetSubject(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, result, "data")
	case http.MethodPatch:
		var request NewSubject
		if !DecodeRequest(w, r, &request) {
			return
		}
		err := UpdateSubject(request, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{})
	case http.MethodDelete:
		err := DeleteSubject(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{})
	}
}
