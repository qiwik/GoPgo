package main

import (
	"testing"
)

func BenchmarkLoad(b *testing.B) {
	s := seed()
	sl := createForBubble(s)

	for i := 0; i < b.N; i++ {
		if err := load(sl); err != nil {
			b.Errorf("generateLoad got err %v want nil", err)
		}
	}
}
