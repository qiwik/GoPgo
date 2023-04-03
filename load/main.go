package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	s := seed()
	sl := createForBubble(s)

	for {
		err := load(sl)
		if err != nil {
			log.Fatalf("we have an error: %v", err)
		}
	}
}

func load(sl []int) error {
	//s := seed()
	//sl := createForBubble(s)

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

func createForBubble(s rand.Source) []int {
	newSl := make([]int, 0)

	for i := 0; i < 10000; i++ {
		n := rand.New(s).Int()
		newSl = append(newSl, n)
	}
	return newSl
}
