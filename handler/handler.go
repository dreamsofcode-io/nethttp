package handler

import "net/http"

func FindByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Handle get
	} else if r.Method == http.MethodPost {
		// Handle post
	} else if r.Method == http.MethodPut {
		// Handle put
	}
}

func GetLatest(w http.ResponseWriter, r *http.Request) {
}
