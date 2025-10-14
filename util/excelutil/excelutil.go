package excelutil

import (
	"container/list"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/redishelper/rediscache"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/util/commutil"

	"github.com/tealeg/xlsx"
	"github.com/xormplus/xorm"
)

type ExportInfo struct {
	TitleList []TtileByKey
	Name      string
	SheetName string
	Path      string
	WebPath   string
}

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

//替换回车符号
func ReplaceRNT(cellValue string) string {
	cellValue = strings.Replace(cellValue, "\n", " ", -1)
	cellValue = strings.Replace(cellValue, "\r", " ", -1)
	cellValue = strings.Replace(cellValue, "\t", " ", -1)
	return cellValue
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

		cellValue = convertToFormatDay(cell.Value)
	} else if cell.Type() == xlsx.CellTypeString {
		cellValue = cell.String()
	} else if cell.Type() == xlsx.CellTypeNumeric {
		//GeneralNumeric Excel格式是数字类型，但是数据库字段类型是文本。正常取原值不会有问题，
		//但此方法会把数字转成科学计数法的文本，导致异常 改成使用 GeneralNumericWithoutScientific
		cellValue, _ = cell.GeneralNumericWithoutScientific()
	} else if cell.Type() == xlsx.CellTypeBool {
		cellValue = commutil.ToString(cell.Bool())
	} else if cell.Type() == xlsx.CellTypeError {
		cellValue = ""
	} else if cell.Type() == xlsx.CellTypeStringFormula {
		cellValue = cell.Formula()
	} else if cell.Type() == xlsx.CellTypeInline {
		cellValue = cell.Value
	}

	//替换回车符号
	cellValue = strings.Replace(cellValue, "\n", " ", -1)
	cellValue = strings.Replace(cellValue, "\r", " ", -1)

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
func convertToFormatDay(excelDaysString string) string {
	// 2006-01-02 距离 1900-01-01的天数
	baseDiffDay := 38719 //在网上工具计算的天数需要加2天，什么原因没弄清楚
	curDiffDay := excelDaysString
	b, err := strconv.Atoi(curDiffDay)
	if err != nil {
		fmt.Println("时间转换失败")
		return excelDaysString
	}
	// 获取excel的日期距离2006-01-02的天数
	realDiffDay := b - baseDiffDay
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
			if strings.Index(dataSource, "9999") == 0 {
				//如果是字典数据源，则只能按名称
				isConvertByName = true
				isConvertByCode = false
				entity.SetIsPidSource(false)
			} else {
				entity.SetIsPidSource(true)
			}
			if !isConvertByName && !isConvertByCode {
				isConvertByName = true // isConvertByCode 原先是编码默认，现在改为名称默认
			}
		}

		var titleName = name
		if isConvertByCode {
			var bianma = rediscache.GetLanguageText("label", langCode, "DsCode")
			if xorm.IsNumeric(bianma) {
				if langCode == "zh" {
					bianma = "编码"
				} else {
					bianma = "code"
				}
			}
			titleName = titleName + bianma
		}

		entity.SetShowname(name)
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

/*
*
读取excel并且返回数据列
*/
func ReadExcel(excelPath string) [][]string {
	xlsx, err := xlsx.OpenFile(excelPath)
	if err != nil {
		fmt.Println("open excel error,", err.Error())
		os.Exit(1)
	}
	rows := xlsx.Sheets[0].Rows
	result := make([][]string, 0)
	for _, row := range rows {
		var cellIndex = len(row.Cells)
		var valuemap = make([]string, 0)
		for i := 0; i < cellIndex; i++ {
			//获取值
			value := GetStringValueFromCell(row.Cells[i])
			valuemap = append(valuemap, value)
		}
		if nil != valuemap && len(valuemap) > 0 {
			result = append(result, valuemap)
		}
	}
	return result
}

func CreateExcelPath() (string, string) {
	timeStr := time.Now().Format("20060102150405")
	path := confighelper.LoadGoEnv()
	path += "temp/excel/" + timeStr + ".xls"
	webPath := "/base/static/temp/excel/" + timeStr + ".xls"
	return path, webPath
}
