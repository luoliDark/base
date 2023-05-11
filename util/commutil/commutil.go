//jsz by 2020.2.2 用于完成数据类型转换，和字符串拼接等

package commutil

import (
	"bytes"
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/satori/go.uuid"
	"github.com/valyala/fastjson"
	"net"
	"net/http"
	"paas/base/loghelper"
	"paas/base/sysmodel/logtype"
	"paas/base/util/snow"
	"regexp"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"
)

func GetUniqueId() string {
	worker := snow.GetInstance().GetWorker()
	return ToString(worker.NextId())
}

// 对多个字符串拼接为一个字符串
func AppendStr(str ...interface{}) string {

	var bufer bytes.Buffer

	for _, arg := range str {
		bufer.WriteString(ToString(arg))
	}

	return bufer.String()
}

//替换特殊字符
func ReplaceSpchar(val string) string {

	if strings.Contains(val, "\t") {
		val = strings.ReplaceAll(val, "\t", " ")
	}

	if strings.Contains(val, "\n") {
		val = strings.ReplaceAll(val, "\n", "\\n")
	}

	if strings.Contains(val, "\r") {
		val = strings.ReplaceAll(val, "\r", "\\r")
	}
	return val
}

// 转为Int型  注：根据操作系统是32 或64 自动转换
func ToInt(str interface{}) int {

	if str == nil || g.IsEmpty(str) {
		return 0
	}

	var re int
	switch vv := str.(type) {
	case string:
		re, _ = strconv.Atoi(vv)
	case int64, int, int16, int8, int32:
		re = gconv.Int(str)
	case float64, float32:
		re = gconv.Int(str)
	default:
		re = gconv.Int(str)
	}

	return re

}

// 转 64
func ToInt64(str interface{}) int64 {

	if str == nil || g.IsEmpty(str) {
		return 0
	}

	var re int64
	switch vv := str.(type) {
	case string:
		re, _ = strconv.ParseInt(vv, 10, 64)
	case int64:
		re = vv
	case float64:
		re = gconv.Int64(str)
	default:
		tmp, ok := str.(int64)
		if !ok {
			_, file, line, _ := runtime.Caller(2)
			fmt.Println(str, "转换为int64失败")
			loghelper.ByError("数据转换int64失败", ToString(vv)+file+ToString(line), "")
		} else {
			re = tmp
		}
	}

	return re

}

// 转为String
func ToString(str interface{}) string {

	if str == nil {
		return ""
	}

	var re string
	var ok bool
	switch vv := str.(type) {
	case string:
		re, ok = str.(string)
		if !ok {
			_, file, line, _ := runtime.Caller(2)
			fmt.Println(str, "转换为String失败")
			loghelper.ByError(logtype.DataTypeErr, file+ToString(line), "")
		}
	case *string:
		r, ok := str.(*string)
		if r != nil {
			re = *r
		}
		if !ok {
			_, file, line, _ := runtime.Caller(2)
			fmt.Println(str, "转换为String失败")
			loghelper.ByError(logtype.DataTypeErr, file+ToString(line), "")
			re = ""
		}
	case []uint8:
		//从redis取出来是这种类型
		var ba []byte
		for _, b := range vv {
			ba = append(ba, b)
		}
		re = string(ba)
	case int, int32:
		re = strconv.Itoa(str.(int))
	case int64:
		re = strconv.FormatInt(gconv.Int64(str), 10)
	case float32:
		re = strconv.FormatFloat(float64(vv), 'f', 6, 64)
	case float64:
		re = strconv.FormatFloat(vv, 'E', -1, 64) //float64
	case bool:
		re = strconv.FormatBool(vv)
	case time.Time:
		re = vv.Format("2006-01-02 15:04:05")
	case map[string]string:
		m, ok := str.(map[string]string)
		if !ok {
			re = ""
		} else {
			re = MapToStr(m)
		}
	default:
		re = fmt.Sprint(str)
	}

	return re
}

// 转为Float64
func ToFloat64(str interface{}) float64 {
	if str == nil || g.IsEmpty(str) {
		return 0
	}
	var re float64

	switch vv := str.(type) {
	case string:
		a, err := strconv.ParseFloat(vv, 64)
		if err != nil {
			_, file, line, _ := runtime.Caller(2)
			fmt.Println(str, "转换为ToFloat64失败")
			loghelper.ByError(logtype.DataTypeErr, ToString(vv)+file+ToString(line), "")
		}
		re = a
	case float64:
		re = vv
	case int, int8, int64, int16, int32:
		re = float64(ToInt(vv))
	default:
		a, err := strconv.ParseFloat(ToString(str), 64)
		if err != nil {
			_, file, line, _ := runtime.Caller(2)
			fmt.Println(str, "转换为ToFloat64失败")
			loghelper.ByError(logtype.DataTypeErr, ToString(vv)+file+ToString(line), "")
		}
		re = a
	}

	return re

}

