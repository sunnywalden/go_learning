package interface_defination

import "testing"

type UserManage interface {
	GetUserName() string
	GetNickName() string

}

type OpsUserManage struct {
	UserName string
	NickName string
}

func GetUserName(u *OpsUserManage) string {
	return u.UserName
}

func GetNickName(u *OpsUserManage) string {
	return u.NickName
}

func TestGetUserName(t *testing.T) {
	om := new(OpsUserManage)
	om.UserName = "sunnywalden"
	om.NickName = "Walden"
	t.Log("Username: ",GetUserName(om))
	t.Log("User nick name: ", GetNickName(om))
}