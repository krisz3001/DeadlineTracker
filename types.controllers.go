package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DeadlineTypes struct {
	DeadlineTypeId   int    `json:"deadlinetypeid"`
	DeadlineTypeName string `json:"deadlinetypename"`
}

type NewDeadlineType struct {
	DeadlineTypeName string `json:"deadlinetypename"`
}

func Controller_DeadlineTypes(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	if token == "" {
		token = CreateToken()
	}
	switch r.Method {
	case http.MethodGet:
		result, err := GetDeadlineTypes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, result, token, "data")
	case http.MethodPost:
		var request NewDeadlineType
		if !DecodeRequest(w, r, &request) {
			return
		}
		err := CreateDeadlineType(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{}, token)
	}
}

func Controller_DeadlineTypes_Id(w http.ResponseWriter, r *http.Request) {
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
		result, err := GetDeadlineType(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, result, token, "data")
	case http.MethodPatch:
		var request NewDeadlineType
		if !DecodeRequest(w, r, &request) {
			return
		}
		err := UpdateDeadlineType(request, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{}, token)
	case http.MethodDelete:
		err := DeleteDeadlineType(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{}, token)
	}
}