// 转为Float32
func ToFloat32(str interface{}) float32 {

	if str == nil || g.IsEmpty(str) {
		return 0
	}

	var re float32

	switch vv := str.(type) {
	case string:
		a, err := strconv.ParseFloat(vv, 32)
		if err != nil {
			_, file, line, _ := runtime.Caller(2)
			fmt.Println(str, "转换为ToFloat32失败")
			loghelper.ByError(logtype.DataTypeErr, ToString(vv)+file+ToString(line), "")
		}
		re = float32((a))
	case float64:
		re = float32(vv)
	case int:
		re = float32(vv)
	default:
		a, err := strconv.ParseFloat(ToString(str), 32)
		if err != nil {
			_, file, line, _ := runtime.Caller(2)
			fmt.Println(str, "转换为ToFloat64失败")
			loghelper.ByError(logtype.DataTypeErr, ToString(vv)+file+ToString(line), "")
		}
		re = float32(a)
	}

	return re

}

// 删除数组指定位置元素
func RmEle(slice []interface{}, index int) []interface{} {
	return append(slice[:index], slice[index+1:])
}

//转换类型
func ToBool(value interface{}) bool {

	if value == nil {
		return false
	}

	result := false
	switch value.(type) {
	case bool:
		result = value.(bool)
	case int, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		//gconv.Bool()
		if gconv.Int(value) == 1 {
			result = true
		} else {
			result = false
		}

	case string:
		if value == "true" || value == "1" {
			result = true
		} else {
			result = false
		}
	default:
		val := ToString(value)
		result = val == "1" || val == "true"
	}

	return result
}

func GetUUID() string {
	// 36 位
	var uuid_str = uuid.NewV4().String()
	uuid_str = strings.ReplaceAll(uuid_str, "-", "")
	return uuid_str
}

// byte 数组转string
func B2S(bs []byte) string {
	bys := new(bytes.Buffer)
	for _, b := range bs {
		bys.WriteString(strconv.Itoa(int(b)))
	}
	return bys.String()
}

const Time_Fomat01 = "2006-01-02 15:04:05"
const Time_Fomat02 = "2006/01/02 15:04:05"
const Time_Fomat03 = "2006-01-02"
const Time_Fomat04 = "2006/01/02"
const Time_Fomat05 = "15:04:05"
const Time_Fomat06 = "20060102150405"
const Time_Fomat07 = "20060102"
const Time_Fomat08 = "20060102150405.000"
const Time_Fomat09 = "150405"
const Time_Fomat10 = "060102150405"
const Time_Fomat12 = "2006-01-02 15:04:05.000" //时分秒毫秒

// 时间格式化
func TimeFormat(date time.Time, formater string) string {
	// 时间格式转换
	return time.Unix(date.Unix(), 0).Format(formater)
}

