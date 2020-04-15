package func_test

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func OutputTime() string {
	date_time := time.Now().String()
	date := strings.Split(date_time, " ")[0]
	fmt.Printf("Today is %s", date)
	return date
}

func TimeSpent (fc func () string) func() string {
	return 	func () string {
		start_time := time.Now()
		res := fc()
		spent_time := time.Since(start_time).Seconds()
		fmt.Println("\n", spent_time, " seconds spent")
		return res
	}
}

func TestTimeSpent(t *testing.T) {
	ts := TimeSpent(OutputTime)
	res := ts()
	t.Log(res)
}