package multichannel_timeout

import (
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

func MusicChannel(MusicChan chan string) {
	var Musics = [...]string{"La la City", "We back to love", "Reality"}
	fmt.Print("You're listening to music channel.\n")
	for _,m := range Musics {
			fmt.Printf("Next song's %s !'\n'", m)
			MusicChan <- m
			time.Sleep(time.Second * 5)

	}
	close(MusicChan)

}


func NewsChannel(NewsChan chan string) {
	var News = [...]string{"Sony invest Bilibili for 4 billion dollars!", "Crown's flowing all over the world!"}
	fmt.Print("You're listening to Breaking News channel.\n")

	for _,n := range News {
		fmt.Printf("%s !'\n'", n)
		NewsChan <- n
		time.Sleep(time.Second * 3)
	}
	close(NewsChan)
}

func MultiChannel(MusicChan chan string,NewsChan chan string) {
	//MusicChan := make(chan string, 5)
	//NewsChan := make(chan string, 5)
	MusicChannel(MusicChan)
	NewsChannel(NewsChan)
	select {
	case res := <- MusicChan:
		fmt.Printf("You're listening to music channel!\n%s\n", res)
		//close()
		//return res,""
	case res := <- NewsChan:
		fmt.Printf("You're listening to news channel!\n%s\n", res)
		//close(MusicChan)
		//return res,""
	case <- time.After(time.Second * 3):
		fmt.Print("Timeout after 3 seconds!")
		//res := "Timeout"
		//return "",res
	default:
		fmt.Print("No channel available!")
		//res := "No channel"
		close(MusicChan)
		close(NewsChan)
	}
}