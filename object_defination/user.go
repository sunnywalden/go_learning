package object_defination

import (
	"fmt"
	"time"
)

type User struct {
	Id string
	Username string `json:"username"`
	Password string `json:"password"`
	Birth time.Time	`json:"birth"`
	Age int	`json:"age"`
}

func GetUserName(u *User) {
	fmt.Printf("Username is %s\n", u.Username)
}
