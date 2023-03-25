package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

const first = "f"
const second = "s"

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
	case first:
		sl = forFirst(s)
	case second:
		sl = forSecond(s)
	}

	return sl
}

func seed() rand.Source { //183 -> 369
	return rand.NewSource(time.Now().UnixNano())
}

func forFirst(s rand.Source) []int { //129 = 129
	newSl := make([]int, 0)

	for i := 0; i < 1000; i++ {
		n := rand.New(s).Int()
		newSl = append(newSl, n)
	}
	return newSl
}

func forSecond(s rand.Source) []int { //129 = 129
	newSl := make([]int, 0)

	for i := 0; i < 10000; i++ {
		n := rand.New(s).Int()
		newSl = append(newSl, n)
	}
	return newSl
}

func sort(sl []int, m string) bool { //136 = 136
	switch m {
	case first:
		bubble(sl)
	case second:
		fusion(sl)
	}

	return true
}

func bubble(sl []int) { //65 = 65
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

func fusion(sl []int) []int {
	if len(sl) == 1 {
		return sl
	}
	left := fusion(sl[0 : len(sl)/2])
	right := fusion(sl[len(sl)/2:])

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
