package reflectutil

import (
	"errors"
	"reflect"
)

// 遍历struct并且自动进行赋值
func MapToStruct(data map[string]interface{}, inStructPtr interface{}) error {
	rType := reflect.TypeOf(inStructPtr)
	rVal := reflect.ValueOf(inStructPtr)
	if rType.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		return errors.New("inStructPtr must be ptr to struct")
	}
	// 遍历结构体
	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		// 得到tag中的字段名
		key := t.Tag.Get("key")
		if v, ok := data[key]; ok {
			// 检查是否需要类型转换
			dataType := reflect.TypeOf(v)
			structType := f.Type()
			if structType == dataType {
				f.Set(reflect.ValueOf(v))
			} else {
				if dataType.ConvertibleTo(structType) {
					// 转换类型
					f.Set(reflect.ValueOf(v).Convert(structType))
				} else {
					return errors.New(t.Name + " type mismatch")
				}
			}
		} else {
			return errors.New(t.Name + " not found")
		}
	}
	return nil
}

// 遍历struct生成map
func StructToMap(inStructPtr *interface{}) map[string]interface{} {
	mapDatas := make(map[string]interface{})
	elem := reflect.ValueOf(inStructPtr).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		mapDatas[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return mapDatas
}
