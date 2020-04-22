package multichannel_timeout

import (
	"context"
	"sync"
	"testing"
)



func TestMusic(t *testing.T) {
	MusicChan := make(chan string, 5)

	ctx,cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	MusicChannel(MusicChan)


	for i :=0;i < 3;i++ {
		wg.Add(1)
		go func(ctx context.Context) {
			for {
				if IsCancel(ctx) {
					Output("I don't want to listen this channel anymore!", "")
					break
				} else {
					_, err := MusicListener(MusicChan)
					if err == "" {
						break
					}
				}
			}
			wg.Done()
		}(ctx)
	}

	wg.Wait()
	cancel()
	Output("Stop listening!\n", "")
}


func TestNews(t *testing.T) {
	NewsChan := make(chan string, 5)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	NewsChannel(NewsChan)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(ctx context.Context) {
			for {
				if IsCancel(ctx) {
					Output("I don't want to listen this channel anymore!", "")
					break
				} else {
					_, err := MusicListener(NewsChan)
					if err == "" {
						break
					}
				}
			}
			wg.Done()
		}(ctx)
	}

	wg.Wait()
	cancel()
	Output("Stop listening!\n", "")
}