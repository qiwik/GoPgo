package main

import "fmt"

type st struct {
	f int
	s int
	t int
}

func main() {
	sl := []st{{1, 2, 3}, {10, 12, 4}, {0, 0, 71}}

	sorts := false

	for !sorts {
		sorts = true
		val1 := sl[0]

		sorts = compare(val1, sl)
	}

	fmt.Println(sl)
}

func compare(current st, sl []st) bool {
	sorts := true

	for i := 1; i < len(sl); i++ {
		val2 := sl[i]
		if first(current, val2) {
			sl[i], sl[i-1] = sl[i-1], sl[i]
			sorts = false
		} else {
			current = val2
		}
	}

	return sorts
}

func first(cur, next st) bool {
	k := 0
	if cur.f > next.f {
		k++
	}

	return second(cur, next, k)
}

func second(cur, next st, k int) bool {
	if cur.s > next.s {
		k++
	}

	if k > 1 {
		return true
	}

	return third(cur, next, k)
}

func third(cur, next st, k int) bool {
	if cur.t > next.t {
		k++
	}

	if k > 1 {
		return true
	}
	return false
}
