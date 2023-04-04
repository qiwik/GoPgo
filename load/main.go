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
	//s := seed()
	//sl := createForBubble(s)

	for {
		err := load()
		if err != nil {
			log.Fatalf("we have an error: %v", err)
		}
	}
}

func load() error {
	// каждый раз содаем новый слайс, чтобы симулировать различные варианты приходящих данных
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

func createForBubble(s rand.Source) []int {
	newSl := make([]int, 0)

	// так как пузырек имеет О(n^2), то смысла создавать слайс с большим количеством элементов просто нет.
	// ограничимся 1000 элементов, что уже очень много для этой сортировки
	for i := 0; i < 1000; i++ {
		n := rand.New(s).Int()
		newSl = append(newSl, n)
	}
	return newSl
}
