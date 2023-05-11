package excelutil

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/luoliDark/base/util/commutil"
	"github.com/tealeg/xlsx"
)

var ostype = runtime.GOOS

// 最小工作表名称长度为 1 个字符。
// 最大工作表名称长度为 31 个字符。
// 这些特殊字符也不允许： / ？* [ ]
func AddSheetByMap(excel *xlsx.File, sheetName string, mapDatas []map[string]interface{}, headers []string) error {
	sh, err := excel.AddSheet(sheetName)
	if err != nil {
		return err
	}
	if len(mapDatas) == 0 {
		return nil
	}
	titles := headers
	if len(titles) == 0 {
		titles = make([]string, 0)
		row := sh.AddRow()
		for key := range mapDatas[0] {
			cell := row.AddCell()
			cell.Value = key
			titles = append(titles, key)
		}
	} else {
		row := sh.AddRow()
		for _, header := range titles {
			cell := row.AddCell()
			cell.Value = header
		}
	}

	for _, itemRowData := range mapDatas {
		row := sh.AddRow()
		for _, title := range titles {
			cell := row.AddCell()
			cell.Value = commutil.ToString(itemRowData[title])
		}
	}
	return nil
}

func AddSheetByTitleMap(excel *xlsx.File, sheetName string, mapDatas []map[string]interface{}, headers []TtileByKey) error {
	sh, err := excel.AddSheet(sheetName)
	if err != nil {
		return err
	}
	if len(mapDatas) == 0 {
		return nil
	}
	titles := headers
	if len(headers) == 0 {
		row := sh.AddRow()
		for key := range mapDatas[0] {
			cell := row.AddCell()
			cell.Value = key
			titles = append(titles, TtileByKey{Title: key, Key: key})
		}
	} else {
		row := sh.AddRow()
		for _, title := range headers {
			cell := row.AddCell()
			cell.Value = title.Title
		}
	}

	for _, itemRowData := range mapDatas {
		row := sh.AddRow()
		for _, title := range titles {
			cell := row.AddCell()
			cell.Value = commutil.ToString(itemRowData[title.Key])
		}
	}
	return nil
}

// map导出excel headers {字段名：标题名}
func ExportExcelFromMapStringByTitleMap(sheetName string, headers []TtileByKey, mapDatas []map[string]string, outPath string) error {
	excel := xlsx.NewFile()
	err := AddSheetByTitleMapString(excel, sheetName, mapDatas, headers)
	if err != nil {
		return err
	}
	err = excel.Save(formatFilePath(outPath))
	if err != nil {
		return err
	}
	return nil
}

func AddSheetByTitleMapString(excel *xlsx.File, sheetName string, mapDatas []map[string]string, headers []TtileByKey) error {
	sh, err := excel.AddSheet(sheetName)
	if err != nil {
		return err
	}
	if len(mapDatas) == 0 {
		return nil
	}
	titles := headers
	if len(headers) == 0 {
		row := sh.AddRow()
		for key := range mapDatas[0] {
			cell := row.AddCell()
			cell.Value = key
			titles = append(titles, TtileByKey{Title: key, Key: key})
		}
	} else {
		row := sh.AddRow()
		for _, title := range headers {
			cell := row.AddCell()
			cell.Value = title.Title
		}
	}

	for _, itemRowData := range mapDatas {
		row := sh.AddRow()
		for _, title := range titles {
			cell := row.AddCell()
			cell.Value = itemRowData[title.Key]
		}
	}
	return nil
}

func AddSheetByTitlePointerMap(excel *xlsx.File, sheetName string, mapDatas []map[string]*string, headers []TtileByKey) error {
	sh, err := excel.AddSheet(sheetName)
	if err != nil {
		return err
	}
	if len(mapDatas) == 0 {
		return nil
	}
	titles := headers
	if len(headers) == 0 {
		row := sh.AddRow()
		for key := range mapDatas[0] {
			cell := row.AddCell()
			cell.Value = key
			titles = append(titles, TtileByKey{Title: key, Key: key})
		}
	} else {
		row := sh.AddRow()
		for _, title := range headers {
			cell := row.AddCell()
			cell.Value = title.Title
		}
	}

	for _, itemRowData := range mapDatas {
		row := sh.AddRow()
		for _, title := range titles {
			cell := row.AddCell()
			if value := itemRowData[title.Key]; value == nil {
				cell.Value = ""
			} else {
				cell.Value = *value
			}
		}
	}
	return nil
}

