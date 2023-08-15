package jsonutil

import (
	"encoding/json"
	"fmt"

	"github.com/luoliDark/base/sysmodel"
	"github.com/valyala/fastjson"
)

// map 转json 字符串
func MapToJson(data map[string]interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		/// 转换失败
		fmt.Errorf("map转json 失败," + err.Error())

		return "", err
	}
	return string(bytes), nil
}

func ListModelToJson(data []sysmodel.Sys_UserWeiXin) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		/// 转换失败
		fmt.Errorf("map转json 失败," + err.Error())

		return "", err
	}
	return string(bytes), nil
}

// listmap 转json 字符串
func ListMapToJson(data []map[string]interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		/// 转换失败
		fmt.Errorf("map数组转json 失败," + err.Error())

		return "", err
	}
	return string(bytes), nil
}

// 通用 转json
func ObjToJson(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		/// 转换失败
		fmt.Errorf("struct 转json 失败," + err.Error())

		return "", err
	}
	return string(bytes), nil
}

// struct 数组转json
func StructToJson(data interface{}) (string, error) {
	return ObjToJson(data)
}

// struct 数组转json
func StructListToJson(data []interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		/// 转换失败
		fmt.Errorf("struct数组 转json 失败," + err.Error())

		return "", err
	}
	return string(bytes), nil
}

// 字符串转 fastjson
func StrToJson(json string) (*fastjson.Value, error) {
	value, e := fastjson.Parse(json)
	if e != nil {
		fmt.Errorf("json 字符串传fastjson 失败," + e.Error())

		return nil, e
	}
	return value, nil
}

// json 字符串转struct 对象
func JsonToStruct(jsonstr string, tar interface{}) (interface{}, error) {
	err := json.Unmarshal([]byte(jsonstr), tar)
	if err != nil {
		fmt.Errorf("json 转 struct 失败," + err.Error())

		return nil, err
	}
	return tar, nil
}

// json 格式转map 字典
func JsonToMap(jsonstr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonstr), result)
	if err != nil {
		fmt.Errorf("json 转map 失败," + err.Error())

		return nil, err
	}
	return result, nil
}

func InterfaceToMap(datamap interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	bytes, err := json.Marshal(datamap)
	if err != nil {
		/// 转换失败
		fmt.Errorf("struct数组 转json 失败," + err.Error())

		return result, err
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Errorf("json 转map 失败," + err.Error())

		return nil, err
	}
	return result, nil
}
