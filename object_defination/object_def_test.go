package object_defination

import (
	"testing"
	myuser "github.com/sunnywalden/go_learning/object_defination"
)



func TestUser(t *testing.T) {
	u1 := &myuser.User{}
	u1.Username = "sunnywalden"
	u1.Age = 30
	u1.Id = "0001"
	u1.Password = "walden0517"
	GetUserName(u1)
}
