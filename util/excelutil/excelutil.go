package excelutil

import (
	"container/list"
	"github.com/tealeg/xlsx"
	"paas/base/sysmodel"
	"paas/base/util/commutil"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetCellString(cell *xlsx.Cell) (cellString string) {

	if reflect.ValueOf(cell).IsNil() {
		return cellString
	}
	cellType := cell.Type()
	switch cellType {
	case xlsx.CellTypeBool:
		cellString = commutil.ToString(cell.Bool())
		break
	case xlsx.CellTypeError:
		panic("获取单元格信息错误")
		break
	case xlsx.CellTypeStringFormula:
		cellString = cell.Formula()
		break
	case xlsx.CellTypeNumeric:
		cellString, _ = cell.GeneralNumeric()
		break
	case xlsx.CellTypeString:
		cellString = cell.String()
		break
	case xlsx.CellTypeInline:
		cellString = cell.Value
		break
	default:
		break

	}
	return cellString
}

func RemoveZero(value string) string {
	if "" == value {
		return ""
	}
	var index = strings.LastIndex(value, ".")
	if index > 0 {
		var string = value[index+1:]
		var pattern = "^[0]*$"
		regexp, _ := regexp.MatchString(pattern, string)
		if regexp {
			return value[0:index]
		} else {
			return value
		}
	} else {
		return value
	}
	return value
}

func GetStringValueFromCell(cell *xlsx.Cell) string {
	/*SimpleDateFormat sFormat = new SimpleDateFormat("MM/dd/yyyy");
	DecimalFormat decimalFormat = new DecimalFormat("#.####");*/
	var cellValue = ""
	if reflect.ValueOf(cell).IsNil() {
		return cellValue
	} else if strings.Index(cell.NumFmt, "h:") >= 0 {
		cellValue = convertToFormatTime(cell.Value) //时分秒
	} else if strings.Index(cell.NumFmt, "yy") >= 0 {
		b, err := strconv.Atoi(cell.Value)
		if err != nil {
			//转成数字失败，不是日期类型（有些文本字段被设置为日期格式了）。不做处理。继续流转
		} else {
			return convertToFormatDay(b)
		}
	}

	if cell.Type() == xlsx.CellTypeString {
		cellValue = cell.String()
	} else if cell.Type() == xlsx.CellTypeNumeric {
		/*if xlsx.CellTypeDate {
			cellValue = cell.Value
		}else {

		}*/
		cellValue, _ = cell.GeneralNumeric()
	} else if cell.Type() == xlsx.CellTypeBool {
		cellValue = commutil.ToString(cell.Bool())
	} else if cell.Type() == xlsx.CellTypeError {
		return cellValue
	} else if cell.Type() == xlsx.CellTypeStringFormula {
		cellValue = cell.Formula()
	} else if cell.Type() == xlsx.CellTypeInline {
		cellValue = cell.Value
	}

	//替换回车符号
	cellValue = strings.Replace(cellValue, "\n", " ", -1)
	cellValue = strings.Replace(cellValue, "\r", " ", -1)

	cellValue = strings.Replace(cellValue, "\"", "“", -1)
	cellValue = strings.Replace(cellValue, "\\", "/", -1)
	cellValue = strings.Replace(cellValue, "'", "‘", -1)
	return cellValue
}

// excel日期字段格式化 yyyy-mm-dd  h:mm:ss  返回数据会加上毫秒 例：2023-04-23 23:59:58.999
func convertToFormatTime(excelDaysString string) string {

	var in float64 = commutil.ToFloat64(excelDaysString)

	excelEpoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)

	tm := excelEpoch.Add(time.Duration(in * float64(24*time.Hour)))

	return tm.Format(commutil.Time_Fomat01)
}

// excel日期字段格式化 yyyy-mm-dd
func convertToFormatDay(curDiffDay int) string {
	// 2006-01-02 距离 1900-01-01的天数
	baseDiffDay := 38719 //在网上工具计算的天数需要加2天，什么原因没弄清楚

	// 获取excel的日期距离2006-01-02的天数
	realDiffDay := curDiffDay - baseDiffDay
	//fmt.Println("realDiffDay:",realDiffDay)
	// 距离2006-01-02 秒数
	realDiffSecond := realDiffDay * 24 * 3600
	//fmt.Println("realDiffSecond:",realDiffSecond)
	// 2006-01-02 15:04:05距离1970-01-01 08:00:00的秒数 网上工具可查出
	baseOriginSecond := 1136185445
	resultTime := time.Unix(int64(baseOriginSecond+realDiffSecond), 0).Format("2006-01-02")
	return resultTime
}

