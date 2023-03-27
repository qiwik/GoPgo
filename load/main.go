package main

import (
	"bytes"
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
	//vars := []string{"bubble_sort", "fusion_sort"}
	//i := rand.Intn(2)

	var requestBody bytes.Buffer
	requestBody.Write([]byte("bubble_sort"))

	client := http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8081/sort", &requestBody)
	if err != nil {
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
