package object_defination

import (
	"testing"
)



func TestUser(t *testing.T) {
	u1 := new(User)
	u1.Username = "sunnywalden"
	u1.Age = 30
	u1.Id = "0001"
	u1.Password = "walden0517"
	GetUserName(u1)
}
