package changable_length_and_defer

import (
	"fmt"
	"testing"
)

func ChangeLenth(sourceList ...[]string) {
	for _,s := range sourceList {
		fmt.Println(s)
	}
}

func LockRelease () {
	fmt.Println("Release lock!")
}

func TestLength(t *testing.T) {
	defer LockRelease()
	list1 := []string{"Sunny", "Walden"}
	list2 := []string{"30", "28", "33"}
	ChangeLenth(list1)
	ChangeLenth(list2)
	fmt.Println("Finished!")
	panic("Eror!")
}
