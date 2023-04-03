package main

import (
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/sort", sorting)
	log.Printf("Serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sorting(w http.ResponseWriter, r *http.Request) {
	req := r.Body

	var sl []int
	err := json.NewDecoder(req).Decode(&sl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		req.Close()
		return
	}

	req.Close()
	bubbleSort(sl)
	w.WriteHeader(http.StatusOK)
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

	stubFunc(sl)
}

func stubFunc(sl []int) {
	for i, s := range sl {
		sl[i] = s * 2
	}
}
