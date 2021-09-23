package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatBool(status)))
}

func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/update" {
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	requestKey := r.FormValue("key")

	if requestKey != c.Key {
		http.Error(w, "403 Forbidden", http.StatusForbidden)
		return
	}

	status = true

	serverTicker.Reset(15 * time.Second)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
