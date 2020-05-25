package reflect_demo

import (
	"fmt"
	"github.com/sunnywalden/go_learning/unit"
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
