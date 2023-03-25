package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

const bubble = "bubble_sort"
const fusion = "fusion_sort"

func main() {
	http.HandleFunc("/sort", sorting)
	log.Printf("Serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sorting(w http.ResponseWriter, r *http.Request) { //203 = 203 - встраивалась до этого?
	src, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	method := string(src)
	_ = sortSlice(method)

	w.WriteHeader(http.StatusOK)
}

func sortSlice(method string) bool { //153 -> 498
	sl := create(method)
	b := sort(sl, method)
	return b
}

func create(m string) []int { //201 -> 345
	sl := make([]int, 0)
	s := seed()

	switch m {
	case bubble:
		sl = createForBubble(s)
	case fusion:
		sl = createForFusion(s)
	}

	return sl
}

func seed() rand.Source { //183 -> 369
	return rand.NewSource(time.Now().UnixNano())
}

func createForBubble(s rand.Source) []int { //129 = 129
	newSl := make([]int, 0)

	for i := 0; i < 1000; i++ {
		n := rand.New(s).Int()
		newSl = append(newSl, n)
	}
	return newSl
}

func createForFusion(s rand.Source) []int { //129 = 129
	newSl := make([]int, 0)

	for i := 0; i < 10000; i++ {
		n := rand.New(s).Int()
		newSl = append(newSl, n)
	}
	return newSl
}

func sort(sl []int, m string) bool { //136 = 136
	switch m {
	case bubble:
		bubbleSort(sl)
	case fusion:
		fusionSort(sl)
	}

	return true
}

func bubbleSort(sl []int) { //65 = 65
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

func fusionSort(sl []int) []int {
	if len(sl) == 1 {
		return sl
	}
	left := fusionSort(sl[0 : len(sl)/2])
	right := fusionSort(sl[len(sl)/2:])

	result := make([]int, len(sl))

	l, r, k := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result[k] = left[l]
			l++
		} else {
			result[k] = right[r]
			r++
		}
		k++
	}

	for l < len(left) {
		result[k] = left[l]
		l++
		k++
	}
	for r < len(right) {
		result[k] = right[r]
		r++
		k++
	}
	return result
}
