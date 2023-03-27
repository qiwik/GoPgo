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
	if sortSlice() {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func sortSlice() bool { //153 -> 498
	sl := create()
	bubbleSort(sl)
	return true
}

func create() []st { //201 -> 345
	return createForBubble()
}

//func seed() rand.Source { //183 -> 369
//	return rand.NewSource(time.Now().UnixNano())
//}

func createForBubble() []st { //129 = 129
	newSl := make([]st, 0)

	for i := 0; i < 10000; i++ {
		newSl = append(newSl, struct {
			f string
			s string
			t string
		}{f: randStringBytes(), s: randStringBytes(), t: randStringBytes()})
	}
	return newSl
}

//func createForFusion(s rand.Source) []int { //129 = 129
//	newSl := make([]int, 0)
//
//	for i := 0; i < 1000; i++ {
//		n := rand.New(s).Int()
//		newSl = append(newSl, n)
//	}
//	return newSl
//}

func bubbleSort(sl []st) { //65 = 65
	var sorts bool

	for !sorts {
		sorts = true
		firstVal := sl[0]
		sorts = iteration(firstVal, sl)
	}
}

//func iteration(current int, sl []int) bool {
//	sorts := true
//
//	for i := 1; i < len(sl); i++ {
//		next := sl[i]
//		if current > next {
//			sl[i], sl[i-1] = sl[i-1], sl[i]
//			sorts = false
//		} else {
//			current = next
//		}
//	}
//
//	return sorts
//}

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

//func fusionSort(sl []int) []int {
//	if len(sl) == 1 {
//		return sl
//	}
//	left := fusionSort(sl[0 : len(sl)/2])
//	right := fusionSort(sl[len(sl)/2:])
//
//	result := make([]int, len(sl))
//
//	l, r, k := 0, 0, 0
//	for l < len(left) && r < len(right) {
//		if left[l] <= right[r] {
//			result[k] = left[l]
//			l++
//		} else {
//			result[k] = right[r]
//			r++
//		}
//		k++
//	}
//
//	for l < len(left) {
//		result[k] = left[l]
//		l++
//		k++
//	}
//	for r < len(right) {
//		result[k] = right[r]
//		r++
//		k++
//	}
//	return result
//}
