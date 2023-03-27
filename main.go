package main

import (
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

type st struct {
	f string
	s string
	t string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
	sl := createForBubble()
	bubbleSort(sl)
}

func createForBubble() []st {
	newSl := make([]st, 0)

	for i := 0; i < 10000; i++ {
		newSl = append(newSl, st{f: randStringBytes(), s: randStringBytes(), t: randStringBytes()})
	}
	return newSl
}

func bubbleSort(sl []st) {
	var sorts bool

	for !sorts {
		sorts = true
		firstVal := sl[0]
		sorts = iteration(firstVal, sl)
	}
}

func iteration(current st, sl []st) bool {
	sorts := true

	for i := 1; i < len(sl); i++ {
		val2 := sl[i]
		if compare(current, val2) {
			sl[i], sl[i-1] = sl[i-1], sl[i]
			sorts = false
		} else {
			current = val2
		}
	}

	return sorts
}

func strSum(one string) (sum int) {
	for s := range one {
		sum += s
	}

	return
}

func compare(cur, next st) bool {
	count := 0
	if strSum(cur.f) > strSum(next.f) {
		count++
	}

	count, success := second(cur, next, count)
	if success {
		return true
	}

	if third(cur, next, count) {
		return true
	}

	return false
}

func second(cur, next st, c int) (int, bool) {
	if strSum(cur.s) > strSum(next.s) {
		c++
	}

	if c > 1 {
		return c, true
	}

	return c, false
}

func third(cur, next st, c int) bool {
	if strSum(cur.t) > strSum(next.t) {
		c++
	}

	if c > 1 {
		return true
	}
	return false
}

func randStringBytes() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
