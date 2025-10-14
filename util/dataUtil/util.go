package dataUtil

import (
	"bytes"
	"reflect"
	"strings"

	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/util/excelutil"

	"github.com/gogf/gf/frame/g"
	"github.com/luoliDark/base/util/commutil"
	"github.com/xormplus/xorm"
)

const (
	ParameterEqual        = "="
	ParameterNotEqual     = "!="
	ParameterLike         = "like"
	ParameterGreater      = ">"
	ParameterLess         = "<"
	ParameterGreaterEqual = ">="
	ParameterLessEqual    = "<="
	ParameterIn           = "in"
	ParameterNotIn        = "not in"
	ParameterOr           = "or"
)

type ParameterInfo struct {
	field    string
	para     interface{}
	condType string
	plist    []ParameterInfo
}

func NewParameterInfo(fieldName, condType, para string) ParameterInfo {
	return ParameterInfo{
		field:    fieldName,
		para:     para,
		condType: condType,
	}
}

type DBExecutor struct {
	list             []ParameterInfo
	engine           *xorm.Engine
	sqlParameterList []interface{}
	sqlStr           string
	table            string
}

func NewNewDBExecutor(engine *xorm.Engine, table string) *DBExecutor {
	return &DBExecutor{
		engine: engine,
		table:  table,
	}
}

func (ph *DBExecutor) AddOrParameterList(l []ParameterInfo) {
	var pl []ParameterInfo
	for _, v := range l {
		if v.field == "" || v.para == "" || v.condType == "" {
			return
		}
		pi := ParameterInfo{
			field:    v.field,
			para:     v.para,
			condType: v.condType,
		}
		pl = append(pl, pi)
	}
	pp := ParameterInfo{}
	pp.condType = ParameterOr
	pp.plist = pl

	ph.list = append(ph.list, pp)
}

func (ph *DBExecutor) Add(fieldName, condType string, para interface{}) {
	if fieldName == "" || condType == "" || para == nil {
		return
	}
	pi := ParameterInfo{
		field:    fieldName,
		para:     para,
		condType: condType,
	}

	ph.list = append(ph.list, pi)
}

func (ph *DBExecutor) AddStr(fieldName, condType, para string) {
	if fieldName == "" || condType == "" || para == "" {
		return
	}
	pi := ParameterInfo{
		field:    fieldName,
		para:     para,
		condType: condType,
	}
	if condType == ParameterLike {
		pi.para = "%" + para + "%"
	}

	ph.list = append(ph.list, pi)
}

func (ph *DBExecutor) AddInt(fieldName, condType string, para, min int) {
	if fieldName == "" || condType == "" || min > para {
		return
	}

	ph.list = append(ph.list, ParameterInfo{
		field:    fieldName,
		para:     para,
		condType: condType,
	})
}

func (ph *DBExecutor) AddStrList(fieldName, condType string, para []string) {
	if fieldName == "" || condType == "" || g.IsEmpty(para) {
		return
	}
	pi := ParameterInfo{
		field:    fieldName,
		para:     para,
		condType: condType,
	}

	ph.list = append(ph.list, pi)
}

func (ph *DBExecutor) AddFloat64List(fieldName, condType string, para []float64) {
	if fieldName == "" || condType == "" || g.IsEmpty(para) {
		return
	}
	pi := ParameterInfo{
		field:    fieldName,
		para:     para,
		condType: condType,
	}

	ph.list = append(ph.list, pi)
}

func (ph *DBExecutor) AddIntList(fieldName, condType string, para []int) {
	if fieldName == "" || condType == "" || g.IsEmpty(para) {
		return
	}

	ph.list = append(ph.list, ParameterInfo{
		field:    fieldName,
		para:     para,
		condType: condType,
	})
}

func (ph *DBExecutor) SetParameters(parameters []ParameterInfo) {
	ph.list = parameters
}

