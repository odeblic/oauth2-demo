package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func generateAccessToken(clientID string, scope string) (string, error) {
	claims := jwt.MapClaims{
		"client_id": clientID,
		"scope":     scope,
		"exp":       time.Now().Add(time.Second * 60).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(SECRET_KEY)
	return accessToken, err
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	grantType := r.FormValue("grant_type")
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")
	authorizationCode := r.FormValue("authorization_code")

	fmt.Printf("\033[33m-----------------\033[32m TOKEN \033[33m-----------------\033[0m\n")
	fmt.Printf("grantType:          \033[35m%s\033[0m\n", grantType)
	fmt.Printf("clientID:           \033[35m%s\033[0m\n", clientID)
	fmt.Printf("clientSecret:       \033[35m%s\033[0m\n", clientSecret)
	fmt.Printf("authorizationCode:  \033[35m%s\033[0m\n", authorizationCode)

	if grantType != "authorization_code" {
		http.Error(w, "Invalid grant type", http.StatusBadRequest)
		return
	}

	client := getClient(clientID)

	if client == nil {
		http.Error(w, "Unknown client", http.StatusUnauthorized)
		return
	}

	if clientSecret != client.Secret {
		http.Error(w, "Invalid client credentials", http.StatusUnauthorized)
		return
	}

	if authorizationCode != client.Code {
		http.Error(w, "Invalid authorization code", http.StatusUnauthorized)
		return
	}

	accessToken, err := generateAccessToken(clientID, "all")

	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	client.Code = ""

	response := TokenResponse{
		AccessToken: accessToken,
		TokenType:   "bearer",
		ExpiresIn:   60,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
