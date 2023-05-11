package jsonutil

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestMapToJson(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "钟新建"
	data["age"] = 28
	data["birthdate"] = "19920726"

	result, _ := MapToJson(data)
	fmt.Print(result)
}

func TestListMapToJson(t *testing.T) {
	data := make([]map[string]interface{}, 2)
	data[0] = map[string]interface{}{
		"name":  "He",
		"age":   21,
		"score": 60.5,
	}
	data[1] = map[string]interface{}{}

	data = append(data, map[string]interface{}{
		"name":  "He",
		"age":   21,
		"score": 60.5,
	})
	fmt.Printf("length is %d", len(data))
	fmt.Printf("size is %d", unsafe.Sizeof(data))
	result, _ := ListMapToJson(data)
	fmt.Println(result)
}
