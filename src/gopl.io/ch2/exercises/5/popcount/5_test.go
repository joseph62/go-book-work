package popcount

import "testing"

func BenchmarkExpressionPopCountPerformance(t *testing.B) {
	benchmarkPerformance(ExpressionPopCount, t)
}

func BenchmarkShiftingPopCountPerformance(t *testing.B) {
	benchmarkPerformance(ShiftingPopCount, t)
}

func BenchmarkRightPopCountPerformance(t *testing.B) {
	benchmarkPerformance(RightPopCount, t)
}

func benchmarkPerformance(f func(uint64) int, t *testing.B) {
	for i := 0; i < t.N; i++ {
		f(uint64(i))
	}
}
