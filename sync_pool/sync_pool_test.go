package sync_pool

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Printf("New employee!\n")
			return ""
		},
	}

	var wg sync.WaitGroup
	employees := []string{"SunnyWalden", "Henry", "Robin"}
	for _,ep := range  employees {
		wg.Add(1)
		go func(ep string) {
			fmt.Printf("Put %s into pool.\n", ep)
			pool.Put(ep)
			wg.Done()
		}(ep)
		wg.Wait()
	}
	for i := 0;i <= len(employees); i++ {
		employee := pool.Get().(string)
		if employee != "" {
			fmt.Printf("Employee %s's coming!\n", employee)
		} else {
			fmt.Printf("No employee's available!\n")
		}
	}
}



