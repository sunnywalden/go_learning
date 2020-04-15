package object_defination

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	Id string
	Username string
	Password string
	birth time.Time
	age int
}

func GetUserName(u *User) {
	fmt.Printf("Username is %s\n", u.Username)
}

func TestUser(t *testing.T) {
	u1 := new(User)
	u1.Username = "sunnywalden"
	u1.age = 30
	u1.Id = "0001"
	u1.Password = "walden0517"
	GetUserName(u1)
}