// map导出excel
func ExportExcelFromMap(sheetName string, headers []string, mapDatas []map[string]interface{}, outPath string) error {
	excel := xlsx.NewFile()
	err := AddSheetByMap(excel, sheetName, mapDatas, headers)
	if err != nil {
		return err
	}
	err = excel.Save(formatFilePath(outPath))
	if err != nil {
		return err
	}
	return nil
}

type TtileByKey struct {
	Key         string //字段编码
	Title       string //标题
	IsNumber    bool
	IsRrequired bool
}

// map导出excel headers {字段名：标题名}
func ExportExcelFromMapByTitleMap(sheetName string, headers []TtileByKey, mapDatas []map[string]interface{}, outPath string) error {
	excel := xlsx.NewFile()
	err := AddSheetByTitleMap(excel, sheetName, mapDatas, headers)
	if err != nil {
		return err
	}
	err = excel.Save(formatFilePath(outPath))
	if err != nil {
		return err
	}
	return nil
}

// map导出excel headers {字段名：标题名}
func ExportExcelFromPointerMapByTitleMap(sheetName string, headers []TtileByKey, mapDatas []map[string]*string, outPath string) error {
	excel := xlsx.NewFile()
	err := AddSheetByTitlePointerMap(excel, sheetName, mapDatas, headers)
	if err != nil {
		return err
	}
	err = excel.Save(formatFilePath(outPath))
	if err != nil {
		return err
	}
	return nil
}

func formatFilePath(pathStr string) string {
	if ostype == "windows" {
		pathStr = strings.ReplaceAll(pathStr, "/", "\\")
	}
	pathStr = strings.ReplaceAll(pathStr, "\\", "/")
	dir := path.Dir(pathStr)
	_, errDir := os.Stat(dir)
	if errDir != nil {
		os.MkdirAll(dir, os.ModePerm)
	}
	return pathStr
}

// struct根据tag导出excel
func ExportExcelFromStruct(sheetName string, records interface{}, outPath string) error {
	excel := xlsx.NewFile()
	err := AddSheetByStruct(excel, sheetName, records)
	if err != nil {
		return err
	}
	err = excel.Save(formatFilePath(outPath))
	if err != nil {
		return err
	}
	return nil
}

func AddSheetByStruct(excel *xlsx.File, sheetName string, records interface{}) error {
	sh, _ := excel.AddSheet(sheetName) // new sheet
	t := reflect.TypeOf(records)
	if t.Kind() != reflect.Slice {
		return fmt.Errorf("对象类型为 %v ，不是Slice无法转换！", t.Kind())
	}
	s := reflect.ValueOf(records)
	firstrow := sh.AddRow()

	for i := 0; i < s.Len(); i++ {
		elem := s.Index(i).Interface()
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		row := sh.AddRow()

		for j := 0; j < elemType.NumField(); j++ {
			field := elemType.Field(j)
			tag := field.Tag.Get("xlsx")
			name := tag
			if tag == "" {
				continue
			}
			// 设置表头
			if i == 0 {
				cell := firstrow.AddCell()
				cell.Value = commutil.ToString(name)
			}
			cell := row.AddCell()
			fieldtype := field.Type.Kind()
			if isIntTypeByField(fieldtype) {
				cell.SetInt(commutil.ToInt(elemValue.Field(j).Interface()))
			} else if isFloatTypeByField(fieldtype) {
				cell.SetFloat(commutil.ToFloat64(elemValue.Field(j).Interface()))
			} else {
				cell.Value = commutil.ToString(elemValue.Field(j).Interface())
			}

		}
	}
	return nil
}

func isIntTypeByField(fieldtype reflect.Kind) bool {
	if fieldtype == reflect.Int || fieldtype == reflect.Int32 ||
		fieldtype == reflect.Int64 || fieldtype == reflect.Int8 {
		return true
	}
	return false
}

func isFloatTypeByField(fieldtype reflect.Kind) bool {
	if fieldtype == reflect.Float64 || fieldtype == reflect.Float32 {
		return true
	}
	return false
}
