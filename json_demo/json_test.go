package json_demo

import (
	"fmt"
	"testing"

	"github.com/sunnywalden/go_learning/object_defination"
)

func testJsonResolve() (res []byte, err error) {
	user := &object_defination.User{
		Id: "0001",
		Username: "henry",
		Password: "Demo",
		Age: 30,
	}

	data, resolveErr := object_defination.JsonResolve(user)
	if resolveErr == nil {
		fmt.Print("Json decode success!\n")
		fmt.Printf("%s\n", data)
	} else {
		fmt.Errorf("%e\n", err)
	}

	return data, err
}

func TestJsonResolve(t *testing.T) {
	data, err := testJsonResolve()
	if err != nil {
		t.Errorf("%e\n", err)
	} else {
		t.Logf("%s\n", data)
	}
}


func TestJsonDecode(t *testing.T) {
	data, err := testJsonResolve()
	if err != nil {
		t.Errorf("%e\n", err)
	} else {
		t.Logf("%s\n", data)
		res, decodeErr := object_defination.JsonGenerator(data, object_defination.User{})
		if decodeErr != nil {
			t.Errorf("%e\n", decodeErr)
		} else {
			t.Log("Json decode success!\n")
			t.Logf("%s\n", res)
		}
	}
}