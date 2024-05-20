package addition

import "testing"

func BenchmarkAddition(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		ToBeTested(1000, 3000)
	}
}
