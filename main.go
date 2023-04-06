package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
)

type example struct {
	FirstField  string `json:"first"`
	SecondField string `json:"second"`
	ThirdField  string `json:"third"`
}

func main() {
	http.HandleFunc("/sort", sorting)
	log.Printf("Serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sorting(w http.ResponseWriter, r *http.Request) {
	var sl []example
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &sl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bubbleSort(sl)
	w.WriteHeader(http.StatusOK)
}

func bubbleSort(sl []example) {
	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j < len(sl); j++ {
			if compare(sl[i], sl[j]) {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}
}

func compare(left, right example) bool {
	comparing := firstFieldCompare(left.FirstField, right.FirstField)
	comparing = secondFieldCompare(left.SecondField, right.SecondField, comparing)
	if comparing > 1 {
		return true
	}
	return thirdFieldCompare(left.ThirdField, right.ThirdField, comparing)
}

func firstFieldCompare(left, right string) int {
	if strSum(left) > strSum(right) {
		return 1
	}

	return 0
}

func secondFieldCompare(left, right string, comparing int) int {
	if strSum(left) > strSum(right) {
		comparing++
	}

	return comparing
}

func thirdFieldCompare(left, right string, comparing int) bool {
	if strSum(left) > strSum(right) {
		comparing++
	}

	if comparing > 1 {
		return true
	}

	return false
}

func strSum(one string) (sum int) {
	for s := range one {
		sum += s
	}

	return
}
