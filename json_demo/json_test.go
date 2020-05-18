package json_demo

import (
	"fmt"
	"testing"

	"github.com/sunnywalden/go_learning/object_defination"
)

func testJsonGene(testData interface{}) (res []byte, err error) {

	data, resolveErr := object_defination.JsonGenerator(testData)
	if resolveErr == nil {
		fmt.Print("Json encode success!\n")
		fmt.Printf("%s\n", data)
	} else {
		fmt.Errorf("%e\n", err)
	}

	return data, err
}

func testJsonReso(jsonData []byte, originType interface{}) (res interface{}, err error) {

	data, resolveErr := object_defination.JsonResolve(jsonData, originType)
	if resolveErr == nil {
		fmt.Print("Json decode success!\n")
		fmt.Printf("%s\n", data)
	} else {
		fmt.Errorf("%e\n", resolveErr)
	}

	return data, resolveErr
}

func BenchmarkJsonEncode(b *testing.B) {
	user := &object_defination.User{
		Id: "0001",
		Username: "henry",
		Password: "Demo",
		Age: 30,
	}

	b.StartTimer()
	data, err := testJsonGene(user)
	b.StopTimer()
	if err != nil {
		b.Errorf("%e\n", err)
	} else {
		b.Logf("%s\n", data)
	}
}


func BenchmarkJsonDecode(b *testing.B) {
	user := &object_defination.User{
		Id: "0001",
		Username: "henry",
		Password: "Demo",
		Age: 30,
	}

	data, err := object_defination.JsonGenerator(user)
	if err != nil {
		b.Errorf("%e\n", err)
	} else {
		b.Logf("%s\n", data)
		b.StartTimer()
		res, decodeErr := object_defination.JsonResolve(data, object_defination.User{})
		b.StopTimer()
		if decodeErr != nil {
			b.Errorf("%e\n", decodeErr)
		} else {
			b.Log("Json decode success!\n")
			b.Logf("%s\n", res)
		}
	}
}

func testJsonDump(testData interface{}) (res []byte, err error) {

	data, resolveErr := JsonDump(testData)
	if resolveErr == nil {
		fmt.Print("Json dump success!\n")
		fmt.Printf("%s\n", data)
	} else {
		fmt.Errorf("%e\n", err)
	}

	return data, err
}

func BenchmarkJsonDump(b *testing.B) {
	user := &object_defination.User{
		Id: "0001",
		Username: "henry",
		Password: "Demo",
		Age: 30,
	}

	b.StartTimer()
	data, err := testJsonDump(user)
	b.StopTimer()
	if err != nil {
		b.Errorf("%e\n", err)
	} else {
		b.Logf("%s\n", data)
	}
}

func testJsonLoad(testData []byte, originType interface{}) (res interface{}, err error) {

	data, resolveErr := JsonLoad(testData, originType)
	if resolveErr == nil {
		fmt.Print("Json load success!\n")
		fmt.Printf("%s\n", data)
	} else {
		fmt.Errorf("%e\n", err)
	}

	return data, err
}

func BenchmarkJsonLoad(b *testing.B) {
	user := &object_defination.User{
		Id: "0001",
		Username: "henry",
		Password: "Demo",
		Age: 30,
	}

	data, err := testJsonDump(user)
	if err != nil {
		b.Errorf("%e\n", err)
	} else {
		b.Logf("%s\n", data)
		b.StartTimer()
		res, err := testJsonLoad(data, object_defination.User{})
		b.StopTimer()
		if err != nil {
			b.Errorf("%s", err)
		} else {
			b.Logf("Json load success!%s", res)
		}
	}
}