package gobench

import (
	"testing"
)

func BenchmarkGoMapAdd(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GoMapAdd()
	}

}
func BenchmarkGoStructAdd(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GoStructAdd()
	}

}
