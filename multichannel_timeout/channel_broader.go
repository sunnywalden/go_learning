package multichannel_timeout

import (
	"context"
	"fmt"
	"time"
)

func Output(res string, err string) {
	if err == "" {
		fmt.Printf("%s \n", res)
	} else {
		fmt.Errorf(err)
	}
}

func TaskCancel(Chan chan string) {
	close(Chan)
}


func IsCancel(ctx context.Context) bool {
	select {
	case <- ctx.Done():
		return true
		default:
			return false

	}
}

func MusicChannel(MusicChan chan string) {
	var Musics = [...]string{"La la City", "We back to love", "Reality"}

	Output("You're listening to music channel.\n", "")

	for _,music := range Musics {


				Output("Next song's " + music + ".\n", "")
				MusicChan <- music
				time.Sleep(time.Second * 4)


	}
}


func NewsChannel(NewsChan chan string) {
	var News = [...]string{"Sony invest Bilibili for 4 billion dollars!", "Crown's flowing all over the world!"}
	Output("You're listening to Breaking News channel.\n","")

	for _,n := range News {
		Output(n + "!" + "\n","")
		NewsChan <- n
		time.Sleep(time.Second * 3)
	}
}

//func MultiChannel(MusicChan chan string,NewsChan chan string) {
//	MusicChannel(MusicChan)
//	NewsChannel(NewsChan)
//	select {
//	case res := <- MusicChan:
//		fmt.Printf("You're listening to music channel!\n%s\n", res)
//	case res := <- NewsChan:
//		fmt.Printf("You're listening to news channel!\n%s\n", res)
//	case <- time.After(time.Second * 3):
//		fmt.Print("Timeout after 3 seconds!")
//	default:
//		fmt.Print("No channel available!")
//		close(MusicChan)
//		close(NewsChan)
//	}
//}