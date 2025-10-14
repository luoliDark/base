//jsz by 2020.2.2 用于完成数据类型转换，和字符串拼接等

package commutil

import (
	"bytes"
	"container/list"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/sysmodel/logtype"
	"github.com/luoliDark/base/util/snow"
	uuid "github.com/satori/go.uuid"
)

// 转为千分位
func NumToQianFW(num interface{}) string {

	str := ToString(num)

	arr := strings.Split(str, ".")
	numStr := arr[0] //如果有小数获取整数部分
	var fnm string
	if len(arr) > 1 {
		fnm = arr[1]
	}
	length := len(numStr)
	if length < 4 {

		if !g.IsEmpty(fnm) {
			fnm = strings.TrimRight(fnm, "0") //去除小数点后的00
			if !g.IsEmpty(fnm) {
				numStr += "." + fnm
			}
		}

		return numStr
	}

	count := (length - 1) / 3 //取于-有多少组三位数
	for i := 0; i < count; i++ {
		s1 := numStr[:length-(i+1)*3]
		s2 := numStr[length-(i+1)*3:]
		numStr = s1 + "," + s2
	}

	if !g.IsEmpty(fnm) {
		fnm = strings.TrimRight(fnm, "0") //去除小数点后的00
		if !g.IsEmpty(fnm) {
			numStr += "." + fnm
		}
	}
	return numStr
}

func GetUniqueId() string {
	worker := snow.GetInstance().GetWorker()
	//worker, err := snow.NewWorker(1)
	//if err != nil {
	//	fmt.Print(err.Error())
	//	return GetUUID()
	//}
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

// 转为Int型  注：根据操作系统是32 或64 自动转换
func ToInt(str interface{}) int {

	if str == nil || g.IsEmpty(str) {
		return 0
	}

	var re int
	switch vv := str.(type) {
	case string:
		re, _ = strconv.Atoi(vv)
	case int64:
		re = gconv.Int(str)
	case float64:
		re = gconv.Int(str)
	case bool:
		b := gconv.Bool(str)
		if b {
			re = 1
		} else {
			re = 0
		}
	default:
		re = gconv.Int(str)
	}

	return re

}

// 转为Int型  注：根据操作系统是32 或64 自动转换
func ToInt64(str interface{}) int64 {

	if str == nil || g.IsEmpty(str) {
		return 0
	}

	var re int64
	switch vv := str.(type) {
	case string:
		tmp, _ := strconv.Atoi(vv)
		re = gconv.Int64(tmp)
	case int64:
		re = gconv.Int64(str)
	case float64:
		re = gconv.Int64(str)
	case bool:
		b := gconv.Bool(str)
		if b {
			re = 1
		} else {
			re = 0
		}
	default:
		re = gconv.Int64(str)
	}

	return re

}

func ToTime(str interface{}) time.Time {
	if IsNullOrEmpty(str) {
		return time.Now()
	}
	time, e := time.Parse("2006-01-02", ToString(str))
	if e != nil {
		fmt.Println(str, "转换为Time失败")
		panic("转换日期失败，" + e.Error())
	}
	return time
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
		re = strconv.FormatFloat(vv, 'f', 6, 64)
		//re = strconv.FormatFloat(vv, 'E', -1, 64) //float64  这样转出来会有E字母
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

	case []string:

		var sbStr strings.Builder
		for _, b := range vv {
			sbStr.WriteString(b)
		}
		re = sbStr.String()

	default:
		re = fmt.Sprint(str)
		if !ok {
			_, file, line, _ := runtime.Caller(2)
			fmt.Println(str, "转换为string失败")
			loghelper.ByError(logtype.DataTypeErr, file+ToString(line), "")
		}
	}

	return re
}

