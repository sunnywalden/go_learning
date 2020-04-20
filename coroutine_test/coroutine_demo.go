package coroutine_test

import (
	"fmt"
)


func PrintPlayer(players []string) {
	for _,p := range players {
		go func(p string) {
			fmt.Printf("Welcome for famous football player %s \n", p)
		}(p)
	}

}

