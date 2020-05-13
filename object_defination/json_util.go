package object_defination

import (
	"encoding/json"
	"fmt"
)

func JsonResolve(obj interface{}) (data []byte, err error) {

	res, err := json.Marshal(obj)
	if err == nil {
		fmt.Printf("%s\n", res)
		return res, nil
	} else {
		fmt.Errorf("resovle user info from json failed!\n%s ", err)
		return nil, err
	}
}


func JsonGenerator(jsonObj []byte, obj interface{}) (data interface{}, err error) {

	resolveErr := json.Unmarshal(jsonObj, &obj)
	if resolveErr == nil {
		fmt.Printf("%s", jsonObj)
		return jsonObj, nil
	} else {
		fmt.Errorf("resovle user info from json failed!\n%s ", err)
		return nil, err
	}
}