package reflect_demo

import (
	unit "go_learning/unit_test"
	"reflect"
	"testing"
)

func TestReflectBasic(t *testing.T) {
	ep := unit.Employee{}
	ep.SetEmployeeName("sunnywalden")
	ep.SetEmployeeAge(30)

	capName := ep.GetEmployeeName()
	t.Log(reflect.ValueOf(capName), reflect.ValueOf(capName).Type())
}

func TestInvokeByReflect(t *testing.T) {
	ep := &unit.Employee{Name: "sunnywalden", Age: 30}
	if nameValue, ok := reflect.TypeOf(*ep).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' value.\n")
	} else {
		t.Log("Employee name is ", nameValue.Name)
	}

	if selfDes, ok := reflect.TypeOf(*ep).FieldByName("SelfDescribe"); !ok {
		t.Error("Failed to get 'SelfDescribe' value.\n")
	} else {
		t.Log("Employee  describe  ", selfDes)
		t.Log("What is employee  describe:  ", selfDes.Tag.Get("info"))
	}

	reflect.ValueOf(ep).MethodByName("SetEmployeeAge").Call([]reflect.Value{reflect.ValueOf(28)})
	t.Log("Updated age:", ep.Age)
	//ep.SetEmployeeName("sunnywalden")
	//ep.SetEmployeeAge(30)

	//capName := ep.GetEmployeeName()
	//t.Log(reflect.ValueOf(capName), reflect.ValueOf(capName).Type())

}
