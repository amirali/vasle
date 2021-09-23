package main

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

func updateStatus(serverUrl string, key string) {
	data := url.Values{}
	data.Set("key", key)

	_, err := http.PostForm(serverUrl, data)
	if err != nil {
		log.Println(err)
	}
}

func runClient(url string, key string) {
	updateStatus(url, key)

	for {
		select {
		case <-time.Tick(10 * time.Second):
			updateStatus(url, key)
		}
	}
}