func (ph *DBExecutor) doSqlEngine(sql string, session *xorm.Engine, id int) string {
	var str string
	if !strings.Contains(sql, "where") {
		str = sql + " where 1=1 "
		if id != 0 {
			str = str + " and b.configid = " + commutil.ToString(id)
		}
	} else {
		str = sql + " "
	}
	var pl []interface{}
	var ok bool
	for _, v := range ph.list {
		switch v.condType {
		case ParameterIn:

			t := reflect.TypeOf(v.para)
			if t == nil {
				return "nil"
			}
			switch t.Kind() {
			case reflect.Slice:
				// 进一步判断切片元素类型
				elemKind := t.Elem().Kind()
				switch elemKind {
				case reflect.String:
					var sl []string
					sl, ok := v.para.([]string)
					if ok {
						str = str + " and " + v.field + ` in ('` + strings.Join(sl, "','") + `')`
					}
				case reflect.Int:
					var il []int
					var sl []string
					il, ok = v.para.([]int)
					for _, vvv := range il {
						sl = append(sl, commutil.ToString(vvv))
					}
					if ok {
						str = str + " and " + v.field + ` in (` + strings.Join(sl, "','") + `)`
					}
				case reflect.Float64:
					var il []float64
					var sl []string
					il, ok = v.para.([]float64)
					for _, vvv := range il {
						sl = append(sl, commutil.ToString(vvv))
					}
					if ok {
						str = str + " and " + v.field + ` in (` + strings.Join(sl, "','") + `)`
					}
				}
			}
		case ParameterNotIn:
			//resultSession.NotIn(v.field, v.para)
		default:
			str = str + " and " + v.field + " " + v.condType + " ?"
			pl = append(pl, v.para)
		}
	}
	ph.sqlParameterList = pl
	ph.sqlStr = str

	return str
}

func (ph *DBExecutor) sqlEngine(session *xorm.Engine) *xorm.Session {
	str := ""
	var pl []interface{}
	resultSession := session.Table(ph.table).Where("1 = 1")
	for _, v := range ph.list {
		switch v.condType {
		case ParameterIn:
			resultSession.In(v.field, v.para)
		case ParameterNotIn:
			resultSession.NotIn(v.field, v.para)
		case ParameterOr:
			var (
				sp []interface{}
				ss []string
			)
			for _, vv := range v.plist {
				ss = append(ss, vv.field+" "+vv.condType+" ?")
				sp = append(sp, vv.para)
			}
			resultSession.And(strings.Join(ss, " or "), sp...)
		default:
			resultSession.And(v.field+" "+v.condType+" ?", v.para)
		}
	}

	return resultSession.Table(ph.table).Where(str, pl...)
}

func (ph *DBExecutor) AllInfoObject(session *xorm.Engine, tableName, orderBy string, data interface{}) error {
	s := ph.sqlEngine(session)
	if orderBy != "" {
		s.OrderBy(orderBy)
	}
	err := s.Find(data)

	return err
}

func (ph *DBExecutor) PageInfoBySqlWithGroupBy(session *xorm.Engine, sql, groupBy, orderBy string, pageIndex, pageSize int, data interface{}) sysmodel.ResultBean {
	str := ph.doSqlEngine(sql, session, 0)
	str += groupBy
	countStr := "select count(1) as cnt from (" + str + ") as tt where 1=1 "
	var results []map[string]int
	err := session.SQL(countStr, ph.sqlParameterList...).Find(&results)
	num := 0
	result := sysmodel.ResultBean{}
	result.ResultTotal = 0
	result.IsSuccess = true
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	if len(results) <= 0 {
		result.ResultTotal = 0
		result.ResultData = nil
		return result
	}
	num, ok := results[0]["cnt"]
	if !ok {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}

	if num <= 0 {
		result.ResultTotal = 0
		result.ResultData = nil
		return result
	}
	result.ResultTotal = commutil.ToInt(num)
	err = session.SQL(str, ph.sqlParameterList...).OrderBy(orderBy).Limit(pageSize, pageSize*(pageIndex-1)).Find(data)
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.ResultData = data
	return result
}

// 根据sql 获取总条数
func GetCountBySql(session *xorm.Engine, sql string) (int, error) {
	countStr := "select count(1) as cnt from (" + sql + ") as tt where 1=1 "
	var results []map[string]int
	err := session.SQL(countStr).Find(&results)
	if err != nil {
		return 0, err
	}
	num := 0
	num, ok := results[0]["cnt"]
	if !ok {
		return 0, nil
	}

	return num, nil
}

