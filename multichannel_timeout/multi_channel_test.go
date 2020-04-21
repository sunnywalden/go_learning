package multichannel_timeout

import (
	"fmt"
	"sync"
	"testing"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Output(res string, err string) {
	if err == "" {
		fmt.Printf("%s \n", res)
	} else {
		fmt.Errorf(err)
	}
}

func TestMusicChannel(t *testing.T) {
	var wg sync.WaitGroup
	var Musics = [...]string{"La la City", "We back to love", "Reality"}
	for _,m := range Musics {
		wg.Add(1)
		go func(m string) {
			res,err := MusicChannel(m)
			Output(<-res, err)
			wg.Done()
		}(m)

	}
	wg.Wait()

}


func TestNewsChannel(t *testing.T) {
    var wg sync.WaitGroup
	var News = [...]string{"Sony invest Bilibili for 4 billion dollars!", "Crown's flowing all over the world!"}

	for _,n := range News {
		wg.Add(1)
		go func(n string) {

			res,err := NewsChannel(n)
			Output(<-res, err)
			wg.Done()

		}(n)
	}
	wg.Wait()
}


func TestMultiChannel(t *testing.T) {

	var News = [...]string{"Sony invest Bilibili for 4 billions!", "Crown's flowing all over the world!"}
	var Musics = [...]string{"La la City", "We back to love", "Reality"}
	MaxLen := Max(len(News), len(Musics))
	for i := 0;i < MaxLen - 1;i++ {
		if i < len(News) && i < len(Musics) {
			res, err := MultiChannel(Musics[i], News[i])
			Output(res, err)

		} else if i > len(News) && i < len(Musics) {
			res, err := MultiChannel(Musics[i], "")
			Output(res, err)
		} else {
			res, err := MultiChannel("", News[i])
			Output(res, err)
		}

	}

}