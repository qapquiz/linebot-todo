package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var channelToken = os.Getenv("CHANNEL_TOKEN")
var channelSecret = os.Getenv("CHANNEL_SECRERT")

func determineListenPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func handlerWebhook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/webhook", handlerWebhook)

	port, err := determineListenPort()
	if err != nil {
		log.Fatal("$PORT not set")
	}

	log.Fatal(http.ListenAndServe(port, nil))
}
