package json_demo

import "github.com/json-iterator/go"

func JsonDump(originData interface{}) (data []byte,err error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonData, encodeErr := json.Marshal(&originData)
	if encodeErr != nil {
		return nil, encodeErr
	} else {
		return jsonData, nil
	}
}

func JsonLoad(jsonData []byte, originType interface{}) (data interface{},err error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	decodeErr := json.Unmarshal(jsonData, &originType)
	if decodeErr != nil {
		return nil, decodeErr
	} else {
		return &originType, nil
	}
}
