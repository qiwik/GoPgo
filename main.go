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
	log.Printf("Serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sorting(w http.ResponseWriter, r *http.Request) {
	sortSlice()
	w.WriteHeader(http.StatusOK)
}

func sortSlice() {
	s := seed()
	sl := createForBubble(s)
	bubbleSort(sl)
}

func seed() rand.Source {
	return rand.NewSource(time.Now().UnixNano())
}

func createForBubble(s rand.Source) []int {
	newSl := make([]int, 0)

	for i := 0; i < 1000; i++ {
		n := rand.New(s).Int()
		newSl = append(newSl, n)
	}
	return newSl
}

func bubbleSort(sl []int) {
	sorts := false

	for !sorts {
		sorts = true
		val1 := sl[0]

		for i := 1; i < len(sl); i++ {
			val2 := sl[i]
			if val1 > val2 {
				sl[i], sl[i-1] = sl[i-1], sl[i]
				sorts = false
			} else {
				val1 = val2
			}
		}
	}
}
