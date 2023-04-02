package main

import (
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	http.HandleFunc("/sort", sorting)
	log.Printf("Serving on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func sorting(w http.ResponseWriter, r *http.Request) {
	sortSlice()
	w.WriteHeader(http.StatusOK)
}

func sortSlice() {
	sl := create()
	bubbleSort(sl)
}

func create() []int {
	s := seed()
	return createForBubble(s)
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

func bubbleSort(sl []int) {
	var sorts bool

	for !sorts {
		sorts = true
		firstVal := sl[0]
		sorts = iteration(firstVal, sl)
	}
}

func iteration(current int, sl []int) bool {
	sorts := true

	for i := 1; i < len(sl); i++ {
		next := sl[i]
		if current > next {
			sl[i], sl[i-1] = sl[i-1], sl[i]
			sorts = false
		} else {
			current = next
		}
	}

	return sorts
}
