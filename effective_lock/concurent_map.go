package effective_lock

import (
	"errors"

	"github.com/orcaman/concurrent-map"
)

func ConcurrentMapGet(m cmap.ConcurrentMap, name string) (info interface{}, err error) {
	res, ok := m.Get(name)
	if !ok {
		return nil, errors.New("get failed")
	}
	return res, nil
}

func ConcurrentMapSet(m cmap.ConcurrentMap, name string, info []string) () {
	m.Set(name, info)
}

func ConcurrentMapDel(m cmap.ConcurrentMap, name string) () {
	m.Remove(name)
}