package reflect_demo

import (
	"fmt"
	unit "go_learning/unit_test"
	"reflect"
	"testing"
)


func SetName(name string,info interface{}) interface{} {
	userInfo := &info
	reflect.ValueOf(userInfo).Elem().Set(reflect.ValueOf(name))
	fmt.Printf("User name: %s\n", *userInfo)
	return info
}

func TestCommonFunc(t *testing.T) {
	employeeHenry := &unit.Employee{}
	vipHenry := new(VIP)

	epHenry := SetName("Henry Zhang", employeeHenry)
	vpHenry := SetName("Hnery", vipHenry)
	t.Log(epHenry)
	t.Log(vpHenry)
}
