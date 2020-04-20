package thread_safe

import (
	"container/list"
	"sync"
	"testing"
)


func TestMutex(t *testing.T) {
	var mut sync.RWMutex
	var wg sync.WaitGroup

	allPlayers := list.New()
	allPlayers.PushBack("Henry Zhang")
	allPlayers.PushBack("David Beckham")
	allPlayers.PushBack("Cristiano Ronaldo")
	allPlayers.PushBack("Mesut Ozil")

	WritePlayers := allPlayers

	for p := allPlayers.Front(); p != nil; p = p.Next() {
		wg.Add(1)
		go func(p *list.Element) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			t.Logf("Player %s's coming!\n", p.Value)
			WritePlayers.Remove(p)
			wg.Done()
		}(p)
	}
	//AllPlayers := []string{"Henry Zhang", "David Beckham", "Cristiano Ronaldo", "Mesut Ozil"}
	//for _,p := range AllPlayers {
	//	wg.Add(1)
	//	go func(p string) {
	//		defer func() {
	//			mut.Unlock()
	//
	//		}()
	//		mut.Lock()
	//		t.Logf("Player %s's coming!\n", p)
	//		wg.Done()
	//	}(p)
	//}
	wg.Wait()
	t.Log("All players showed up!\n")
}
