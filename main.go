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

func randStringBytes() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	http.HandleFunc("/sort", sorting)
	log.Printf("Serving on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func sorting(w http.ResponseWriter, r *http.Request) { //203 = 203 - встраивалась до этого?
	sortSlice()
	w.WriteHeader(http.StatusOK)
}

func sortSlice() { //153 -> 498
	sl := createForBubble()
	bubbleSort(sl)
}

func createForBubble() []st { //129 = 129
	newSl := make([]st, 0)

	for i := 0; i < 10000; i++ {
		newSl = append(newSl, st{f: randStringBytes(), s: randStringBytes(), t: randStringBytes()})
	}
	return newSl
}

func bubbleSort(sl []st) { //65 = 65
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
	k := 0
	if strSum(cur.f) > strSum(next.f) {
		k++
	}

	if second(cur, next, k) {
		return true
	}

	if third(cur, next, k) {
		return true
	}

	return false
}

func second(cur, next st, k int) bool {
	if strSum(cur.s) > strSum(next.s) {
		k++
	}

	if k > 1 {
		return true
	}

	return false
}

func third(cur, next st, k int) bool {
	if strSum(cur.t) > strSum(next.t) {
		k++
	}

	if k > 1 {
		return true
	}
	return false
}
