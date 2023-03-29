package main

import (
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

type st struct {
	first  string
	second string
	third  string
	fourth string
	fifth  string
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
		newSl = append(newSl, st{
			first:  randStringBytes(),
			second: randStringBytes(),
			third:  randStringBytes(),
			fourth: randStringBytes(),
			fifth:  randStringBytes(),
		})
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
		next := sl[i]
		if compare(current, next) {
			sl[i], sl[i-1] = sl[i-1], sl[i]
			sorts = false
		} else {
			current = next
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
	if strSum(cur.first) > strSum(next.first) {
		count++
	}

	count = secondStr(cur, next, count)

	count, success := thirdStr(cur, next, count)
	if success {
		return true
	}

	count, success = fourthStr(cur, next, count)
	if success {
		return true
	}

	return fifthStr(cur, next, count)
}

func secondStr(cur, next st, c int) int {
	if strSum(cur.second) > strSum(next.second) {
		c++
	}

	return c
}

func thirdStr(cur, next st, c int) (int, bool) {
	if strSum(cur.third) > strSum(next.third) {
		c++
	}

	if c > 2 {
		return c, true
	}
	return c, false
}

func fourthStr(cur, next st, c int) (int, bool) {
	if strSum(cur.fourth) > strSum(next.fourth) {
		c++
	}

	if c > 2 {
		return c, true
	}
	return c, false
}

func fifthStr(cur, next st, c int) bool {
	if strSum(cur.fifth) > strSum(next.fifth) {
		c++
	}

	if c > 2 {
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
