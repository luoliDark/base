package arraysUtil

import (
	"strings"

	"github.com/luoliDark/base/util/commutil"
)

//数组去重
func RemoveRepeatedElementAndEmpty(arr []string) []string {
	newArr := make([]string, 0)
	for _, item := range arr {
		if "" == strings.TrimSpace(item) {
			continue
		}

		repeat := false
		if len(newArr) > 0 {
			for _, v := range newArr {
				if v == item {
					repeat = true
					break
				}
			}
		}

		if repeat {
			continue
		}

		newArr = append(newArr, item)
	}
	return newArr
}

//获取集合中某一字段的字符串集合
func GetListFieldResult(arr []map[string]string, field string) []string {
	newArr := make([]string, 0)
	for _, item := range arr {
		result := item[field]
		if commutil.IsNullOrEmpty(result) {
			continue
		}
		newArr = append(newArr, result)
	}
	return newArr
}

//合并字符串数组
func MergeStrArrays(arr ...[]string) []string {
	newArr := make([]string, 0)
	for _, item := range arr {
		for _, str := range item {
			newArr = append(newArr, str)
		}
	}
	return newArr
}

//合并字符串数组
func MergeAndDistinctStrArrays(arr ...[]string) []string {
	newArr := make([]string, 0)
	for _, item := range arr {
		for _, str := range item {
			newArr = append(newArr, str)
		}
	}
	return RemoveRepeatedElementAndEmpty(newArr)
}

//根据某一字段查找数组中匹配得值
func Filter(key, value string, arr []map[string]string) map[string]string {
	for _, item := range arr {
		s := item[key]
		if strings.EqualFold(value, s) {
			return item
		}
	}
	return nil
}
