package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

func main() {
	username := os.Args[1]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var events []Event
	if err = json.NewDecoder(res.Body).Decode(&events); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(events) == 0 {
		return
	}
	for _, e := range events {
		fmt.Printf("- %s on %s\n", e.Type, e.Repo.Name)
	}

}
