package popcount

import "testing"

func BenchmarkLoopPopCountPerformance(t *testing.B) {
	benchmarkPerformance(LoopPopCount, t)
}

func BenchmarkExpressionPopCountPerformance(t *testing.B) {
	benchmarkPerformance(ExpressionPopCount, t)
}

func benchmarkPerformance(f func(uint64) int, t *testing.B) {
	for i := 0; i < t.N; i++ {
		f(uint64(i))
	}
}