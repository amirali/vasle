package main

import (
	"fmt"
	"log"
	"net/http"
)

func runServer(port int, key string) {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/update", UpdateStatusHandler)

	go func() {
		select {
		case <-serverTicker.C:
			status = false
		}
	}()

	host := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(host, nil))
}
