package buffer_channel

import (
	"fmt"
	"time"
)

func AsyncTask(t string) chan string {
	Chan := make(chan string, 5)

		go func(t string) {
			fmt.Printf("World famous Player %s 's coming!'\n'", t)
			Chan <- t
		}(t)
		return Chan

}

func AsyncTasks(ts []string) {
	for _,t := range ts {
		go func(t string) {
			res := AsyncTask(t)
			time.Sleep(time.Second * 5)
			fmt.Printf("All you know %s acts so well in every game he played.\n", <-res)
		}(t)
	}
}

func PrepareMatch() {
	time.Sleep(time.Second * 15)
	fmt.Print("Match'll start 15 minutes later!\n")
}
