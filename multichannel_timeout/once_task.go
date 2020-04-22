package multichannel_timeout

import "sync"

func OnceTask(once sync.Once, f func())  {
	//var once sync.Once
	once.Do(f)
}