func (ph *DBExecutor) PageInfoBySqlWithGroupByWithExportMap(session *xorm.Engine, isExport int, sql, groupBy, orderBy string, pageIndex, pageSize int, exportInfo *excelutil.ExportInfo, id int) sysmodel.ResultBean {
	str := ph.doSqlEngine(sql, session, id)
	str += " " + groupBy
	result := sysmodel.ResultBean{}
	if isExport == 1 {
		m, err := session.SQL(str, ph.sqlParameterList...).OrderBy(orderBy).QueryString()
		if err != nil {
			result.IsSuccess = false
			result.ErrorMsg = err.Error()
			return result
		}
		err = excelutil.ExportExcelFromMapStringByTitleMap(exportInfo.SheetName, exportInfo.TitleList, m, exportInfo.Path)
		if err != nil {
			result.IsSuccess = false
			result.ErrorMsg = err.Error()
			return result
		}
		result.SetSuccess(exportInfo.WebPath)
		return result
	}
	countStr := "select count(1) as cnt from (" + str + ") as tt     "
	var results []map[string]int
	err := session.SQL(countStr, ph.sqlParameterList...).Find(&results)
	num := 0

	result.ResultTotal = 0
	result.IsSuccess = true
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	if len(results) <= 0 {
		result.ResultTotal = 0
		result.ResultData = nil
		return result
	}
	num, ok := results[0]["cnt"]
	if !ok {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}

	if num <= 0 {
		result.ResultTotal = 0
		result.ResultData = nil
		return result
	}
	result.ResultTotal = commutil.ToInt(num)
	if pageIndex == 1 {
	}
	ph.sqlParameterList = append(ph.sqlParameterList, (pageIndex-1)*pageSize)
	ph.sqlParameterList = append(ph.sqlParameterList, pageSize)
	mapData, err := session.SQL(str+" order by "+orderBy+" LIMIT  ? ,?", ph.sqlParameterList...).QueryString()
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.ResultData = mapData
	return result
}

func (ph *DBExecutor) PageInfoBySqlWithGroupByWithExport(session *xorm.Engine, isExport int, sql, groupBy, orderBy string, pageIndex, pageSize int, data interface{}, exportInfo *excelutil.ExportInfo) sysmodel.ResultBean {
	str := ph.doSqlEngine(sql, session, 0)
	str += groupBy
	result := sysmodel.ResultBean{}
	if isExport == 1 {
		m, err := session.SQL(str, ph.sqlParameterList...).OrderBy(orderBy).QueryString()
		if err != nil {
			result.IsSuccess = false
			result.ErrorMsg = err.Error()
			return result
		}
		err = excelutil.ExportExcelFromMapStringByTitleMap(exportInfo.SheetName, exportInfo.TitleList, m, exportInfo.Path)
		if err != nil {
			result.IsSuccess = false
			result.ErrorMsg = err.Error()
			return result
		}
		result.SetSuccess(exportInfo.WebPath)
		return result
	}
	countStr := "select count(1) as cnt from (" + str + ") as tt where 1=1 "
	var results []map[string]int
	err := session.SQL(countStr, ph.sqlParameterList...).Find(&results)
	num := 0

	result.ResultTotal = 0
	result.IsSuccess = true
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	if len(results) <= 0 {
		result.ResultTotal = 0
		result.ResultData = nil
		return result
	}
	num, ok := results[0]["cnt"]
	if !ok {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}

	if num <= 0 {
		result.ResultTotal = 0
		result.ResultData = nil
		return result
	}
	result.ResultTotal = commutil.ToInt(num)
	err = session.SQL(str, ph.sqlParameterList...).OrderBy(orderBy).Limit(pageSize, pageSize*(pageIndex-1)).Find(data)
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.ResultData = data
	return result
}

func (ph *DBExecutor) PageInfoBySql(session *xorm.Engine, sql, orderBy string, pageIndex, pageSize int, data interface{}) sysmodel.ResultBean {
	str := ph.doSqlEngine(sql, session, 0)
	countStr := "select count(1) as cnt from (" + str + ") as tt where 1=1 "
	var results []map[string]int
	err := session.SQL(countStr, ph.sqlParameterList...).Find(&results)
	num := 0
	result := sysmodel.ResultBean{}
	result.ResultTotal = 0
	result.IsSuccess = true
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	if len(results) <= 0 {
		result.ResultTotal = 0
		result.ResultData = nil
		return result
	}
	num, ok := results[0]["cnt"]
	if !ok {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}

	if num <= 0 {
		result.ResultTotal = 0
		result.ResultData = nil
		return result
	}
	result.ResultTotal = commutil.ToInt(num)
	err = session.SQL(str, ph.sqlParameterList...).OrderBy(orderBy).Limit(pageSize, pageSize*(pageIndex-1)).Find(data)
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.ResultData = data
	return result
}

