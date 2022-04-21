package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginSuccessResponse struct {
	Username string `json:"username"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")
	res, err := api.usersRepo.Login(username, password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	fmt.Println(res)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginSuccessResponse{Username: username})
}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	err := api.usersRepo.Logout(username)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(AuthErrorResponse{Error: ""})
}