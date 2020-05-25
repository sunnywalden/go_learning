package multichannel_timeout

import (
	"fmt"
	"sync"
	"testing"
)

func OnceDemo() {
	age := 30
	age += 1
	fmt.Printf("Age:%d.\n", age)
}

func TestOnceTask(t *testing.T) {
	var wg sync.WaitGroup
	var once sync.Once
	for i := 0;i < 5;i++ {
		wg.Add(1)

		go func(i int,once sync.Once) {
			fmt.Printf("Task %d is running.\n"  ,i)
			OnceTask(once, OnceDemo)
			wg.Done()
		}(i, once)
	}
	wg.Wait()
}
