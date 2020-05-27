package effective_lock

import (
	"testing"

	"github.com/orcaman/concurrent-map"
	"github.com/sunnywalden/go_learning/effective_lock"
)


var waldenInfo = []string{"SunnyWalden", "30"}
var henryInfo  = []string{"Henry Zhang", "28"}


func initData(userInfo []string) (res []string) {
	var info []string
	for _,use := range userInfo {
		info = append(info, use)
	}
	return info
}


func BenchmarkSyncMapSet(b *testing.B) {
	m := cmap.New()


	b.StartTimer()
	effective_lock.ConcurrentMapSet(m, "walden", waldenInfo)
	//effective_lock.ConcurrentMapSet(m, "henry", henryInfo)
	b.StopTimer()
}

func BenchmarkSyncMapGet(b *testing.B) {
	m := cmap.New()
	effective_lock.ConcurrentMapSet(m, "walden", waldenInfo)
	effective_lock.ConcurrentMapSet(m, "henry", henryInfo)

	b.StartTimer()
	info ,err := effective_lock.ConcurrentMapGet(m, "walden")
	if err != nil {
		b.Log(err)
	} else {
		b.Logf("%s info %s", "walden", info)
	}
	b.StopTimer()
}