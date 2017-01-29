package utils

import "encoding/json"

//ToJSONOrPanic ...
func ToJSONOrPanic(instance interface{}) string {
	encodedMessage, err := json.Marshal(instance)
	if err != nil {
		panic(err)
	}
	return string(encodedMessage)
}

//ToJSON ...
func ToJSON(instance interface{}) (string, error) {
	encodedMessage, err := json.Marshal(instance)
	return string(encodedMessage), err
}
