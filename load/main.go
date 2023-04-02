package main

import (
	"log"
	"net/http"
)

func main() {
	for {
		err := load()
		if err != nil {
			log.Fatalf("we have an error: %v", err)
		}
	}
}

func load() error {
	client := http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8080/sort", nil)
	if err != nil {
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
