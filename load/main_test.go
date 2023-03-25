package main

import (
	"testing"
)

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := load(); err != nil {
			b.Errorf("generateLoad got err %v want nil", err)
		}
	}
}