// 获取 客户端请求 ip 地址
func GetRequestIP(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("Remote_addr"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

func GetLangCode() (langCode string) {
	return ""

}

func ForListFindElement(list *list.List, index int) (obj *list.Element) {
	var i = 0
	for element := list.Front(); nil != element; element = element.Next() {
		if i == index {
			return element
		}
		i++
	}
	return obj
}

func IsNullOrEmpty(obj interface{}) bool {
	if "" == obj || nil == obj {
		return true
	}
	return false
}

func CheckClassType(obj string, classType string) bool {
	if "" == obj || "" == classType {
		//如果传入值为空，也就是数据库未配置，那么就认为验证正确
		return true
	}
	//正则表达式
	var pattern = ""

	//防止大小写问题，全部先转换为小写
	classType = strings.ToLower(classType)

	if strings.Index(classType, "int") > -1 {
		//数字类型，验证是否全部为数字
		pattern = "^-?\\d+$"
		reg, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		return reg.MatchString(obj)

	} else if strings.Index(classType, "decimal") > -1 {
		//小数类型，验证是否全部为数字，
		pattern = "^(-?\\d+)(\\.\\d+)?$"
		reg, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		return reg.MatchString(obj)
	} else if strings.Index(classType, "date") > -1 {
		//时间格式
		return true
	} else if strings.Index(classType, "float") > -1 {
		pattern = "^(-?\\d+)(\\.\\d+)?$"
		reg, error := regexp.Compile(pattern)
		if error != nil {
			panic(error)
		}
		return reg.MatchString(obj)
	}
	return true
}

//从list中查找指定数据的下标
func Find(slice []string, val interface{}) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

//将生成的json封装成resultbean 格式的json返回
func GetResultBeanByJson(jsonStr string) string {

	if jsonStr == "" {
		beanJson := AppendStr("{\"IsSuccess\":false,\"ErrorCode\":\"\",\"ErrorMsg\":\"\",\"ResultData\":\"\"}")
		return beanJson
	} else {
		beanJson := AppendStr("{\"IsSuccess\":true,\"ErrorCode\":\"\",\"ErrorMsg\":\"\",\"ResultData\":", jsonStr, "}")
		return beanJson
	}
}

//rows转map 注只限定查出两名
func RowsToMapVsMap(rows []map[string]string, key string) map[string]map[string]string {
	resultM := make(map[string]map[string]string)
	if len(rows) == 0 {
		return make(map[string]map[string]string)
	} else {
		for _, m := range rows {
			k := m[key]
			if k != "" {
				resultM[k] = m
			}
		}
	}
	return resultM
}

//rows转map[string]map[string]string
func RowsToMapList(key string, rows []map[string]string) map[string]map[string]string {
	resultM := make(map[string]map[string]string)
	if len(rows) == 0 {
		return resultM
	} else {
		for _, m := range rows {
			k := m[key]
			if k != "" {
				resultM[k] = m
			}
		}
	}
	return resultM
}

//rows转map 注只限定查出两名并且别名必须为k,v
func RowsToMap(rows []map[string]string) map[string]string {
	return RowsToMapByKeys(rows, "k", "v")
}

//rows转map 注只限定查出两名
func RowsToMapByKeys(rows []map[string]string, key, value string) map[string]string {
	resultM := make(map[string]string)
	if len(rows) == 0 {
		return resultM
	} else {
		for _, m := range rows {
			k := m[key]
			v := m[value]
			if k != "" {
				resultM[k] = v
			}
		}
	}
	return resultM
}

//rows转map 注只限定查出两名并且别名必须为k,v
func RowsToSyncMap(rows []map[string]string) sync.Map {
	resultM := sync.Map{}
	if len(rows) > 0 {
		for _, m := range rows {
			key := m["k"]
			value := m["v"]
			if key != "" {
				resultM.Store(key, value)
			}
		}
	}
	return resultM

}

func GetSyncMapValue(mapObj sync.Map, key string) interface{} {
	if !IsNullOrEmpty(mapObj) {
		value, ok := mapObj.Load(key)
		if !ok {
			panic("map取值失败")
		}
		return value
	}
	return ""
}

//rows 将指定字段转为以豆号分隔的字符串
func RowsToIdStrByCol(rows []map[string]string, splitCol string) string {

	var ids strings.Builder
	if rows == nil {
		return ""
	} else {
		for _, m := range rows {
			key := m[splitCol]
			if key != "" {
				ids.WriteString("'")
				ids.WriteString(key)
				ids.WriteString("',")
			}
		}
	}

	return strings.TrimRight(ids.String(), ",")

}

//rows 将指定数组转为以豆号分隔的字符串
func RowsToIdStrByArr(rows []string) string {

	var ids strings.Builder
	if rows == nil {
		return ""
	} else {
		for _, key := range rows {
			if key != "" {
				ids.WriteString("'")
				ids.WriteString(key)
				ids.WriteString("',")
			}
		}
	}

	return strings.TrimRight(ids.String(), ",")

}

//将 1,2,3 转为 ‘1’，‘2’，‘3’
func StringsToIdStr(ids string) string {
	return RowsToIdStrByArr(strings.Split(ids, ","))
}

//将1，2，3 转为[1,2,3]
func IdsToArray(ids string, splitChar string) []string {
	return strings.Split(ids, splitChar)
}

//将map的key转为 ‘1’，‘2’，‘3’
func MapToIdStr(m map[string]string) string {
	var ids bytes.Buffer
	for key, _ := range m {
		if key != "" {
			ids.WriteString("'")
			ids.WriteString(key)
			ids.WriteString("',")
		}
	}
	return strings.TrimRight(ids.String(), ",")
}

//将map中的k,v转为字符串 ，用行记录map中数据对日志
func MapToStr(m map[string]string) string {

	if m == nil {
		return ""
	}

	var ids bytes.Buffer
	for key, v := range m {
		if key != "" {
			ids.WriteString(key)
			ids.WriteString("=")
			ids.WriteString(v)
			ids.WriteString(";")
		}
	}
	return strings.TrimRight(ids.String(), ";")
}

func GetTableFields(tableName string, tablebm string) (fields string) {
	if IsNullOrEmpty(tableName) {
		return fields
	}
	var fieldBuilder strings.Builder
	tableName = strings.ToLower(tableName)
	if !IsNullOrEmpty(tablebm) {
		tablebm = tablebm + "."
	}
	switch tableName {
	case "sys_fgridfield":
		fieldBuilder.WriteString("" + tablebm + "name as showname," + tablebm + "SqlCol as sqlcol," + tablebm + "DataSource as datasource," + tablebm + "IsConvertByCode as isconvertbycode,")
		fieldBuilder.WriteString("" + tablebm + "IsConvertByName as isconvertbyname," + tablebm + "IsSingle as issingle,")
		//" + tablebm + "IsRequired as isrequired,  暂时木有isrequired字段
		fieldBuilder.WriteString("" + tablebm + "SqlDataType as sqldatatype," + tablebm + "Select_Show as select_show")
		break
	case "sys_fpagefield":
		fieldBuilder.WriteString("" + tablebm + "name as showname," + tablebm + "SqlCol as sqlcol," + tablebm + "DataSource as datasource," + tablebm + "IsConvertByCode as isconvertbycode,")
		//" + tablebm + "IsRequired as isrequired,
		fieldBuilder.WriteString("" + tablebm + "IsConvertByName as isconvertbyname," + tablebm + "IsSingle as issingle,")
		fieldBuilder.WriteString("" + tablebm + "SqlDataType as sqldatatype," + tablebm + "Select_Show as select_show")
		break
	}
	return fieldBuilder.String()
}

func FormatDateTime(dateparam time.Time) string {
	if IsNullOrEmpty(dateparam) {
		return ""
	}
	return dateparam.Format(Time_Fomat01)
}

// time 指定格式转 string
func GetNowTimeByFormatStr(Format string) string {
	return time.Now().Format(Format)
}

//将时间戳转换为日期 针对毫秒级
func TimeSpanToDate(timeUnix int64) time.Time {
	tu := int64(timeUnix / 1000)

	t := time.Unix(tu, 0) //先转为秒再转日期

	return t
}

//将时间戳转换为日期
func StrToDateTime(t1 string) time.Time {
	time, err := time.ParseInLocation("2006-01-02 15:04:05", t1, time.Local)
	if err != nil {
		loghelper.ByError("将时间戳转换为日期失败", err.Error()+" 转换的字符串="+t1, "")
	}
	return time
}

//获取rest层捕到的全部服务器错误
func GetHttpCatchErrMsg(err interface{}, userId string) string {
	errbuf := bytes.Buffer{}
	stack := ToString(debug.Stack())
	errbuf.WriteString(fmt.Sprint(err)) // error 和 string 等类型 都能获取到对应值。
	errstr := errbuf.String()

	loghelper.ByRestCatchErr("自动捕获到异常，", AppendStr(errstr, "; 源代码文件 : ", stack), userId)
	return errstr
}

//将对象转为json
func ObjectToJson(obj interface{}) string {
	js3, err := json.Marshal(obj)
	if err != nil {
		return "convert err"
	} else {
		return ToString(js3)
	}
}

//将对象转为json
func ToJson(obj interface{}) string {
	js3, err := json.Marshal(obj)
	if err != nil {
		return "convert err"
	} else {
		return ToString(js3)
	}
}

func IsIntType(sqlDataType string) bool {
	if IsNullOrEmpty(sqlDataType) {
		return false
	}
	//////设置数据类型为小写字母
	if !IsNullOrEmpty(sqlDataType) {
		sqlDataType = strings.ToLower(sqlDataType)
	}

	return sqlDataType == ("number") || strings.Contains(sqlDataType, "int") || strings.HasPrefix(sqlDataType, "decimal(")
}

/**
根据传入map[string]string转换为map[string]interface
*/
func TypeConverter(datalist []map[string]string) []map[string]interface{} {
	reslt, _ := json.Marshal(datalist)
	var resultdata interface{}
	json.Unmarshal(reslt, resultdata)
	return resultdata.([]map[string]interface{})
}

func TypeConverterMapString(datalist map[string]interface{}) map[string]string {
	reslt, _ := json.Marshal(datalist)
	var resultdata interface{}
	json.Unmarshal(reslt, resultdata)
	return resultdata.(map[string]string)
}

/**
根据传入时间获取前一年时间戳
*/
func GetNowTimeBeforeYear(date time.Time) int64 {
	///后去当前时间的前一年时间
	beforeYearTime := date.AddDate(-2, 0, 0)
	///获取时间戳
	return beforeYearTime.Unix()
}

//捕获ajax请求后台服务器错误
func CatchError() {
	if err := recover(); err != nil {
		GetHttpCatchErrMsg(err, "-1")
	}
}

//捕获ajax请求后台服务器错误
func CatchErrorByTitle(title string) {
	if err := recover(); err != nil {
		GetHttpCatchErrMsg(fmt.Sprintf("title:%v,err:%v", title, err), "-1")
	}
}

func If(condition bool, trueVal, falseVal interface{}) string {
	if condition {
		return ToString(trueVal)
	}
	return ToString(falseVal)
}

func AssmblyMap(valuemap map[string]string, key string, valueObj *fastjson.Value) map[string]string {
	var value = valueObj.Get(key)
	//println("map获取的值："+key+"=============="+ToString(value))
	if nil != value && !IsNullOrEmpty(value) {
		var result = value.String()
		//println("map获取到的值"+result)
		if result == "true" {
			result = "1"
		} else if result == "false" {
			result = "0"
		}
		//log.Print("map获取的值："+key+"=============="+ToString(value))
		//println("map获取到的值：" + result)
		valuemap[key] = strings.Trim(result, "\"")
	}
	return valuemap
}

func TrimStr(val string) string {
	if !IsNullOrEmpty(val) {
		val = strings.Trim(val, "\"")
	}
	return val
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func ListConverterMap(datamaplist []map[string]string, key string) map[string]map[string]string {
	var resultmap = make(map[string]map[string]string)
	///判断集合是否为空
	for _, mapobj := range datamaplist {
		///获取对应key
		var value = mapobj[key]
		if IsNullOrEmpty(value) {
			continue
		}
		resultmap[ToString(value)] = mapobj
	}
	return resultmap
}

//获取相差时间
func GetHourDiffer(start_time, end_time string) int64 {
	var hour int64
	t1, err := time.ParseInLocation(Time_Fomat01, start_time, time.Local)
	t2, err := time.ParseInLocation(Time_Fomat01, end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

func ParentIDIsNull(parentID string) bool {
	if parentID == "" || parentID == "0" || parentID == "undefined" || parentID == "null" {
		return true
	}
	return false
}

//通过数组获取key值，组成in条件语句
func GetListCodeStrInWhere(list []map[string]string, key string, column string) string {
	codeStr := ""
	if len(list) <= 0 {
		return codeStr
	}
	for _, object := range list {
		code := object[key]
		if !IsNullOrEmpty(code) {
			codeStr += "'" + code + "',"
		}
	}
	if !IsNullOrEmpty(codeStr) {
		codeStr = strings.TrimRight(codeStr, ",")
		codeStr = column + " in (" + codeStr + ")"
	}
	return codeStr
}

//通过map键的唯一性去重
func RemoveRepeatedElement(s []string) []string {
	result := make([]string, 0)
	m := make(map[string]bool) //map的值不重要
	for _, v := range s {
		if _, ok := m[v]; !ok {
			if !IsNullOrEmpty(v) {
				result = append(result, v)
				m[v] = true
			}
		}
	}
	return result
}

//获取当前 时间
func GetNowTime() string {
	dateparam := time.Now()
	return dateparam.Format(Time_Fomat01)
}

//获取当前 时间
func GetNowYYDDMM() string {
	dateparam := time.Now()
	return dateparam.Format(Time_Fomat03)
}

//获取当前年份
func GetNowYear() int {
	return time.Now().Year()
}

//获取当前月份
func GetNowMonth() int {
	return time.Now().Year()
}

//转换时间 例：2021-09-08 23：22：44
func ParseTime(timeStr string) time.Time {
	t, err := time.Parse(Time_Fomat01, timeStr)
	if err != nil {
		loghelper.ByError("转换为时间失败", err.Error(), "")
	}
	return t
}

//转换时间 例：2021-09-08 23：22：44
func TimeToStr(t time.Time) string {
	return t.Format(Time_Fomat01)
}

// 获取两个时间相差的天数，0表同一天，正数表endDay>startDay，负数表endDay<startDay
func GetDiffDays(startDay, endDay time.Time) int {

	endDay = time.Date(endDay.Year(), endDay.Month(), endDay.Day(), 0, 0, 0, 0, time.Local)
	startDay = time.Date(startDay.Year(), startDay.Month(), startDay.Day(), 0, 0, 0, 0, time.Local)

	return int(endDay.Sub(startDay).Hours() / 24)

}
