package multichannel_timeout

import (
	"context"
	"sync"
	"testing"

	mulChan "github.com/sunnywalden/go_learning/multichannel_timeout"
)



func TestMusic(t *testing.T) {
	MusicChan := make(chan string, 5)

	ctx,cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	mulChan.MusicChannel(MusicChan)


	for i :=0;i < 3;i++ {
		wg.Add(1)
		go func(ctx context.Context) {
			for {
				if IsCancel(ctx) {
					mulChan.Output("I don't want to listen this channel anymore!", "")
					break
				} else {
					_, err := mulChan.MusicListener(MusicChan)
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
	mulChan.Output("Stop listening!\n", "")
}


func TestNews(t *testing.T) {
	NewsChan := make(chan string, 5)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	mulChan.NewsChannel(NewsChan)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(ctx context.Context) {
			for {
				if IsCancel(ctx) {
					mulChan.Output("I don't want to listen this channel anymore!", "")
					break
				} else {
					_, err := mulChan.MusicListener(NewsChan)
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
	mulChan.Output("Stop listening!\n", "")
}