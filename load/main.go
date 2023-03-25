package main

import (
	"bytes"
	"log"
	"math/rand"
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
	vars := []string{"s", "f"}
	i := rand.Intn(2)

	var requestBody bytes.Buffer
	requestBody.Write([]byte(vars[i]))

	client := http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8080/sort", &requestBody)
	if err != nil {
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
