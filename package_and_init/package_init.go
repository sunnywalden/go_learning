package package_and_init

import "fmt"

func init() {
	fmt.Print("###########################")
}

func SayHi(name string) {
	defer func() {
		fmt.Print("###########################")
	}()
	fmt.Printf("Hello %s", name)
}
