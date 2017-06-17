package utils

import (
	"encoding/json"
	"log"
	"errors"
	//"reflect"
)

func ParseBytesToJson(bytes []byte) (retJson map[string]interface{}, err error) {
	if err = json.Unmarshal(bytes, &retJson); err != nil {
		log.Println(err)
	}
	return
}

func ParseStringToJson(str string) (retJson map[string]interface{}, err error) {
	retJson, err = ParseBytesToJson([]byte(str))
	return
}

func GetValueFromInterfaceMap(m interface{}, key string) (result interface{}, err error) {
	data, ok := m.(map[string]interface{})
	if !ok {
		err = errors.New("m is not map[string]interface{} type.")
		return
	}
	result = data[key]
	return
}

func GetStringFromInterfaceMap(m interface{}, key string) (result string, err error) {
	data, err := GetValueFromInterfaceMap(m, key)
	if err != nil {
		return
	}
	result, ok := data.(string)
	if !ok {
		errInfo := "value is not string type."
		log.Println(errInfo)
		err = errors.New(errInfo)
		return
	}
	return
}

func GetFloat64FromInterfaceMap(m interface{}, key string) (result float64, err error) {
	data, err := GetValueFromInterfaceMap(m, key)
	if err != nil {
		return
	}
	
	result, ok := data.(float64)
	if !ok {
		errInfo := "value is not float64 type." //reflect.TypeOf(data)
		log.Println(errInfo)
		err = errors.New(errInfo)
		return
	}
	return
}