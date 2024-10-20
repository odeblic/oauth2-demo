package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type ConsentFields struct {
	ClientID    string
	Scope       string
	RedirectURI string
	State       string
	Code        string
}

func consentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	r.ParseForm()

	responseType := r.FormValue("response_type")
	clientID := r.FormValue("client_id")
	redirectURI := r.FormValue("redirect_uri")
	scope := r.FormValue("scope")
	state := r.FormValue("state")

	fmt.Printf("\033[33m----------------\033[32m CONSENT \033[33m----------------\033[0m\n")
	fmt.Printf("responseType:       \033[35m%s\033[0m\n", responseType)
	fmt.Printf("clientID:           \033[35m%s\033[0m\n", clientID)
	fmt.Printf("redirectURI:        \033[35m%s\033[0m\n", redirectURI)
	fmt.Printf("scope:              \033[35m%s\033[0m\n", scope)
	fmt.Printf("state:              \033[35m%s\033[0m\n", state)

	consentTemplate, err := template.ParseFiles("consent.html")

	if err != nil {
		message := "Could not parse template: " + err.Error()
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	data := ConsentFields{}
	data.Scope = scope
	data.State = state
	data.ClientID = clientID
	data.RedirectURI = redirectURI

	err = consentTemplate.Execute(w, data)

	if err != nil {
		message := "Could not execute template: " + err.Error()
		http.Error(w, message, http.StatusInternalServerError)
	}
}
