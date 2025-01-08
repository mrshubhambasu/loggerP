package handlers

import (
	"authservice/utils"
	"encoding/json"
	"net/http"
)

func LoginEP(w http.ResponseWriter, r *http.Request) {
	// Login handler
	var Cred struct {
		Username string
		Password string
	}

	err := json.NewDecoder(r.Body).Decode(&Cred)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if Cred.Username != "admin" || Cred.Password != "admin" {
		http.Error(w, "Invalid Creds", http.StatusBadRequest)
		return
	}

	// Gen JWT

	token, err := utils.GenerateJWT(Cred.Username)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return token to client who is logging in
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
