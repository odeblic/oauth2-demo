package main

import (
	"log"
	"net/http"
	"sync"
)

type Client struct {
	Secret string
	Code   string
}

const SALT = "8342"

// var mu sync.Mutex
var SECRET_KEY = []byte("0123456789")
var clientList sync.Map

func getClient(clientID string) *Client {
	object, present := clientList.Load(clientID)
	if present {
		return object.(*Client)
	}
	return nil
}

func main() {
	clientList.Store("app-0", &Client{"secret-000", ""})
	clientList.Store("app-1", &Client{"secret-111", ""})
	clientList.Store("app-2", &Client{"secret-222", ""})
	clientList.Store("app-3", &Client{"secret-333", ""})

	http.HandleFunc("/consent", consentHandler)
	http.HandleFunc("/authorize", authorizeHandler)
	http.HandleFunc("/token", tokenHandler)

	log.Println("OAuth2 Server running on https://localhost:5002")
	log.Fatal(http.ListenAndServeTLS(":5002", "cert.pem", "key.pem", nil))
}
