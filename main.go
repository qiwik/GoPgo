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
	var sl []int
	err := json.NewDecoder(r.Body).Decode(&sl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// не использую defer, потому что тогда не произойдет встраивания
		r.Body.Close()
		return
	}

	// не использую defer, потому что тогда не произойдет встраивания
	r.Body.Close()
	bubbleSort(sl)
	w.WriteHeader(http.StatusOK)
}

func bubbleSort(sl []int) {
	var sorts bool

	for !sorts {
		sorts = true
		firstVal := sl[0]
		sorts = iteration(firstVal, sl)
	}

	stubDivide(sl)
}

// Отделяем блок кода, который имеет конкретную функцию
func iteration(current int, sl []int) bool {
	sorts := true

	for i := 1; i < len(sl); i++ {
		next := sl[i]
		current, sorts = compare(current, next, i, sl)
	}

	// нагружаем функцию iteration(). Сначала вычитываем новый порядок, затем делим каждый элемент, потом снова вычитываем
	// для того, чтобы убедиться, что все ок
	stubLoop(sl)
	stubDivide(sl)
	stubLoop(sl)

	return sorts
}

// Опциональная функция. Результаты с ней идут с суффиксом _1 в файлах. Без нее без суффикса
func compare(cur, next, i int, sl []int) (int, bool) {
	if cur > next {
		sl[i], sl[i-1] = sl[i-1], sl[i]
		return cur, false
	}
	return next, true
}

// Симулируем некую полезную работу над слайсом
func stubDivide(sl []int) {
	for i, s := range sl {
		sl[i] = s / 2
	}
}

// Симулируем обход слайса после каждой итерации для того, чтобы, например, выписать обновленный порядок в консоль
func stubLoop(sl []int) {
	for i, _ := range sl {
		_ = sl[i]
	}
}
