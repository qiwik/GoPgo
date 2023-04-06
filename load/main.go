package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type example struct {
	FirstField  int `json:"first"`
	SecondField int `json:"second"`
	ThirdField  int `json:"third"`
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
	// каждый раз создаем новый слайс, чтобы симулировать различные варианты приходящих данных
	s := seed()
	sl := createForBubble(s)

	b, err := json.Marshal(sl)
	if err != nil {
		return err
	}

	client := http.Client{}
	request, err := http.NewRequest("POST", "http://localhost:8080/sort", bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

func seed() rand.Source {
	return rand.NewSource(time.Now().UnixNano())
}

func createForBubble(s rand.Source) []example {
	newSl := make([]example, 0)

	for i := 0; i < 1000; i++ {
		newSl = append(newSl, example{
			FirstField:  rand.New(s).Int(),
			SecondField: rand.New(s).Int(),
			ThirdField:  rand.New(s).Int(),
		})
	}
	return newSl
}
