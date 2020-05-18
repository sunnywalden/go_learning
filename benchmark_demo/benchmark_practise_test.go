package benchmark_demo

import (
	"github.com/sunnywalden/go_learning/unit"
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
