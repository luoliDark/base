package sqlutils

import (
	"errors"
	"html/template"
	"reflect"
	"regexp"
	"strings"

	"github.com/luoliDark/base/util/commutil"
	"github.com/xormplus/xorm"
)

// 根据 sql,和输入的参数执行sql语句，并且返回执行结果map格式的
func ExecuteSqlV2(db *xorm.Session, sql string, params interface{}) ([]map[string]string, error) {

	params_ := []interface{}{}
	tmpl, err := template.New("mytemplate").Parse(sql)
	if err != nil {
		return nil, err
	}
	var result strings.Builder
	tmpl.Execute(&result, params)
	toMap := structToMap(params)
	sql = result.String()
	pattern := `#\{(.+?)}`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(sql, -1)
	sqlCopy := re.ReplaceAllString(sql, "?")
	sqlCopy = strings.Replace(sqlCopy, "&lt;", "<", -1)
	params_ = append(params_, sqlCopy)
	for _, match := range matches {
		params_ = append(params_, toMap[match[1]])
	}
	return db.QueryString(params_...)
}

// 根据 sql,和输入的参数,返回处理好的sql语句和参数
func SqlV2(sql string, params interface{}) (string, []interface{}, error) {

	params_ := []interface{}{}
	tmpl, err := template.New("mytemplate").Parse(sql)
	if err != nil {
		return "", nil, err
	}
	var result strings.Builder
	tmpl.Execute(&result, params)
	toMap := structToMap(params)
	sql = result.String()
	pattern := `#\{(.+?)}`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(sql, -1)
	sqlCopy := re.ReplaceAllString(sql, "?")
	params_ = append(params_, sqlCopy)
	for _, match := range matches {
		params_ = append(params_, toMap[match[1]])
	}
	i := params_[0]
	print(i)
	return commutil.ToString(i), params_[1:], nil
}

// ExecuteSqlV2Count 根据sql,和输入的参数执行sql语句，统计查询出来的数量
func ExecuteSqlV2Count(db *xorm.Session, sql string, params interface{}) (int, error) {

	params_ := []interface{}{}
	tmpl, err := template.New("mytemplate").Parse(sql)
	if err != nil {
		return 0, err
	}
	var result strings.Builder
	tmpl.Execute(&result, params)
	toMap := structToMap(params)
	sql = result.String()
	limitPattern := regexp.MustCompile(`(?s)limit.+`)
	sql = limitPattern.ReplaceAllString(sql, "")
	pattern := `#\{(.+?)}`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(sql, -1)
	// 创建正则表达式模式，匹配 "select * from" 到 "from" 之间的内容
	countPattern := regexp.MustCompile(`(?s)select.+?from`)
	// 替换匹配的部分为 "count(1)"
	sql = countPattern.ReplaceAllString(sql, "select count(1) from")
	sqlCopy := re.ReplaceAllString(sql, "?")
	params_ = append(params_, sqlCopy)
	for _, match := range matches {
		params_ = append(params_, toMap[match[1]])
	}
	queryString, err := db.QueryString(params_...)
	return commutil.ToInt(queryString[0]["count(1)"]), err
}

// ExecuteSqlV2Count 根据sql,和输入的参数执行sql语句，对sql进行分页查询
func ExecuteSqlV2Page(db *xorm.Session, sql string, params interface{}) ([]map[string]string, error) {

	params_ := []interface{}{}
	tmpl, err := template.New("mytemplate").Parse(sql)
	if err != nil {
		return nil, err
	}
	var result strings.Builder
	tmpl.Execute(&result, params)
	toMap := structToMap(params)
	sql = result.String()
	limitPattern := regexp.MustCompile(`(?s)limit.+`)
	sql = limitPattern.ReplaceAllString(sql, "")
	pattern := `#\{(.+?)}`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(sql, -1)
	// 创建正则表达式模式，匹配 "select * from" 到 "from" 之间的内容
	countPattern := regexp.MustCompile(`(?s)select.+?from`)
	// 替换匹配的部分为 "count(1)"
	sql = countPattern.ReplaceAllString(sql, "select count(1) from")
	sqlCopy := re.ReplaceAllString(sql, "?")
	limit := "limit ? offset ?"
	sqlCopy += limit
	params_ = append(params_, sqlCopy)
	for _, match := range matches {
		params_ = append(params_, toMap[match[1]])
	}
	pageSize := commutil.ToInt(toMap["PageSize"])
	pageIndex := commutil.ToInt(toMap["PageIndex"])
	if pageIndex == 0 {
		pageIndex = 1
	}
	params_ = append(params_, pageSize)
	params_ = append(params_, (pageIndex-1)*pageSize)
	queryString, err := db.QueryString(params_...)
	return queryString, err
}