func (ph *DBExecutor) PageInfoObjectWithExport(session *xorm.Engine, isExport int, tableName, orderBy string, pageIndex, pageSize int, data interface{}, exportInfo *excelutil.ExportInfo) sysmodel.ResultBean {
	result := sysmodel.ResultBean{}
	ss := ph.sqlEngine(session)
	if isExport == 1 {
		m, err := ph.sqlEngine(session).OrderBy(orderBy).QueryString()
		if err != nil {
			result.IsSuccess = false
			result.ErrorMsg = err.Error()
			return result
		}
		err = excelutil.ExportExcelFromMapStringByTitleMap(exportInfo.SheetName, exportInfo.TitleList, m, exportInfo.Path)
		if err != nil {
			result.IsSuccess = false
			result.ErrorMsg = err.Error()
			return result
		}
		result.SetSuccess(exportInfo.WebPath)
		return result
	}
	num, err := ss.Count()
	result.ResultTotal = 0
	result.IsSuccess = true
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	if num <= 0 {
		return result
	}
	result.ResultTotal = commutil.ToInt(num)
	err = ph.sqlEngine(session).OrderBy(orderBy).Limit(pageSize, pageSize*(pageIndex-1)).Find(data)
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.ResultData = data
	return result
}

func (ph *DBExecutor) PageInfoObjectV2(session *xorm.Engine, tableName, orderBy string, pageIndex, pageSize, isAll int, data interface{}) sysmodel.ResultBean {
	ss := ph.sqlEngine(session)
	result := sysmodel.ResultBean{}
	result.ResultTotal = 0
	var err error
	result.IsSuccess = true
	if isAll == 1 {
		err = ph.sqlEngine(session).OrderBy(orderBy).Find(data)
	} else {
		num, err := ss.Count()
		if err != nil {
			result.IsSuccess = false
			result.ErrorMsg = err.Error()
			return result
		}
		result.ResultTotal = commutil.ToInt(num)
		if num <= 0 {
			return result
		}
		err = ph.sqlEngine(session).OrderBy(orderBy).Limit(pageSize, pageSize*(pageIndex-1)).Find(data)
	}
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.ResultData = data
	return result
}

func (ph *DBExecutor) PageInfoObject(session *xorm.Engine, tableName, orderBy string, pageIndex, pageSize int, data interface{}) sysmodel.ResultBean {
	ss := ph.sqlEngine(session)
	num, err := ss.Count()
	result := sysmodel.ResultBean{}
	result.ResultTotal = 0
	result.IsSuccess = true
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}

	if num <= 0 {
		return result
	}
	result.ResultTotal = commutil.ToInt(num)
	err = ph.sqlEngine(session).OrderBy(orderBy).Limit(pageSize, pageSize*(pageIndex-1)).Find(data)
	if err != nil {
		result.IsSuccess = false
		result.ErrorMsg = err.Error()
		return result
	}
	result.ResultData = data
	return result
}

func (ph *DBExecutor) PageInfo(session *xorm.Engine, tableName, orderBy string, pageIndex, pageSize int, data interface{}) (interface{}, int64, error) {
	ss := ph.sqlEngine(session)
	num, err := ss.Count()
	if err != nil {
		return data, 0, err
	}
	result := &sysmodel.ResultBean{}

	if num <= 0 {
		return data, 0, nil
	}
	result.ResultTotal = commutil.ToInt(num)
	err = ph.sqlEngine(session).OrderBy(orderBy).Limit(pageSize, pageSize*(pageIndex-1)).Find(data)
	if err != nil {
		return data, 0, err
	}
	return data, num, nil
}

func (ph *DBExecutor) GetList(session *xorm.Engine, tableName, orderBy string, data interface{}) (interface{}, error) {
	err := ph.sqlEngine(session).OrderBy(orderBy).Find(data)
	return data, err
}

func (ph *DBExecutor) Sql(session *xorm.Session) *xorm.Session {
	str := ""
	var pl []interface{}
	for _, v := range ph.list {
		if str != "" {
			str += " and "

		}
		str += v.field + " " + v.condType + " ?"

		pl = append(pl, v.para)
	}

	return session.Where(str, pl...)
}

func ExecDeleteByList(session *xorm.Session, tableName, fieldName string, valueList []string) error {
	sql := new(bytes.Buffer)
	sql.WriteString("delete from " + tableName + " where " + fieldName + " in (")
	for k, v := range valueList {
		sql.WriteString("'" + v + "'")
		if k != len(valueList)-1 {
			sql.WriteString(",")
		}
	}
	sql.WriteString(")")

	_, err := session.Exec(sql.String())
	if err != nil {
		return err
	}

	return nil
}
