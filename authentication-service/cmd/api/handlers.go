package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) ValidateToken(w http.ResponseWriter, r *http.Request) {
	// Extract the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		app.errorJSON(w, errors.New("authorization header is required"), http.StatusUnauthorized)
	}

	// Check if the token is in "Bearer <token>" format
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		app.errorJSON(w, errors.New("invalid authorization format"), http.StatusUnauthorized)
	}

	token := parts[1]

	// Validate the token
	if !app.verifyToken(token) {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	payload := jsonResponse{
		Error:   false,
		Message: "token is valid",
		Data:    nil,
	}
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) verifyToken(token string) bool {
	// TODO: validate the token
	return token == "valid-token"
}