// 转保留2位小数
func ToF2(value interface{}) float64 {

	newV, err := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	if err != nil {
		newV = ToFloat64(value)
	}
	return newV

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
	case int:
		re = float64(vv)
	default:
		a, err := strconv.ParseFloat(ToString(vv), 64)
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
	case float64:
		re = float32(vv)
	default:
		a, err := strconv.ParseFloat(ToString(str), 32)
		if err != nil {
			_, file, line, _ := runtime.Caller(2)
			fmt.Println(str, "转换为ToFloat32失败")
			loghelper.ByError(logtype.DataTypeErr, ToString(vv)+file+ToString(line), "")
		}
		re = float32((a))
	}
	return re

}

// 删除数组指定位置元素
func RmEle(slice []interface{}, index int) []interface{} {
	return append(slice[:index], slice[index+1:])
}

func IsTrue(value interface{}) bool {
	return ToBool(value)
}
func IsFalse(value interface{}) bool {
	return !ToBool(value)
}

// 转换类型
func ToBool(value interface{}) bool {

	if value == nil {
		return false
	}

	result := false
	switch value := value.(type) {
	case bool:
		result = value
	case int, int8, int16, int32, int64, uint8, uint16, uint32, uint64:

		if value == 1 {
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

// GenerateUniqueShortIDAndKeyWord 生成指定长度的唯一键
func GenerateUniqueShortIDAndKeyWord(keyword string, length int) string {
	timestamp := time.Now().UnixNano()           // 获取当前时间戳（纳秒）
	random := rand.Int63()                       // 获取一个随机数（int64）
	id := ToString(timestamp) + ToString(random) // 将结果转换为16进制字符串
	var keyWrodLength = 0
	if !IsNullOrEmpty(keyword) {
		keyWrodLength = len(keyword)
	}
	length = length - keyWrodLength
	if len(id) > length { // 如果结果超过了指定长度，则截取
		id = id[:length]
	} else if len(id) < length { // 如果结果短于指定长度，则填充0（可选）
		id = fmt.Sprintf("%0"+strconv.Itoa(length)+"s", id)
	}
	return keyword + id
}

func GenerateUniqueShortID(length int) string {
	timestamp := time.Now().UnixNano() // 获取当前时间戳（纳秒）
	random := rand.Int63()             // 获取一个随机数（int64）
	combined := timestamp ^ random     // 结合时间戳和随机数，确保唯一性（XOR操作）
	id := fmt.Sprintf("%x", combined)  // 将结果转换为16进制字符串
	if len(id) > length {              // 如果结果超过了指定长度，则截取
		id = id[:length]
	} else if len(id) < length { // 如果结果短于指定长度，则填充0（可选）
		id = fmt.Sprintf("%0"+strconv.Itoa(length)+"s", id)
	}
	return id
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
const Time_Fomat11 = "2006-01"
const Time_Fomat12 = "2006-01-02 15:04:05.000" //时分秒毫秒

// 时间格式化
func TimeFormat(date time.Time, formater string) string {

	if formater == Time_Fomat11 {
		s := time.Unix(date.Unix(), 0).Format(Time_Fomat03)
		arr := strings.Split(s, "-")
		if len(arr) >= 2 {
			return arr[0] + "-" + arr[1]
		} else {
			return ""
		}

	} else {
		// 时间格式转换
		return time.Unix(date.Unix(), 0).Format(formater)
	}

}

// 时间格式化
func TimParse(date string, formater string) time.Time {

	date = strings.ReplaceAll(date, "\t", "")
	date = strings.ReplaceAll(date, "\r", "")
	date = strings.ReplaceAll(date, "\n", "")

	// 时间格式转换
	timeval, e := time.Parse(formater, date)
	if nil != e {
		panic("时间转换失败：" + e.Error() + "传入的时间是：" + date)
	}

	return timeval
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

func ArrayToString(arrays []string) string {
	var result = ""
	//遍历数组拼接字符串
	for _, v := range arrays {
		result += v + ","
	}
	//判断是否为空
	if !IsNullOrEmpty(result) {
		result = result[0 : len(result)-1]
	}
	return result
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

// 从字符串找提出纯数字
func GetNumberFromStr(str string) string {
	reg := regexp.MustCompile(`[0-9]`)
	s2 := reg.FindAllString(str, -1)
	return strings.Join(s2, "")
}

func IsNullOrEmpty(obj interface{}) bool {
	if "" == obj || nil == obj {
		return true
	}
	return false
}

func IsNotNullOrEmpty(obj interface{}) bool {
	return !IsNullOrEmpty(obj)
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

// 从list中查找指定数据的下标
func Find(slice []string, val interface{}) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// 将生成的json封装成resultbean 格式的json返回
func GetResultBeanByJson(jsonStr string) string {

	if jsonStr == "" {
		beanJson := AppendStr("{\"IsSuccess\":false,\"ErrorCode\":\"\",\"ErrorMsg\":\"\",\"ResultData\":\"\"}")
		return beanJson
	} else {
		beanJson := AppendStr("{\"IsSuccess\":true,\"ErrorCode\":\"\",\"ErrorMsg\":\"\",\"ResultData\":", jsonStr, "}")
		return beanJson
	}
}

// rows转map[string]map[string]string
func RowsToMapVsMap(rows []map[string]string, key string) map[string]map[string]string {
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

// rows转map[string]map[string]string
func RowsToMapList(key string, rows []map[string]string) map[string]map[string]string {
	resultM := make(map[string]map[string]string)
	if len(rows) == 0 {
		return resultM
	} else {
		for _, m := range rows {
			k := m[key]
			if k != "" {
				_, ok := resultM[k]
				if ok {
					//相同的key  并且已经删除则跳过
					if m["isdiscard"] == "1" {
						continue
					}
					resultM[k] = m
				} else {
					resultM[k] = m
				}
			}

		}
	}
	return resultM
}

// rows转map 注只限定查出两名并且别名必须为k,v
func RowsToMap(rows []map[string]string) map[string]string {

	resultM := make(map[string]string)
	if rows == nil {
		return make(map[string]string)
	} else {
		for _, m := range rows {
			key := m["k"]
			value := m["v"]
			if key != "" {
				resultM[key] = value
			}
		}
	}

	return resultM

}

// rows转map 注只限定查出两名并且别名必须为k,v
func RowsToSyncMap(rows []map[string]string) sync.Map {

	resultM := sync.Map{}
	if rows == nil {
		return sync.Map{}
	} else {
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

// rows 将指定字段转为以豆号分隔的字符串
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

// rows 某个字段转为数组
func RowsToArr(rows []map[string]string) []string {

	if rows == nil {
		return nil
	} else {

		lst := make([]string, len(rows))
		for index, m := range rows {
			key := m["id"]
			lst[index] = key
		}

		return lst

	}

}

// rows 将指定数组转为以豆号分隔的字符串
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

// 将 1,2,3 转为 ‘1’，‘2’，‘3’
func StringsToIdStr(ids string) string {
	return RowsToIdStrByArr(strings.Split(ids, ","))
}

// 将map的key转为 ‘1’，‘2’，‘3’
func MapToIdStr(m map[string]string) string {
	if m == nil {
		return ""
	}

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

// 将map中的k,v转为字符串 ，用行记录map中数据对日志
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
func TimeFormatStr(Format string) string {
	return time.Now().Format(Format)
}

// 获取当前 时间
func GetNowTime() string {
	dateparam := time.Now()
	return dateparam.Format(Time_Fomat01)
}

// 获取当前 时间 2006-01-02
func GetNowYYDDMM() string {
	dateparam := time.Now()
	return dateparam.Format(Time_Fomat03)
}

// 获取当前 时间 20060102150405
func GetNowYYYYMMDDHHmmss() string {
	dateparam := time.Now()
	return dateparam.Format(Time_Fomat06)
}

func GetNowByFormat(layout string) string {
	dateparam := time.Now()
	return dateparam.Format(layout)
}

// 获取当前 时间 20060102
func GetNowYYYYMMDD() string {
	dateparam := time.Now()
	return dateparam.Format(Time_Fomat07)
}

// 获取当前 时间 含毫秒
func GetNowYYYYMMDDHHMMSSsss() string {
	dateparam := time.Now()
	format := dateparam.Format(Time_Fomat08)
	replace := strings.Replace(format, ".", "", 1)
	return replace
}

// 获取当前年份
func GetNowYear() int {
	return time.Now().Year()
}

// 获取当前月份
func GetNowMonth() int {
	return int(time.Now().Month())
}

// 获取rest层捕到的全部服务器错误
func GetHttpCatchErrMsg(err interface{}, userId string) string {
	errstr := "内部错误"
	if s := GetCatchErrMsg(err); s != "" {
		errstr += s
	}
	stack := ToString(debug.Stack())
	loghelper.ByRestCatchErr("rest请求错误", errstr+" 源代码文件："+stack, userId)
	return errstr
}

// 获取rest层捕到的全部服务器错误
func GetCatchErrMsg(err interface{}) string {
	errType := reflect.TypeOf(err).Name()
	errstr := ""
	if errType == "error" {
		newErr := err.(error)
		if newErr != nil {
			errstr += newErr.Error()
		}
	} else if errType == "string" {
		errstr += ToString(err)
	} else {
		errstr = fmt.Sprint(err)
	}
	return errstr
}

// 将对象转为json
func ObjectToJson(obj interface{}) string {
	js3, err := json.Marshal(obj)
	if err != nil {
		return "convert err"
	} else {
		return ToString(js3)
	}
}

type reflectWithString struct {
	v reflect.Value
	s string
}

func StructToJSONStr(obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	//var data = make(map[string]interface{})
	sv := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		sv[i] = t.Field(i).Name
		//data[t.Field(i).Name] = v.Field(i).Interface()
	}
	sort.Slice(sv, func(i, j int) bool { return sv[i] < sv[j] })
	var result bytes.Buffer
	result.WriteString("{")
	for i, key := range sv {

		field := v.FieldByName(key)
		fieldValue := field.Interface()
		switch field.Kind() {
		case reflect.Struct:
			result.WriteString("\"" + key + "\":")
			structResult := StructToJSONStr(fieldValue)
			result.WriteString(structResult)
		case reflect.Slice:
			result.WriteString("\"" + key + "\":[")
			//array := fieldValue.([]interface{})
			of := reflect.ValueOf(fieldValue)
			for i := 0; i < of.Len(); i++ {
				index := of.Index(i)
				structObj := index.Interface()
				result.WriteString("\"" + key + "\":")
				structResult := StructToJSONStr(structObj)
				result.WriteString(structResult)
				if i < of.Len()-1 {
					result.WriteString(",")
				}
			}
			//for _,obj := range array{
			//	result.WriteString("\""+key+"\":")
			//	structResult := StructToMap(obj)
			//	result.WriteString(structResult)
			//}
			result.WriteString("]")
		default:
			result.WriteString("\"" + key + "\":\"" + ToString(fieldValue) + "\"")
		}
		if i < len(sv)-1 {
			result.WriteString(",")
		}
	}
	result.WriteString("}")

	return result.String()
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

/*
*
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

/*
*
根据传入时间获取前一年时间戳
*/
func GetNowTimeBeforeYear(date time.Time) int64 {
	///后去当前时间的前一年时间
	beforeYearTime := date.AddDate(-2, 0, 0)
	///获取时间戳
	return beforeYearTime.Unix()
}

// 捕获ajax请求后台服务器错误
func CatchError() {
	if err := recover(); err != nil {
		errMsg := GetHttpCatchErrMsg(err, "-1")
		loghelper.ByError("自动捕获到异常", errMsg, "-1")
	}
}

func CatchError2(errType string) {
	if err := recover(); err != nil {
		errMsg := GetHttpCatchErrMsg(err, "-1")
		loghelper.ByError(errType+"自动捕获到异常", errMsg, "-1")
	}
}

func If(condition bool, trueVal, falseVal interface{}) string {
	if condition {
		return ToString(trueVal)
	}
	return ToString(falseVal)
}

func TrimStr(val string) string {
	if !IsNullOrEmpty(val) {
		val = strings.Trim(val, "\"")
	}
	return val
}

// rows 将指定xorm结构体转为以逗号分隔的字符串
func ColByStruct(data interface{}) string {
	t := reflect.TypeOf(data)
	var str = ""
	for i := 0; i < t.NumField(); i++ {
		str += t.Field(i).Tag.Get("xorm") + ","
	}
	return strings.TrimRight(str, ",")
}

// MD5 加密
func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

// 获取两个时间相差的天数，0表同一天，正数表endDay>startDay，负数表endDay<startDay
func GetDiffDays(startDay, endDay time.Time) int {

	endDay = time.Date(endDay.Year(), endDay.Month(), endDay.Day(), 0, 0, 0, 0, time.Local)
	startDay = time.Date(startDay.Year(), startDay.Month(), startDay.Day(), 0, 0, 0, 0, time.Local)

	return int(endDay.Sub(startDay).Hours() / 24)

}

// 计算日期相差多少月
func GetSubMonth(startDay, endDay time.Time) (month int) {
	y1 := endDay.Year()
	m1 := int(endDay.Month())
	d1 := endDay.Day()

	y2 := startDay.Year()
	m2 := int(startDay.Month())
	d2 := startDay.Day()

	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	return month + 1
}

// 获取字符串前几位
func SubString(str string, startIndx, endIndex int) string {
	rs := []rune(str)
	d2 := string(rs[startIndx:endIndex])
	return d2
}

// 合并map
func MergeMap(mObj ...map[string]string) map[string]string {
	newObj := map[string]string{}
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}

func ToJson(obj interface{}) string {
	js3, err := json.Marshal(obj)
	if err != nil {
		return "convert err"
	} else {
		return ToString(js3)
	}
}

// 合并map
func MergeList(mObj ...[]map[string]string) []map[string]string {
	newObj := []map[string]string{}
	for _, m := range mObj {
		for _, object := range m {
			newObj = append(newObj, object)
		}
	}
	return newObj
}

// TraverseMapInStringOrder 按字母顺序遍历map
func TraverseMapInStringOrder(params map[string]string) (map[string][]string, string, map[string][]string) {
	str := strings.Builder{}
	var handler = make(map[string][]string)
	var appidHandler = make(map[string][]string)
	//第一个参数放sign
	appidHandler["sign"] = []string{params[""]}
	keys := make([]string, 0)
	for k, _ := range params {
		keys = append(keys, k)
		appidHandler[k] = []string{params[k]}
	}
	sort.Strings(keys)
	for _, k := range keys {
		handler[k] = []string{params[k]}
		str.WriteString(k)
		str.WriteString("=")
		str.WriteString(params[k])
		str.WriteString("&")
	}
	return handler, str.String(), appidHandler
}

// CommonTraverseMapInStringOrder 按字母顺序遍历map
func CommonTraverseMapInStringOrder(params map[string]string) (map[string][]string, string, map[string][]string, []string) {
	str := strings.Builder{}
	var handler = make(map[string][]string)
	var appidHandler = make(map[string][]string)
	//先对key进行排序
	keys := make([]string, 0)
	for k, _ := range params {
		keys = append(keys, k)
		appidHandler[k] = []string{params[k]}
	}
	sort.Strings(keys)
	//将排序后的key再次进行拼接成字符串返回
	for _, k := range keys {
		handler[k] = []string{params[k]}
		str.WriteString(k)
		str.WriteString("=")
		str.WriteString(params[k])
		str.WriteString("&")
	}
	return handler, str.String(), appidHandler, keys
}

// 按字母顺序遍历map
func TraverseMapOrder(params map[string]string) []string {

	/*	asciilist := make([]int,0)
		for k, _ := range params {
			asciilist = append(asciilist, gconv.Int(k))
		}
		sort.Ints(asciilist)
		keys := make([]string, 0)
		for _, k := range asciilist {
			keys = append(keys, string(k))
		}*/
	keys := make([]string, 0)
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Web 参数替换 undefined null 空
func ReplaceWebEmptyParams(s string) string {
	if s == "undefined" || s == "null" {
		return ""
	}
	return s
}

// 字符串切片去重
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}
