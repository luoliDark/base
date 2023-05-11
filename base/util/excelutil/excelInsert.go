package excelutil

import (
	"base/base/constant"
	"base/base/util/commutil"
	"fmt"
	"regexp"
	"runingproject/services/excel/export"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gogf/gf/frame/g"
)

var ExcelRegexpNumberType = regexp.MustCompile(`\b(int|tinyint|smallint|mediumint|bigint|float|double|decimal|numeric)\b`)

// BatchInsertListDataByExcelize
// 引用包：github.com/360EntSecGroup-Skylar/excelize
// 批量插入数据到excel,
func BatchInsertListDataByExcelize(f *excelize.File, SheetName string, rows []map[string]string, titles []map[string]string) (err error) {
	if len(titles) == 0 {
		return nil
	}
	//因为是多协程写入，同时导出需要限制人数。 避免cpu占用， 造成服务无法访问。
	//防止前面检查失效，写入前再做检查，是否已经有人在导入大数据量写入excel。
	if export.IsBigDataExport(len(rows)) && export.IsAllowBigDataExport() {
		return fmt.Errorf("已有用户在导出数据，本次单据导出数据量大，服务器忙不过来了，请稍后五分钟再试！")
	}

	//中文SheetName名称 好像不能插入？ 待排查
	if SheetName == "" {
		SheetName = "Sheet1"
	}
	//中文标题
	var ChTitleList []interface{}
	//英文标题
	var EnTitleList []string
	var EnTitleMapByNumber = make(map[string]bool, 0) //数字格式的字段
	for _, colM := range titles {
		sqlcol := colM["sqlcol"]
		name := colM["name"]
		datasouce := colM["datasource"]
		ChTitleList = append(ChTitleList, name)
		if !g.IsEmpty(datasouce) {
			EnTitleList = append(EnTitleList, sqlcol+"_show")
		} else {
			EnTitleList = append(EnTitleList, sqlcol)
			sqldatatype := strings.Trim(colM["sqldatatype"], " ")
			if ExcelRegexpNumberType.Match([]byte(sqldatatype)) {
				EnTitleMapByNumber[sqlcol] = true
			}
		}
	}

	rowIndex := 1
	//插入标题
	f.SetSheetRow(SheetName, commutil.AppendStr("A", rowIndex), &ChTitleList)
	// 注： excel插件内存中也是用数组存放数据， 需要先在最后一行申请一行空的数据。 这样excel数组就会扩容到导出大小 。
	// 1、提交扩容可避免插入行时 插入时间和内存消耗剧增。
	// 2、使用协程异步执行会导致 数组越界。因为插入时的行号小于 excel数组大小 所以需要提前扩容
	rowsLen := len(rows)
	v2 := make([]interface{}, len(EnTitleList))
	f.SetSheetRow(SheetName, commutil.AppendStr("A", rowsLen+1), &v2)

	isNumber := false
	var errNumber error
	var toFloatVal float64
	//循环插入内容
	EnTitleListLen := len(EnTitleList)
	var rowsvalueListVSList = make([][]interface{}, 0)
	var rowsvalueList []interface{}
	wg := sync.WaitGroup{}
	// 取十分之四+ 1 的协程数  测试：16逻辑cpu，execCpuNum = 5 时， 导出6W条数据，两个人12W cpu占比 70% 、
	execCpuNum := constant.CPUNumber/10*4 + 1
	// 分次插入，得出每次插入数量 ，不能超过cpu数量， 避免协程过多造成服务cpu暴涨。
	oneInsertLen := rowsLen / execCpuNum
	execCpuNum, oneInsertLen = updateInseretParams(execCpuNum, oneInsertLen, rowsLen)
	for i := 0; i < rowsLen; i++ {
		rowsvalueList = make([]interface{}, 0, EnTitleListLen)
		//循环对比标题
		for _, enName := range EnTitleList {
			dataValue := rows[i][enName]
			isNumber = false
			// 如何表单配置对应的是number类型的就把数据转成float
			if EnTitleMapByNumber[enName] {
				toFloatVal, errNumber = strconv.ParseFloat(dataValue, 64)
				isNumber = errNumber == nil
			}
			if isNumber {
				rowsvalueList = append(rowsvalueList, toFloatVal)
			} else {
				rowsvalueList = append(rowsvalueList, dataValue)
			}
		}
		rowsvalueListVSList = append(rowsvalueListVSList, rowsvalueList)
		if (i > 0 && i%oneInsertLen == 0) || i == rowsLen-1 {
			wg.Add(1)
			go func(rowIndex int, rowsvalueListVSList [][]interface{}) {
				defer func() {
					wg.Done()
					r := recover()
					if err == nil && r != nil {
						err = fmt.Errorf("并发写入excel错误：%v,%s", r, debug.Stack())
						fmt.Println(err)
					}
				}()
				for _, v := range rowsvalueListVSList {
					rowIndex++
					t := time.Now()
					if rowIndex%2000 == 0 {
						fmt.Println(rowIndex, t)
					}
					//第一行被字段占用，从第二行开始写入整行数据
					//必须使用函数传入，否则协程取到的变量可能不是当前使用的，
					f.SetSheetRow("Sheet1", commutil.AppendStr("A", rowIndex), &v)
				}
			}(rowIndex, rowsvalueListVSList)
			//下一次开始的行号
			rowIndex = rowIndex + len(rowsvalueListVSList)
			rowsvalueListVSList = make([][]interface{}, 0, oneInsertLen)
		}
	}
	wg.Wait()

	return
}

//修正
func updateInseretParams(execCpuNum int, oneInsertLen int, rowsLen int) (int, int) {
	if execCpuNum == 1 {
		execCpuNum = 2
	}
	if oneInsertLen == 0 {
		oneInsertLen = rowsLen
	}
	return execCpuNum, oneInsertLen
}
