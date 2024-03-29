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
	_, level := IsAuthorized(r)
	if level < 1 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case http.MethodGet:
		result, err := GetDeadlineTypes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, result, "data")
	case http.MethodPost:
		if level < 2 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		var request NewDeadlineType
		if !DecodeRequest(w, r, &request) {
			return
		}
		if exists, err := DoesTypeExist(request); exists {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Error(w, "Type already exists with that name", http.StatusConflict)
			return
		}
		err := CreateDeadlineType(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{})
	}
}

func Controller_DeadlineTypes_Id(w http.ResponseWriter, r *http.Request) {
	_, level := IsAuthorized(r)
	if level < 1 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
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
		SendResponse(w, result, "data")
	case http.MethodPatch:
		if level < 2 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		var request NewDeadlineType
		if !DecodeRequest(w, r, &request) {
			return
		}
		err := UpdateDeadlineType(request, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{})
	case http.MethodDelete:
		if level < 2 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		err := DeleteDeadlineType(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{})
	}
}