// ExecuteSqlV2Sum 根据sql,和输入的参数执行sql语句，对sql语句进行汇总
func ExecuteSqlV2Sum(db *xorm.Session, sql string, params interface{}) ([]map[string]string, error) {

	params_ := []interface{}{}
	tmpl, err := template.New("mytemplate").Parse(sql)
	if err != nil {
		return nil, err
	}
	var result strings.Builder
	tmpl.Execute(&result, params)
	toMap := structToMap(params)
	sql = result.String()
	limitPattern := regexp.MustCompile(`(?s)limit.+`)
	sql = limitPattern.ReplaceAllString(sql, "")
	pattern := `#\{(.+?)}`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(sql, -1)
	// 创建正则表达式模式，匹配 "select * from" 到 "from" 之间的内容
	countPattern := regexp.MustCompile(`(?s)select.+?from`)
	// 替换匹配的部分为 "count(1)"
	sql = countPattern.ReplaceAllString(sql, "select count(1) from")
	sqlCopy := re.ReplaceAllString(sql, "?")
	params_ = append(params_, sqlCopy)
	for _, match := range matches {
		params_ = append(params_, toMap[match[1]])
	}
	queryString, err := db.QueryString(params_...)
	return queryString, err
}
func ExecuteSqlStruct(db *xorm.Session, sql string, params interface{}, target interface{}) error {
	data, err := ExecuteSqlV2(db, sql, params)
	if err != nil {
		return err
	}
	err = assignMapsToStructs(data, target)
	if err != nil {
		return err
	}
	return nil
}

func structToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	valueType := reflect.ValueOf(data)
	valueTypeKind := valueType.Kind()
	//reflect.
	if valueTypeKind == reflect.Struct {
		typeOfT := valueType.Type()

		for i := 0; i < valueType.NumField(); i++ {
			fieldName := typeOfT.Field(i).Name
			fieldValue := valueType.Field(i).Interface()
			result[fieldName] = fieldValue
		}
	}

	return result
}
func assignMapsToStructs(data []map[string]string, target interface{}) error {
	// 获取目标结构体类型
	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Ptr || targetType.Elem().Kind() != reflect.Slice {
		return errors.New("目标必须是指向切片的指针")
	}
	targetSliceType := targetType.Elem().Elem()
	if targetSliceType.Kind() != reflect.Struct {
		return errors.New("切片元素必须是结构体")
	}

	// 创建目标切片
	targetValue := reflect.ValueOf(target).Elem()
	targetSlice := reflect.MakeSlice(targetValue.Type(), len(data), len(data))

	for i, item := range data {
		// 创建目标结构体
		targetStruct := reflect.New(targetSliceType).Elem()

		for key, value := range item {
			if value == "" {
				continue
			}
			// key开头为大写
			key = strings.ToUpper(key[:1]) + key[1:]
			field := targetStruct.FieldByName(key)
			if !field.IsValid() {
				//return errors.New("结构体中不存在字段：" + key)
				continue
			}
			// 如果数组的值是切片
			// 试着把他转换成字符串
			kind := field.Type().Kind()
			if kind == reflect.String {
				field.SetString(value)
			} else if kind == reflect.Float64 {
				// 字符串转float64
				toFloat64 := commutil.ToFloat64(value)
				field.SetFloat(toFloat64)
			} else if kind == reflect.Float32 {
				toFloat32 := commutil.ToFloat32(value)
				field.SetFloat(float64(toFloat32))
			} else if kind == reflect.Int || kind == reflect.Int32 || kind == reflect.Int64 {
				toInt := commutil.ToInt(value)
				field.SetInt(int64(toInt))
			} else if kind == reflect.Int32 {
				toInt := commutil.ToInt(value)
				field.SetInt(int64(toInt))
			} else {
				field.SetString(value)
			}
		}

		targetSlice.Index(i).Set(targetStruct)
	}

	targetValue.Set(targetSlice)
	return nil
}
