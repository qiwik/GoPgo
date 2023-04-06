package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type example struct {
	FirstField  string `json:"first"`
	SecondField string `json:"second"`
	ThirdField  string `json:"third"`
}

func main() {
	for {
		err := load()
		if err != nil {
			log.Fatalf("we have an error: %v", err)
		}
	}
}

func load() error {
	sl := createForBubble()
	b, err := json.Marshal(sl)
	if err != nil {
		return err
	}

	client := http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8080/sort", bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

func createForBubble() []example {
	newSl := make([]example, 0)

	for i := 0; i < 1000; i++ {
		newSl = append(newSl, example{
			FirstField:  randStringBytes(),
			SecondField: randStringBytes(),
			ThirdField:  randStringBytes(),
		})
	}
	return newSl
}

func randStringBytes() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
