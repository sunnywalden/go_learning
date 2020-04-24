package benchmark_demo

import (
	unit "go_learning/unit_test"
	"testing"
)

func BenchmarkCapitalise(b *testing.B) {
	testStr := "sunnyWalden"
	b.StartTimer()
	unit.Capitalize(testStr)
	b.StopTimer()
}

func BenchmarkCapitalised(b *testing.B) {
	testStr := "sunnyWalden"
	b.StartTimer()
	unit.Capitalized(testStr)
	b.StopTimer()
}
