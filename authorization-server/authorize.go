package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
)

func generateAuthorizationCode(clientID string) string {
	data := clientID + "-" + SALT
	hash := md5.Sum([]byte(data))
	hexa := hex.EncodeToString(hash[:])
	authorizationCode := hexa[:6]
	return authorizationCode
}

func authorizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	r.ParseForm()

	consent := r.FormValue("consent")
	responseType := r.FormValue("response_type")
	clientID := r.FormValue("client_id")
	redirectURI := r.FormValue("redirect_uri")
	scope := r.FormValue("scope")
	state := r.FormValue("state")

	fmt.Printf("\033[33m---------------\033[32m AUTHORIZE \033[33m---------------\033[0m\n")
	fmt.Printf("consent:            \033[35m%s\033[0m\n", consent)
	fmt.Printf("responseType:       \033[35m%s\033[0m\n", responseType)
	fmt.Printf("clientID:           \033[35m%s\033[0m\n", clientID)
	fmt.Printf("redirectURI:        \033[35m%s\033[0m\n", redirectURI)
	fmt.Printf("scope:              \033[35m%s\033[0m\n", scope)
	fmt.Printf("state:              \033[35m%s\033[0m\n", state)

	if consent != "yes" {
		http.Error(w, "Consent not given", http.StatusUnauthorized)
		return
	}

	if responseType != "code" {
		http.Error(w, "Invalid response type", http.StatusBadRequest)
		return
	}

	client := getClient(clientID)

	if client == nil {
		http.Error(w, "Unknown client", http.StatusUnauthorized)
		return
	}

	client.Code = generateAuthorizationCode(clientID)

	http.Redirect(w, r, fmt.Sprintf("%s?code=%s&state=%s", redirectURI, client.Code, state), http.StatusFound)
}
