package effective_lock

import (
	"sync"
	"testing"

	"github.com/orcaman/concurrent-map"
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


func BenchmarkConMapSet(b *testing.B) {
	m := cmap.New()


	b.StartTimer()
	ConcurrentMapSet(&m, "walden", &waldenInfo)
	ConcurrentMapSet(&m, "henry", &henryInfo)
	b.StopTimer()
}

func BenchmarkConMapGet(b *testing.B) {
	m := cmap.New()
	ConcurrentMapSet(&m, "walden", &waldenInfo)
	ConcurrentMapSet(&m, "henry", &henryInfo)

	b.StartTimer()
	for _,name := range [2]string{"walden", "henry"} {
		info, err := ConcurrentMapGet(&m, name)
		if err != nil {
			b.Logf("%s\n", err)
		} else {
			b.Logf("%s info %s\n", name, *info)
		}
	}
	b.StopTimer()
}

func BenchmarkConMapDel(b *testing.B) {
	m := cmap.New()
	ConcurrentMapSet(&m, "walden", &waldenInfo)
	ConcurrentMapSet(&m, "henry", &henryInfo)

	b.StartTimer()
	for _,name := range [2]string{"walden", "henry"} {
		ConcurrentMapDel(&m, name)
	}
	b.StopTimer()
}

func BenchmarkSyncMapSet(b *testing.B) {
	s := sync.Map{}


	b.StartTimer()
	SyncMapSet(&s, "walden", &waldenInfo)
	b.StopTimer()
}

func BenchmarkSyncMapGet(b *testing.B) {
	s := sync.Map{}
	SyncMapSet(&s, "walden", &waldenInfo)
	SyncMapSet(&s, "henry", &henryInfo)

	b.StartTimer()
	for _,name := range []string{"walden", "henry"} {
		res, err := SyncMapGet(&s, name)
		if err != nil {
			b.Logf("%s\n", err)
		} else {
			b.Logf("get res from sync map: %s\n", *res)
		}
	}
	b.StopTimer()
}


func BenchmarkSyncMapDel(b *testing.B) {
	s := sync.Map{}
	SyncMapSet(&s, "walden", &waldenInfo)
	SyncMapSet(&s, "henry", &henryInfo)

	b.StartTimer()
	for _,name := range []string{"walden", "henry"} {
		SyncMapDel(&s, name)
	}
	b.StopTimer()
}