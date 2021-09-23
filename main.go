package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

var configPath = os.Getenv("HOME") + "/.config/valse/"
var configFile = configPath + "config.toml"

type configSchema struct {
	Key    string `toml:"key"`
	Port   int    `toml:"port"`
	Client struct {
		Url string `toml:"url"`
	} `toml:"client"`
}

var c configSchema
var status bool

var serverTicker *time.Ticker

func init() {
	if _, err := os.Stat(configFile); err == nil {
		file, err := os.Open(filepath.Clean(configFile))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if _, err = toml.NewDecoder(file).Decode(&c); err != nil {
			log.Fatal(err)
		}
	} else {
		err := os.MkdirAll(configPath, 0755)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(filepath.Clean(configFile))
		if err != nil {
			log.Fatal(err)
		}

		c.Key = "0123456789abcdef"
		c.Port = 9090
		c.Client.Url = "http://localhost:9090/update"

		if err := toml.NewEncoder(file).Encode(c); err != nil {
			log.Fatal(err)
		}

		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	command := os.Args[1]

	switch command {
	case "server":
		serverTicker = time.NewTicker(15 * time.Second)
		runServer(c.Port, c.Key)

	case "client":
		runClient(c.Client.Url, c.Key)
	}
}
