package sysmodel

import (
	"base/util/commutil"
	"bytes"
	"strings"
)

// 查询字段实体类，用于列表查询时传回用户输入的字段数据
type QueryField struct {
	FieldName   string `xorm:"fieldname" json:"fieldname"` //字段名称
	KeyWord     string `xorm:"keyword" json:"keyword"`     // 查询数据
	SqlDatatype string `xorm:"sqldatatype" json:"sqldatatype"`
	Op          string `xorm:"op" json:"op"` // 比较字符 in  like   注意使用 () 和 %
}

// 生成WhereSql
func GetWhereSqlByQueryField(Field []QueryField) string {
	if len(Field) == 0 {
		return ""
	}
	bf := bytes.Buffer{}
	idx := 0
	for _, field := range Field {
		if field.FieldName == "" {
			continue
		}
		if idx > 0 {
			bf.WriteString(" and ")
		}
		idx++
		bf.WriteString(field.FieldName)
		bf.WriteString(" ")
		bf.WriteString(field.Op)
		value := field.KeyWord
		Op := strings.ToLower(field.Op)
		if Op == "in" || Op == "not in" {
			valueList := strings.Split(value, ",")
			for idx, s := range valueList {
				if idx > 0 {
					bf.WriteString(",")
				} else {
					bf.WriteString(" (")
				}
				bf.WriteString(" '")
				bf.WriteString(s)
				bf.WriteString("' ")
			}
			bf.WriteString(")")
		} else if Op == "like" || Op == "not like" {
			bf.WriteString(" '%")
			bf.WriteString(value)
			bf.WriteString("%' ")
		} else {
			bf.WriteString(" '")
			bf.WriteString(value)
			bf.WriteString("' ")
		}
	}
	if bf.Len() > 0 {
		bf.WriteString(" ")
	}
	return bf.String()
}

// 生成 WhereSql ,替换为SQL占位符? 查询
//  参数： Field 请求字段，
//  参数： paramsMap 参数的占位Map，
// 返回 WhereSql a = ? and a like ?
func GetWhereSqlOrParamMapByQueryField(Field []QueryField, paramsMap map[string]interface{}) string {
	if len(Field) == 0 {
		return ""
	}
	if len(paramsMap) == 0 {
		paramsMap = make(map[string]interface{}, len(Field))
	}
	bf := bytes.Buffer{}
	index := 0
	for _, field := range Field {
		if field.FieldName == "" {
			continue
		}
		if index > 0 {
			bf.WriteString(" and ")
		}
		index++
		bf.WriteString(field.FieldName)
		bf.WriteString(" ")
		bf.WriteString(field.Op)
		bf.WriteString(" ")
		value := field.KeyWord
		Op := strings.ToLower(field.Op)
		if Op == "in" || Op == "not in" {
			bf.WriteString(" ( ")
			valueList := strings.Split(value, ",")
			for idx, s := range valueList {
				if idx > 0 {
					bf.WriteString(",")
				}
				// in 或 not in  需要替换为  id in (?id1,?id2,?id3)  map 中对应值为 id1:1,id2:2,id3:3
				key := commutil.AppendStr(field.FieldName, "_", idx)
				paramsMap[key] = s
				bf.WriteString(commutil.AppendStr("?", key))
			}
			bf.WriteString(" )")
		} else if Op == "like" || Op == "not like" {
			paramsMap[field.FieldName] = value
			bf.WriteString("?" + field.FieldName)
		} else {
			paramsMap[field.FieldName] = value
			bf.WriteString("?" + field.FieldName)
		}
	}
	if bf.Len() > 0 {
		bf.WriteString(" ")
	}
	return bf.String()
}
