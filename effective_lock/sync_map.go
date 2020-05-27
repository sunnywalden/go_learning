package effective_lock

import (
	"errors"
	"sync"
)
//type SyncMap struct {
//	Get(s *sync.Map) ()
//}

func SyncMapGet(s *sync.Map, name string) (info interface{}, err error) {
	res, ok := s.Load(name)
	if ok {
		return nil, errors.New("get failed")
	}
	return res, nil
}

func SyncMapSet(s *sync.Map, name string, info interface{}) () {
	s.Store(name, info)
}

func SyncMapDel(s *sync.Map, name string) () {
	s.Delete(name)
}