func HasTitle(list []sysmodel.ExcelEntity, obj string, param string) int {
	var b = -1
	if len(list) <= 0 || "" == obj || "" == param {
		return b
	}
	for index, excelEntity := range list {
		var string = ""
		if param == "titlename" {
			string = excelEntity.GetTitle()
		} else if param == "sqlcol" {
			string = excelEntity.GetSqlCol()
		} else if param == "datasource" {
			string = excelEntity.GetDataSource()
		} else {
			string = excelEntity.GetTitle()
		}
		if string == obj {
			return index
		}
	}
	return b
}

func ChangeCellColor(file *xlsx.File, list *list.List) *xlsx.File {
	for value := list.Front(); nil != value; value = value.Next() {
		tempVal := value.Value.(string)
		sheetName := tempVal[0:strings.Index(tempVal, "-")]

		row := commutil.ToInt(tempVal[strings.Index(tempVal, "-")+1 : strings.LastIndex(tempVal, "-")])
		col := commutil.ToInt(tempVal[strings.LastIndex(tempVal, "-")+1:])

		sheet := file.Sheet[sheetName]
		row2 := sheet.Row(row)
		cell := row2.Cells[col]
		if commutil.IsNullOrEmpty(cell) {
			cell = row2.AddCell()
		}
		style := &xlsx.Style{}
		style.Fill.BgColor = "EFEFDE"
		style.Fill.PatternType = "EFEFDE"
		cell.SetStyle(style)

	}
	return file
}

func IsNullRow(row *xlsx.Row) (result bool) {

	if row == nil {
		return true
	}
	result = true
	for index := range row.Cells {
		cell := row.Cells[index]
		value := ""
		if cell != nil {
			switch cell.Type() {
			case xlsx.CellTypeInline:
				value = cell.Value
				break
			case xlsx.CellTypeString:
				value = cell.String()
				break
			case xlsx.CellTypeNumeric:
				value, _ = cell.GeneralNumeric()
				break
			case xlsx.CellTypeBool:
				value = commutil.ToString(cell.Bool())
				break
			case xlsx.CellTypeStringFormula:
				value = commutil.ToString(cell.Formula())
				break
			default:
				break
			}

			if "" != strings.ReplaceAll(value, " ", "") {
				result = false
				break
			}
		}
	}

	return result
}

func DataToExcelEntity(dataEntity []map[string]string, isMain bool, langCode string) (listobj []sysmodel.ExcelEntity) {
	if nil == dataEntity || len(dataEntity) <= 0 || commutil.IsNullOrEmpty(isMain) {
		return listobj
	}
	//当前字段是否已经到处编码列
	for i := 0; i < len(dataEntity); i++ {
		tempMap := dataEntity[i]
		var name = commutil.ToString(tempMap["showname"]) //name
		var sqlCol = strings.ToLower(commutil.ToString(tempMap["sqlcol"]))
		var sqlDataType = commutil.ToString(tempMap["sqldatatype"])
		var dataSource = commutil.ToString(tempMap["datasource"])
		var isConvertByCode = commutil.ToBool(tempMap["isconvertbycode"])
		var isConvertByName = commutil.ToBool(tempMap["isconvertbyname"])
		var isRequired = commutil.ToBool(tempMap["isrequired"])
		var isSingle = commutil.IsNullOrEmpty(tempMap["issingle"]) || commutil.ToBool(tempMap["issingle"])
		//初始化对象
		var entity = sysmodel.ExcelEntity{}
		///判断数据源是否为空
		if !commutil.IsNullOrEmpty(dataSource) {
			entity.SetIsPidSource(true)
			if !isConvertByName && !isConvertByCode {
				isConvertByName = true // isConvertByCode 原先是编码默认，现在改为名称默认
			}
		}

		var titleName = name

		entity.SetShowname(titleName)
		entity.SetIsConvertByCode(isConvertByCode)
		entity.SetIsConvertByName(isConvertByName)
		entity.SetDataSource(dataSource)
		entity.SetIsMain(isMain)
		entity.SetIsRequired(isRequired)
		entity.SetSqlCol(sqlCol)
		entity.SetSqlDataType(sqlDataType)
		entity.SetTitle(titleName)
		entity.SetIsSingle(isSingle)
		listobj = append(listobj, entity)

	}

	return listobj
}
