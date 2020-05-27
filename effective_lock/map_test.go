package effective_lock

import (
	"testing"

	"github.com/orcaman/concurrent-map"
)

func BenchmarkSyncMap(b *testing.B) {
	b.StartTimer()
	m := cmap.New()
	var info []string
	info = append(info, "SunnyWalden", "30")
	ConcurrentMapSet(m, "walden", info)
	b.StopTimer()
}
