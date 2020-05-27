package effective_lock

import (
	"github.com/orcaman/concurrent-map"
)

func ConcurrentMapGet(m cmap.ConcurrentMap, name string) (info []string, err error) {
	res, err := m.get(name)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ConcurrentMapSet(m cmap.ConcurrentMap, name string, info []string) (success bool, err error) {
	_, setErr := m.set(name, info)
	if setErr != nil {
		return false, setErr
	}
	return true, nil
}

func ConcurrentMapDel(m cmap.ConcurrentMap, name string) (success bool, err error) {
	_, delErr := m.del(name)
	if delErr != nil {
		return false, delErr
	}
	return true, nil
}