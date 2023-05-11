//jsz by 2020.2.2 用于所有SQL语句的执行，所有SQL必须调用本类方法进行执行
//后期如出现更好用的ORM框架时 可在此文件中替换即可，不影响业务系统调用

package dbhelper

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/xormplus/xorm"
	xormplus_core "github.com/xormplus/xorm/core"
	"paas/base/confighelper"
	"paas/base/db/conn"
	"paas/base/loghelper"
	"paas/base/sysmodel"
	"paas/base/sysmodel/logtype"
	"paas/base/util/commutil"
	"paas/base/util/jsonutil"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 拼接sql 执行失败脚本
func concatErr(userID string, err error, sqlOrArgs ...interface{}) string {
	var sql string
	var buffer bytes.Buffer

	for index, _ := range sqlOrArgs {
		if index == 0 {
			sql = commutil.ToString(sqlOrArgs[index])
			buffer.WriteString(sql)
		} else {

			buffer.WriteString(" 参数")
			buffer.WriteString(gconv.String(index))
			buffer.WriteString("=")
			buffer.WriteString(commutil.ToString(sqlOrArgs[index]))
		}
	}
	buffer.WriteString("错误信息为：")
	buffer.WriteString(err.Error())
	return buffer.String()
}

// QuerySql 执行SQL查询，可传递参数，并返回[]map[string]string
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
func Query(userID string, IsMasterDB bool, sqlOrArgs ...interface{}) ([]map[string]string, error) {

	// 通用方法无需区分数据库类型
	tim := loghelper.BeginimeRecord() //记录开始时间
	engine, _ := conn.GetConnection(userID, IsMasterDB)
	results, err := engine.QueryString(sqlOrArgs...)

	if err != nil {
		var sqlError = concatErr(userID, err, sqlOrArgs...)
		loghelper.ByError(logtype.QueryErr, sqlError, userID)
		return nil, err
	}

	loghelper.EndTimeRecord(userID, "执行SQL", tim, sqlOrArgs...) //记录结束时间
	return results, nil
}

//从指定session进行查询
// QuerySql 执行SQL查询，可传递参数，并返回[]map[string]string
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
func QueryByTran(session *xorm.Session, userID string, IsMasterDB bool, sqlOrArgs ...interface{}) ([]map[string]string, error) {

	// 通用方法无需区分数据库类型
	tim := loghelper.BeginimeRecord() //记录开始时间
	results, err := session.QueryString(sqlOrArgs...)
	if err != nil {
		var sqlError = concatErr(userID, err, sqlOrArgs...)
		loghelper.ByError(logtype.QueryErr, sqlError, userID)
		return nil, err
	}
	loghelper.EndTimeRecord(userID, "执行SQL", tim, sqlOrArgs...) //记录结束时间
	return results, nil
}

// QuerySql 执行SQL查询，可传递参数，并返回JSON
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
func QueryByJson(userID string, IsMasterDB bool, sqlOrArgs ...interface{}) (string, error) {
	tim := loghelper.BeginimeRecord() //记录开始时间
	engine, _ := conn.GetConnection(userID, IsMasterDB)
	result, err := engine.QueryString(sqlOrArgs...)
	if err != nil {

		var sqlError = concatErr(userID, err, sqlOrArgs...)
		loghelper.ByError(logtype.QueryErr, sqlError, userID)
		return "", err
	}

	loghelper.EndTimeRecord(userID, "执行SQL", tim, sqlOrArgs...) //记录结束时间
	return jsonutil.ObjToJson(result)
}

// QueryPaging 执行SQL分页查询，可传递参数，并返回[]map[string]string
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
func QueryPaging(userID string, IsMasterDB bool, orderBy string, pageIndex int, pageSize int, sqlOrArgs ...interface{}) (lstMap []map[string]string, err error) {

	if len(sqlOrArgs) < 1 {
		return nil, errors.New("分页查询sql为空")
	}
	tim := loghelper.BeginimeRecord() //记录开始时间
	switch confighelper.GetCurrdb() {
	case "mysql":
		lstMap, err = queryPaging_mysql(userID, IsMasterDB, orderBy, pageIndex, pageSize, sqlOrArgs...)
	case "mssql":
		lstMap, err = queryPaging_Sqlserver(userID, IsMasterDB, orderBy, pageIndex, pageSize, sqlOrArgs...)
	default:
		loghelper.ByHighError(logtype.NoDBErr, "请检查config.ini", userID)
		panic("当前数据库未配置,请检查config.ini")
	}
	loghelper.EndTimeRecord(userID, "执行SQL", tim, sqlOrArgs...) //记录结束时间
	return lstMap, err
}

// QueryPaging 执行SQL分页查询，可传递参数，并返回[]map[string]string
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
func QueryPagingTran(session *xorm.Session, userID string, IsMasterDB bool, orderBy string, pageIndex int, pageSize int, sqlOrArgs ...interface{}) (lstMap []map[string]string, err error) {

	if len(sqlOrArgs) < 1 {
		return nil, errors.New("分页查询sql为空")
	}
	tim := loghelper.BeginimeRecord() //记录开始时间
	switch confighelper.GetCurrdb() {
	case "mysql":
		lstMap, err = queryPaging_mysqlTran(session, userID, IsMasterDB, orderBy, pageIndex, pageSize, sqlOrArgs...)
	case "mssql":
		lstMap, err = queryPaging_SqlserverTran(session, userID, IsMasterDB, orderBy, pageIndex, pageSize, sqlOrArgs...)
	default:
		loghelper.ByHighError(logtype.NoDBErr, "请检查config.ini", userID)
		panic("当前数据库未配置,请检查config.ini")
	}
	loghelper.EndTimeRecord(userID, "执行SQL", tim, sqlOrArgs...) //记录结束时间
	return lstMap, err
}

// QueryPaging 执行SQL分页查询，可传递参数，并返回[]map[string]string
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
func QueryPagingByJson(userID string, IsMasterDB bool, orderBy string, pageIndex int, pageSize int, sqlOrArgs ...interface{}) (strJson string, err error) {

	tim := loghelper.BeginimeRecord() //记录开始时间
	switch confighelper.GetCurrdb() {
	case "mysql":
		strJson, err = queryPagingByJson_mysql(userID, IsMasterDB, orderBy, pageIndex, pageSize, sqlOrArgs...)
	case "mssql":
		strJson, err = queryPagingByJson_Sqlserver(userID, IsMasterDB, orderBy, pageIndex, pageSize, sqlOrArgs...)
	default:
		loghelper.ByHighError(logtype.NoDBErr, "请检查config.ini", userID)
		panic("当前数据库未配置,请检查config.ini")
	}
	loghelper.EndTimeRecord(userID, "执行分页查询", tim, sqlOrArgs...) //记录结束时间
	return strJson, err
}

//执行SQL语句，
func ExecSql(userID string, IsMasterDB bool, sqlOrArgs ...interface{}) (bool, error) {
	if sqlOrArgs == nil || len(sqlOrArgs) == 0 {
		return false, nil
	}
	tim := loghelper.BeginimeRecord() //记录开始时间
	db, _ := conn.GetConnection(userID, true)
	_, err := db.Exec(sqlOrArgs...)
	if err != nil {
		var sqlError = concatErr(userID, err, sqlOrArgs...)
		loghelper.ByError(logtype.ExecSqlErr, sqlError, userID)
		return false, err
	}
	loghelper.EndTimeRecord(userID, logtype.ExecSqlErr, tim, sqlOrArgs...) //记录结束时间
	return true, nil
}

//执行SQL语句，
func ExecSqlByTran(session *xorm.Session, userID string, IsMasterDB bool, sqlOrArgs ...interface{}) (bool, error) {

	tim := loghelper.BeginimeRecord() //记录开始时间
	_, err := session.Exec(sqlOrArgs...)
	if err != nil {
		var sqlError = concatErr(userID, err, sqlOrArgs...)
		loghelper.ByError(logtype.ExecSqlErr, sqlError, userID)
		return false, err
	}
	loghelper.EndTimeRecord(userID, logtype.ExecSqlErr, tim, sqlOrArgs...) //记录结束时间
	return true, nil
}

//执行多条SQL语句，
func ExecMoreSql(userID string, IsMasterDB bool, sqlList []string) (bool, error) {
	tim := loghelper.BeginimeRecord() //记录开始时间
	// 用事务执行多条SQL语句，，
	if sqlList == nil || len(sqlList) == 0 {
		// 没有可执行sql 直接返回 true
		return true, nil
	}
	db, _ := conn.GetConnection(userID, true)

	// 开启事务
	session := db.NewSession()
	defer session.Close()
	err := session.Begin()
	for _, value := range sqlList {
		// 依次执行
		_, err := session.Exec(value)
		if err != nil {
			// 执行失败 回滚，并且返回false，记录错误sql
			var sqlError = concatErr(userID, err, value, "")
			loghelper.ByError(logtype.ExecMoreSqlErr, sqlError, userID)
			session.Rollback()
			return false, err
		}
	}
	// 所有都执行完成  提交事务
	err = session.Commit()
	if err != nil {
		var sqlError = concatErr(userID, err, strings.Join(sqlList, ";"))
		loghelper.ByError(logtype.ExecMoreSqlErr, sqlError, userID)
		return false, err
	}
	loghelper.EndTimeRecord(userID, logtype.ExecMoreSqlErr, tim, strings.Join(sqlList, ";")) //记录结束时间
	return true, nil
}

// 执行存储过程返回出参，出参必须包含Result
// ProcName 存储过程名称
// OutArgs  出参名称，不包含@，出参需要定义在入参之后
// ProcInArgs 入参参数列表，按照存储过程指定的顺序传递
func ExecProc_OutParamValue(userID string, IsMasterDB bool, ProcName string, ProcParList []sysmodel.ProcPar) (result map[string]string, err error) {
	tim := loghelper.BeginimeRecord() //记录开始时间
	switch strings.ToLower(conn.DBType) {
	case "mysql":
		// 拼接mysql 存储过程sql 语句
		result, err = execProc_OutParamMysql(userID, IsMasterDB, ProcName, ProcParList)
	case "mssql":
		// 拼接sqlserver 存储过程语句
		result, err = execProc_OutParamSqlserver(userID, IsMasterDB, ProcName, ProcParList)
	}
	slice := make([]interface{}, 0)
	slice = append(slice, ProcName)
	parJson, _ := jsonutil.ObjToJson(ProcParList)
	slice = append(slice, parJson)
	loghelper.EndTimeRecord(userID, "执行存储过程", tim, slice...) //记录结束时间
	return result, err
}

// 执行存储过程，并且返回结果集以及出参
func ExecProc_ResultSetValue(userID string, IsMasterDB bool, ProcName string, ProcParList []sysmodel.ProcPar) (result []map[string]string, outValues map[string]string, err error) {
	tim := loghelper.BeginimeRecord() //记录开始时间
	switch strings.ToLower(conn.DBType) {
	case "mysql":
		result, outValues, err = execProc_ResultSetValue_Mysql(userID, IsMasterDB, ProcName, ProcParList)
	case "mssql":
		result, outValues, err = execProc_ResultSetValue_Sqlserver(userID, IsMasterDB, ProcName, ProcParList)
	}
	slice := make([]interface{}, 0)
	slice = append(slice, ProcName)
	parJson, _ := jsonutil.ObjToJson(ProcParList)
	slice = append(slice, parJson)
	loghelper.EndTimeRecord(userID, "执行存储过程", tim, slice...) //记录结束时间
	return result, outValues, err
}

// QueryFirst 执行SQL查询，可传递参数，并返回首行
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
func QueryFirst(userID string, IsMasterDB bool, sqlOrArgs ...interface{}) (map[string]string, error) {

	tim := loghelper.BeginimeRecord() //记录开始时间
	if sqlOrArgs == nil || len(sqlOrArgs) == 0 {
		// 没有传递sql
		return nil, nil
	}
	var result map[string]string
	var err error
	switch confighelper.GetCurrdb() {
	case "mysql":
		result, err = queryFirst_mysql(userID, IsMasterDB, sqlOrArgs...)
	case "mssql":
		//result, err = queryFirstCol_Sqlserver(userID, IsMasterDB, sqlOrArgs...)
	default:
		loghelper.ByHighError(logtype.NoDBErr, "请检查config.ini", userID)
		panic("当前数据库未配置,请检查config.ini")
	}
	loghelper.EndTimeRecord(userID, logtype.ExecMoreSqlErr, tim, sqlOrArgs...) //记录结束时间

	return result, err
}

// QueryFirst
func QueryFirstByTran(session *xorm.Session, userID string, IsMasterDB bool, sqlOrArgs ...interface{}) (map[string]string, error) {
	tim := loghelper.BeginimeRecord() //记录开始时间
	if sqlOrArgs == nil || len(sqlOrArgs) == 0 {
		// 没有传递sql
		return nil, nil
	}
	var result map[string]string
	var err error
	switch confighelper.GetCurrdb() {
	case "mysql":
		result, err = queryFirst_mysqlTran(session, userID, IsMasterDB, sqlOrArgs...)
	case "mssql":
		//result, err = queryFirstCol_Sqlserver(userID, IsMasterDB, sqlOrArgs...)
	default:
		loghelper.ByHighError(logtype.NoDBErr, "请检查config.ini", userID)
		panic("当前数据库未配置,请检查config.ini")
	}
	loghelper.EndTimeRecord(userID, logtype.ExecMoreSqlErr, tim, sqlOrArgs...) //记录结束时间

	return result, err
}

// QueryFirstCol 执行SQL查询，可传递参数，并返回首行首列
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
func QueryFirstCol(userID string, IsMasterDB bool, sqlOrArgs ...interface{}) (string, error) {

	tim := loghelper.BeginimeRecord() //记录开始时间
	if sqlOrArgs == nil || len(sqlOrArgs) == 0 {
		// 没有传递sql
		return "", nil
	}
	result := ""
	var err error
	switch confighelper.GetCurrdb() {
	case "mysql":
		result, err = queryFirstCol_mysql(userID, IsMasterDB, sqlOrArgs...)
	case "mssql":
		result, err = queryFirstCol_Sqlserver(userID, IsMasterDB, sqlOrArgs...)
	default:
		loghelper.ByHighError(logtype.NoDBErr, "请检查config.ini", userID)
		panic("当前数据库未配置,请检查config.ini")
	}
	loghelper.EndTimeRecord(userID, logtype.ExecMoreSqlErr, tim, sqlOrArgs...) //记录结束时间

	return result, err
}

// QueryFirstCol 执行SQL查询，可传递参数，并返回首行首列
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
func QueryFirstColByTran(session *xorm.Session, userID string, IsMasterDB bool, sqlOrArgs ...interface{}) (string, error) {

	tim := loghelper.BeginimeRecord() //记录开始时间
	if sqlOrArgs == nil || len(sqlOrArgs) == 0 {
		// 没有传递sql
		return "", nil
	}
	result := ""
	var err error
	switch confighelper.GetCurrdb() {
	case "mysql":
		result, err = queryFirstCol_mysqlTran(session, userID, IsMasterDB, sqlOrArgs...)
	case "mssql":
		result, err = queryFirstCol_SqlserverTran(session, userID, IsMasterDB, sqlOrArgs...)
	default:
		loghelper.ByHighError(logtype.NoDBErr, "请检查config.ini", userID)
		panic("当前数据库未配置,请检查config.ini")
	}
	loghelper.EndTimeRecord(userID, logtype.ExecMoreSqlErr, tim, sqlOrArgs...) //记录结束时间

	return result, err
}

//在指定事务中执行SQL语句，并且不进行commit ，也不能rollback
func ExecSqlByTransaction(userID string, connSession xorm.Session, sqlOrArgs ...interface{}) (bool, error) {

	tim := loghelper.BeginimeRecord() //记录开始时间
	_, err := connSession.Exec(sqlOrArgs...)
	if err != nil {
		var sqlError = concatErr(userID, err, sqlOrArgs...)
		loghelper.ByError(logtype.ExecSqlByTranErr, sqlError, userID)
		return false, err
	}
	loghelper.EndTimeRecord(userID, logtype.ExecSqlByTranErr, tim, sqlOrArgs...) //记录结束时间
	return true, nil
}

//执行多条SQL语句，
func ExecMoreSqlByTransaction(userID string, connSession *xorm.Session, sqlList []string) (bool, error) {

	tim := loghelper.BeginimeRecord() //记录开始时间
	if len(sqlList) == 0 {
		return true, nil
	}
	// 遍历执行
	for _, sql := range sqlList {
		_, err := connSession.Exec(sql)
		if err != nil {
			var sqlError = concatErr(userID, err, sql)
			loghelper.ByError(logtype.ExecMoreSqlBrTranErr, sqlError, userID)
			return false, err
		}
	}
	loghelper.EndTimeRecord(userID, logtype.ExecMoreSqlBrTranErr, tim, strings.Join(sqlList, ";")) //记录结束时间
	return true, nil
}

//指定事务中执行S执行存储过程， 存储过程必须有一个名为result 的output输出参数
func ExecProcInTran_OutParamValue(session *xorm.Session, userID string, IsMasterDB bool, ProcName string, ProcParList []sysmodel.ProcPar) (result map[string]string, err error) {

	tim := loghelper.BeginimeRecord() //记录开始时间
	switch confighelper.GetCurrdb() {
	case "mysql":
		return execProcInTran_OutParamMysql(session, userID, IsMasterDB, ProcName, ProcParList)
	case "mssql":
		return execProcInTran_OutParamSqlServer(session, userID, IsMasterDB, ProcName, ProcParList)
	default:
		loghelper.ByHighError(logtype.NoDBErr, "请检查config.ini", userID)
		panic("当前数据库未配置,请检查config.ini")
	}
	slice := make([]interface{}, 0)
	slice = append(slice, ProcName)
	parJson, _ := jsonutil.ObjToJson(ProcParList)
	slice = append(slice, parJson)
	loghelper.EndTimeRecord(userID, logtype.ExecProcByTranErr, tim, slice...) //记录结束时间
	return nil, nil
}

//根据传入数据及表信息，获取insert 语句
//tableInfo 表对象信息 需要对tabletype表类型 进行说明
//userID 用户ID 可为空
//PrimaryKey 主健value或外健值 例：某个 Billid 可为空
//rows 注：传入的数据类型必须为[]map[string]string
//newGuid批次号，每次插表都要有唯一的批次
func BatchInsertByTransaction(tableInfo sysmodel.SqlTableStruct, session *xorm.Session, rows []map[string]interface{}, userID string, ForeignKeyVal string, newGuid string, entid string) (bool, error) {

	//非空检查
	if rows == nil || len(rows) == 0 {
		return true, nil
	}

	tim := loghelper.BeginimeRecord() //记录开始时间

	//生成insert sql pre前缀语句
	var bufferSql bytes.Buffer
	bufferSql.WriteString("insert into ")
	bufferSql.WriteString(tableInfo.TableName)
	bufferSql.WriteString("(")

	//字段清单
	firstRow := rows[0]
	colArr := make([]string, 0, len(firstRow))

	for k, _ := range firstRow {

		//过虑不需要保存的字段 只限表单保存
		if tableInfo.Pid > 0 && (k == "row_id" || k == "rowstate" || strings.Index(k, "_show") != -1) {
			continue
		}

		if (tableInfo.IsIdentity || tableInfo.GridID > 0) && (tableInfo.IsIdentity && k == tableInfo.PrimaryKey) {
			continue
		}

		colArr = append(colArr, k)
		bufferSql.WriteString(k)
		bufferSql.WriteString(",")
	}

	//数据中是否有外健
	var isHaveFk bool = false
	tmp, ok := firstRow[tableInfo.ForeignKey]
	if ok && !g.IsEmpty(tmp) {
		isHaveFk = true
	}

	// insertdate
	bufferSql.WriteString("insertdate")

	bufferSql.WriteString(",entid")

	//pid or griid
	if tableInfo.GridID > 0 {
		bufferSql.WriteString(",currgridid")
		bufferSql.WriteString(",currpid")
		if !isHaveFk && tableInfo.ForeignKey != "" {
			bufferSql.WriteString(",")
			bufferSql.WriteString(tableInfo.ForeignKey)
		}

	}

	bufferSql.WriteString(") values ")
	prefixSql := bufferSql.String()

	// 遍历执行 注：每600行会先插入一次
	isReset := false
	insertCnt := 0
	var insertSqls string
	var bufferInsertSql bytes.Buffer
	for currInd, oneRow := range rows {

		bufferInsertSql.WriteString("(")
		for _, col := range colArr {
			val := oneRow[col]
			if val == "" {
				bufferInsertSql.WriteString("NULL")
			} else {
				bufferInsertSql.WriteString("'")
				bufferInsertSql.WriteString(commutil.ToString(val))
				bufferInsertSql.WriteString("'")
			}
			bufferInsertSql.WriteString(",")
		}

		//insertDate
		bufferInsertSql.WriteString("'")
		bufferInsertSql.WriteString(time.Now().Format(commutil.Time_Fomat01))
		bufferInsertSql.WriteString("'")
		//pidOrGrid
		if tableInfo.GridID > 0 {
			bufferInsertSql.WriteString(",")
			bufferInsertSql.WriteString(entid)
			bufferInsertSql.WriteString(",")
			bufferInsertSql.WriteString(commutil.ToString(tableInfo.GridID))
			bufferInsertSql.WriteString(",")
			bufferInsertSql.WriteString(commutil.ToString(tableInfo.Pid))

			//外健
			if !isHaveFk && tableInfo.ForeignKey != "" {
				bufferInsertSql.WriteString(",'")
				bufferInsertSql.WriteString(ForeignKeyVal)
				bufferInsertSql.WriteString("'")
			}

		}

		bufferInsertSql.WriteString(")")

		insertCnt++

		if insertCnt >= 600 || currInd == len(rows)-1 {
			isReset = true
			insertCnt = 0
		}

		if !isReset {
			bufferInsertSql.WriteString(",")
		}

		//达到600条先进行一次提交
		if isReset {

			insertSql := prefixSql + bufferInsertSql.String()
			//fmt.Println(insertSql)
			_, err := session.Exec(insertSql)
			if err != nil {
				var sqlError = concatErr(userID, err, insertSqls)
				loghelper.ByError(logtype.BatchInsertErr, sqlError, userID)
				return false, err
			}
			bufferInsertSql.Reset()
			isReset = false
		}

	}
	loghelper.EndTimeRecord(userID, logtype.BatchInsertErr, tim, commutil.AppendStr(prefixSql, "数据条数：", commutil.ToString(len(rows)))) //记录结束时间
	return true, nil
}

//根据传入数据及表信息，获取update 语句
//rows 要更新的数据
//PrimaryKey主健ID 注：只用于更新主表时可用，更新子表只根据detailid更新
func BatchUpdateByTransaction(tableInfo sysmodel.SqlTableStruct, session *xorm.Session, rows []map[string]interface{}, userID string) (bool, error) {

	//非空检查
	if rows == nil || len(rows) == 0 {
		return true, nil
	}

	tim := loghelper.BeginimeRecord() //记录开始时间

	//字段清单
	firstRow := rows[0]
	colMap := make(map[string]*bytes.Buffer, 0)

	for k, _ := range firstRow {

		//过虑不需要保存的字段 只限表单保存
		if tableInfo.Pid > 0 && (k == tableInfo.PrimaryKey || k == "update_date" || k == "row_id" || k == "rowstate" || strings.Index(k, "_show") != -1) {
			continue
		}

		buf := bytes.Buffer{}
		colMap[k] = &buf
	}

	var bufferUpdateSql bytes.Buffer
	var primaryKeyBuffer bytes.Buffer
	preFixSql := commutil.AppendStr("update ", tableInfo.TableName, " set ")

	//用于分拼处理记录起始位置 每次更新300条
	rowsCnt := len(rows)
	for startIndex := 0; startIndex <= rowsCnt; {

		colCnt := 0
		//遍历每个字段，为其拼case when
		for col, buf := range colMap {

			buf.WriteString(col)
			buf.WriteString("=case ")
			buf.WriteString(tableInfo.PrimaryKey)
			buf.WriteString(" ")

			for ind := startIndex; ind < startIndex+300 && ind < rowsCnt; ind++ {
				oneRow := rows[ind]
				buf.WriteString(" when ")

				pk_value := commutil.ToString(oneRow[tableInfo.PrimaryKey])
				if tableInfo.IsStringByPrimaryKey {
					pk_value = "'" + pk_value + "'"
				}
				buf.WriteString(pk_value)

				if colCnt == 0 {
					primaryKeyBuffer.WriteString("'")
					primaryKeyBuffer.WriteString(commutil.ToString(oneRow[tableInfo.PrimaryKey]))
					primaryKeyBuffer.WriteString("',")
				}

				val := oneRow[col]
				if val != "" {
					buf.WriteString(" then '")
					buf.WriteString(commutil.ToString(val))
					buf.WriteString("' ")
				} else {
					buf.WriteString(" then NULL ")
				}

				if ind == startIndex+300-1 || ind == rowsCnt-1 {
					buf.WriteString(" end ")
				}

			}

			colCnt++

		}

		//拼接所有字段的case when 为一个update sql
		bufferUpdateSql.WriteString(preFixSql)
		cnt := 0
		for _, buf := range colMap {
			str := buf.String()
			buf.Reset()
			bufferUpdateSql.WriteString(str)
			if cnt < len(colMap)-1 {
				bufferUpdateSql.WriteString(",")
			}
			cnt++
		}

		//update_date
		now := commutil.GetNowTime()
		bufferUpdateSql.WriteString(" ,update_date ='")
		bufferUpdateSql.WriteString(now)
		bufferUpdateSql.WriteString("'")

		//执行update sql 并清空buffer 中的SQL代码
		bufferUpdateSql.WriteString(" where ")
		bufferUpdateSql.WriteString(tableInfo.PrimaryKey)
		bufferUpdateSql.WriteString(" in (")
		bufferUpdateSql.WriteString(strings.TrimRight(primaryKeyBuffer.String(), ","))
		bufferUpdateSql.WriteString(" )")
		upSql := bufferUpdateSql.String()
		//fmt.Println(upSql)
		_, err := session.Exec(upSql)
		if err != nil {
			loghelper.ByError(logtype.BatchUpdateErr, tableInfo.TableName+upSql+err.Error(), userID)
			return false, err
		}
		bufferUpdateSql.Reset()
		primaryKeyBuffer.Reset()
		startIndex = startIndex + 300 //300条update一次
	}

	loghelper.EndTimeRecord(userID, logtype.BatchUpdateErr, tim, commutil.AppendStr(tableInfo.TableName, "数据条数：", commutil.ToString(len(rows)))) //记录结束时间
	return true, nil
}

//根据传入数据及表信息，获取update 语句
//rows 要更新的数据
//PrimaryKey主健ID 注：只用于更新主表时可用，更新子表只根据detailid更新
func BatchUpdateByTransactionByMapString(tableInfo sysmodel.SqlTableStruct, session *xorm.Session, rows []map[string]string, userID string) (bool, error) {

	rowsLen := len(rows)
	//非空检查
	if rows == nil || rowsLen == 0 {
		return true, nil
	}

	tim := loghelper.BeginimeRecord() //记录开始时间

	//字段清单
	firstRow := rows[0]
	colMap := make(map[string]*bytes.Buffer, 0)

	for k, _ := range firstRow {
		if k == tableInfo.PrimaryKey {
			continue
		}
		//过虑不需要保存的字段 只限表单保存
		if tableInfo.Pid > 0 && (k == "update_date" || k == "row_id" || k == "rowstate" || strings.Index(k, "_show") != -1) {
			continue
		}

		buf := bytes.Buffer{}
		colMap[k] = &buf
	}

	var bufferUpdateSql bytes.Buffer
	var primaryKeyBuffer bytes.Buffer
	preFixSql := commutil.AppendStr("update ", tableInfo.TableName, " set ")

	//用于分拼处理记录起始位置 每次更新300条
	rowsCnt := len(rows)
	for startIndex := 0; startIndex <= rowsCnt; {

		colCnt := 0
		//遍历每个字段，为其拼case when
		for col, buf := range colMap {
			if rowsLen == 1 {
				buf.WriteString(col)
				buf.WriteString("=")
				oneRow := rows[0]
				val := oneRow[col]
				if val != "" {
					buf.WriteString(" '")
					buf.WriteString(val)
					buf.WriteString("' ")
				} else {
					buf.WriteString(" NULL ")
				}
				if colCnt == 0 {
					if tableInfo.IsStringByPrimaryKey {
						primaryKeyBuffer.WriteString("'")
						primaryKeyBuffer.WriteString(oneRow[tableInfo.PrimaryKey])
						primaryKeyBuffer.WriteString("'")
					} else {
						primaryKeyBuffer.WriteString(oneRow[tableInfo.PrimaryKey])
					}
					primaryKeyBuffer.WriteString(",")
					colCnt++
				}
				continue
			}

			buf.WriteString(col)
			buf.WriteString("=case ")
			buf.WriteString(tableInfo.PrimaryKey)
			buf.WriteString(" ")

			for ind := startIndex; ind < startIndex+300 && ind < rowsCnt; ind++ {
				oneRow := rows[ind]
				buf.WriteString(" when ")

				pk_value := oneRow[tableInfo.PrimaryKey]
				if tableInfo.IsStringByPrimaryKey {
					pk_value = "'" + pk_value + "'"
				}
				buf.WriteString(pk_value)

				if colCnt == 0 {
					if tableInfo.IsStringByPrimaryKey {
						primaryKeyBuffer.WriteString("'")
						primaryKeyBuffer.WriteString(oneRow[tableInfo.PrimaryKey])
						primaryKeyBuffer.WriteString("'")
					} else {
						primaryKeyBuffer.WriteString(oneRow[tableInfo.PrimaryKey])
					}
					primaryKeyBuffer.WriteString(",")
				}

				val := oneRow[col]
				if val != "" {
					buf.WriteString(" then '")
					buf.WriteString(val)
					buf.WriteString("' ")
				} else {
					buf.WriteString(" then NULL ")
				}

				if ind == startIndex+300-1 || ind == rowsCnt-1 {
					buf.WriteString(" end ")
				}

			}

			colCnt++

		}

		//拼接所有字段的case when 为一个update sql
		bufferUpdateSql.WriteString(preFixSql)
		cnt := 0
		for _, buf := range colMap {
			str := buf.String()
			buf.Reset()
			bufferUpdateSql.WriteString(str)
			if cnt < len(colMap)-1 {
				bufferUpdateSql.WriteString(",")
			}
			cnt++
		}

		if tableInfo.Pid > 0 {
			//update_date
			now := commutil.GetNowTime()
			bufferUpdateSql.WriteString(" ,update_date ='")
			bufferUpdateSql.WriteString(now)
			bufferUpdateSql.WriteString("'")
		}

		//执行update sql 并清空buffer 中的SQL代码
		bufferUpdateSql.WriteString(" where ")
		bufferUpdateSql.WriteString(tableInfo.PrimaryKey)
		bufferUpdateSql.WriteString(" in (")
		bufferUpdateSql.WriteString(strings.TrimRight(primaryKeyBuffer.String(), ","))
		bufferUpdateSql.WriteString(" )")
		upSql := bufferUpdateSql.String()
		//fmt.Println(upSql)
		_, err := session.Exec(upSql)
		if err != nil {
			loghelper.ByError(logtype.BatchUpdateErr, tableInfo.TableName+upSql+err.Error(), userID)
			return false, err
		}
		bufferUpdateSql.Reset()
		primaryKeyBuffer.Reset()
		startIndex = startIndex + 300 //300条update一次
	}

	loghelper.EndTimeRecord(userID, logtype.BatchUpdateErr, tim, commutil.AppendStr(tableInfo.TableName, "数据条数：", commutil.ToString(len(rows)))) //记录结束时间
	return true, nil
}

//-------多个字段作为更新主键
//根据传入数据及表信息，获取update 语句
//rows 要更新的数据
//PrimaryKey主健ID 注：只用于更新主表时可用，更新子表只根据detailid更新
func BatchUpdateByTransactionByMorePrimarykeys(MorePrimarykeys []string, tableInfo sysmodel.SqlTableStruct, session *xorm.Session, rows []map[string]string, userID string) (bool, error) {
	if len(MorePrimarykeys) == 0 {
		return false, errors.New("更新维度字段不能为空！")
	}

	//非空检查
	if rows == nil || len(rows) == 0 {
		return true, nil
	}

	tim := loghelper.BeginimeRecord() //记录开始时间

	//字段清单
	firstRow := rows[0]
	colMap := make(map[string]*bytes.Buffer, 0)
	tableInfo.PrimaryKey = MorePrimarykeys[0]

	for k, _ := range firstRow {

		//过虑不需要保存的字段 只限表单保存
		if tableInfo.Pid > 0 && (k == tableInfo.PrimaryKey || k == "update_date" || k == "row_id" || k == "rowstate" || strings.Index(k, "_show") != -1) {
			continue
		}

		buf := bytes.Buffer{}
		colMap[k] = &buf
	}

	var bufferUpdateSql bytes.Buffer
	var primaryKeyBuffer bytes.Buffer
	var pk_value string
	preFixSql := commutil.AppendStr("update ", tableInfo.TableName, " set ")

	//用于分拼处理记录起始位置 每次更新300条
	rowsCnt := len(rows)
	for startIndex := 0; startIndex <= rowsCnt; {

		colCnt := 0
		//遍历每个字段，为其拼case when
		for col, buf := range colMap {

			buf.WriteString(col)
			buf.WriteString("=case ")

			for ind := startIndex; ind < startIndex+300 && ind < rowsCnt; ind++ {
				oneRow := rows[ind]
				buf.WriteString(" when ")

				//更新条件
				for index, key := range MorePrimarykeys {
					if index > 0 {
						buf.WriteString(" and ")
						if colCnt == 0 {
							primaryKeyBuffer.WriteString(" and ")
						}
					}
					buf.WriteString(key)
					buf.WriteString(" = ")

					pk_value = commutil.ToString(oneRow[key])
					if tableInfo.IsStringByPrimaryKey {
						pk_value = "'" + pk_value + "'"
					}
					buf.WriteString(pk_value)

					if colCnt == 0 {
						if index == 0 {
							primaryKeyBuffer.WriteString(" ( ")
						}
						primaryKeyBuffer.WriteString(key)
						primaryKeyBuffer.WriteString("='")
						primaryKeyBuffer.WriteString(commutil.ToString(oneRow[key]))
						primaryKeyBuffer.WriteString("' ")
					}
				}
				if colCnt == 0 {
					primaryKeyBuffer.WriteString(" ) or")
				}

				val := oneRow[col]
				if val != "" {
					buf.WriteString(" then '")
					buf.WriteString(commutil.ToString(val))
					buf.WriteString("' ")
				} else {
					buf.WriteString(" then NULL ")
				}

				if ind == startIndex+300-1 || ind == rowsCnt-1 {
					buf.WriteString(" end ")
				}

			}

			colCnt++

		}

		//拼接所有字段的case when 为一个update sql
		bufferUpdateSql.WriteString(preFixSql)
		cnt := 0
		for _, buf := range colMap {
			str := buf.String()
			buf.Reset()
			bufferUpdateSql.WriteString(str)
			if cnt < len(colMap)-1 {
				bufferUpdateSql.WriteString(",")
			}
			cnt++
		}

		//update_date
		now := commutil.GetNowTime()
		bufferUpdateSql.WriteString(" ,update_date ='")
		bufferUpdateSql.WriteString(now)
		bufferUpdateSql.WriteString("'")

		//执行update sql 并清空buffer 中的SQL代码
		bufferUpdateSql.WriteString(" where ")
		bufferUpdateSql.WriteString(strings.TrimRight(primaryKeyBuffer.String(), "or"))
		//bufferUpdateSql.WriteString(tableInfo.PrimaryKey)
		//bufferUpdateSql.WriteString(" in (")
		//bufferUpdateSql.WriteString(" )")
		upSql := bufferUpdateSql.String()
		//fmt.Println(upSql)
		_, err := session.Exec(upSql)
		if err != nil {
			loghelper.ByError(logtype.BatchUpdateErr, tableInfo.TableName+upSql+err.Error(), userID)
			return false, err
		}
		bufferUpdateSql.Reset()
		primaryKeyBuffer.Reset()
		startIndex = startIndex + 300 //300条update一次
	}

	loghelper.EndTimeRecord(userID, logtype.BatchUpdateErr, tim, commutil.AppendStr(tableInfo.TableName, "数据条数：", commutil.ToString(len(rows)))) //记录结束时间
	return true, nil
}

//根据传入数据及表信息，获取insert 语句
//tableInfo 表对象信息 需要对tabletype表类型 进行说明
//userID 用户ID 可为空
//PrimaryKey 主健value或外健值 例：某个 Billid 可为空
//rows 注：传入的数据类型必须为[]map[string]string
//newGuid批次号，每次插表都要有唯一的批次
func BatchInsertByMapString(tabName string, session *xorm.Session, rows []map[string]string, userID string) (bool, error) {

	//非空检查
	if rows == nil || len(rows) == 0 {
		return true, nil
	}

	tim := loghelper.BeginimeRecord() //记录开始时间

	//生成insert sql pre前缀语句
	var bufferSql bytes.Buffer
	bufferSql.WriteString("insert into ")
	bufferSql.WriteString(tabName)
	bufferSql.WriteString("(")

	//字段清单
	firstRow := rows[0]
	colArr := make([]string, 0, len(firstRow))

	cnt := len(firstRow)
	index := 0
	for k, _ := range firstRow {
		colArr = append(colArr, k)
		bufferSql.WriteString(k)
		if index < cnt-1 {
			bufferSql.WriteString(",")
		}
		index++
	}

	bufferSql.WriteString(") values ")
	prefixSql := bufferSql.String()

	// 遍历执行 注：每600行会先插入一次
	isReset := false
	insertCnt := 0
	oneInsertMaxRowCnt := 600
	var insertSqls string
	var bufferInsertSql bytes.Buffer
	for currInd, oneRow := range rows {

		bufferInsertSql.WriteString("(")
		index = 0
		for _, col := range colArr {
			val := oneRow[col]
			if val == "" {
				bufferInsertSql.WriteString("NULL")
			} else {
				bufferInsertSql.WriteString("'")
				bufferInsertSql.WriteString(val)
				bufferInsertSql.WriteString("'")
			}

			if index < cnt-1 {
				bufferInsertSql.WriteString(",")
			}
			index++
		}

		bufferInsertSql.WriteString(")")

		insertCnt++

		if insertCnt >= oneInsertMaxRowCnt || currInd == len(rows)-1 {
			isReset = true
			insertCnt = 0
		}

		if !isReset {
			bufferInsertSql.WriteString(",")
		}

		//达到600条先进行一次提交
		if isReset {

			insertSql := prefixSql + bufferInsertSql.String()
			//fmt.Println(insertSql)
			_, err := session.Exec(insertSql)
			if err != nil {
				var sqlError = concatErr(userID, err, insertSqls)
				loghelper.ByError(logtype.BatchInsertErr, sqlError, userID)
				return false, err
			}
			bufferInsertSql.Reset()
			isReset = false
		}

	}
	loghelper.EndTimeRecord(userID, logtype.BatchInsertErr, tim, commutil.AppendStr(prefixSql, "数据条数：", commutil.ToString(len(rows)))) //记录结束时间
	return true, nil
}

//获取ifnull关健字
func GetIFNull() string {
	switch strings.ToLower(conn.DBType) {
	case "mysql":
		return "ifnull"
	case "mssql":
		return "isnull"
	default:
		return "ifnull"
	}

}

// 解析前台ajax传递的查询字段，并转为sql及查询参数
func GetFilterSql(lstQueryField []sysmodel.QueryField, IsOpenRealCtr bool, pid int, isParamsTran bool) (filterSql string, queryPars []interface{}) {
	queryPars = make([]interface{}, 0)
	var bufferWhere bytes.Buffer
	isfirst := true
	for _, entity := range lstQueryField {
		if g.IsEmpty(entity.KeyWord) {
			continue
		}
		if !isfirst {
			bufferWhere.WriteString(" and ")
		}
		isfirst = false
		//50508 表示凭证导入界面
		//50504 快捷付款 (员工 、供应商）
		//50609  快捷付款 (汇票)
		isSelectReal := !strings.Contains(entity.Op, "like") && (entity.FieldName == "billno" || entity.FieldName == "realnumber") && (IsOpenRealCtr || pid == 50508 || pid == 50504 || pid == 50609)
		GetFilterParamsToListSelect(&entity, isParamsTran)
		if isSelectReal {
			//单据启用了信封号功能 ，需要将两个字段联合查询
			var billno, realnumber string
			if strings.Contains(entity.FieldName, "billno") {
				realnumber = strings.ReplaceAll(entity.FieldName, "billno", "realnumber")
				billno = entity.FieldName
			} else {
				billno = strings.ReplaceAll(entity.FieldName, "realnumber", "billno")
				realnumber = entity.FieldName
			}

			if entity.Op == "in" || entity.Op == "not in" {
				var inBuffer strings.Builder
				inBuffer.WriteString(" (")
				lst := strings.Split(entity.KeyWord, ",")
				// queryPars2 用于分割流水号在不同占位符，否则直接添加在 queryPars
				// 查询1,2 会变成 billno in (1,1) or realnumber in (2,2)
				var queryPars2 []interface{}
				isfirstcol := true
				for _, s := range lst {
					if s == "" || strings.Trim(s, " ") == "" {
						continue
					}
					if !isfirstcol {
						inBuffer.WriteString(",")
					}
					isfirstcol = false
					queryPars = append(queryPars, s)   // 两个字段 需要复制一个 s
					queryPars2 = append(queryPars2, s) // 两个字段 需要复制一个 s
					inBuffer.WriteString("?")
				}
				queryPars = append(queryPars, queryPars2...)
				inBuffer.WriteString(" )")
				bufferWhere.WriteString(" (")

				bufferWhere.WriteString(" ")
				bufferWhere.WriteString(billno)
				bufferWhere.WriteString(" ")
				bufferWhere.WriteString(entity.Op)
				bufferWhere.WriteString(inBuffer.String())

				bufferWhere.WriteString(" or ")
				bufferWhere.WriteString(realnumber)
				bufferWhere.WriteString(" ")
				bufferWhere.WriteString(entity.Op)
				bufferWhere.WriteString(inBuffer.String())

				bufferWhere.WriteString(" )")

			} else {
				bufferWhere.WriteString(" (")
				bufferWhere.WriteString(" ")
				bufferWhere.WriteString(billno)
				bufferWhere.WriteString(" ")
				bufferWhere.WriteString(entity.Op)
				bufferWhere.WriteString(" ? ")

				bufferWhere.WriteString(" or ")

				bufferWhere.WriteString(" ")
				bufferWhere.WriteString(realnumber)
				bufferWhere.WriteString(" ")
				bufferWhere.WriteString(entity.Op)
				bufferWhere.WriteString(" ? ")

				bufferWhere.WriteString(" )")

				queryPars = append(queryPars, entity.KeyWord) // 两个字段 需要复制一个 s
				queryPars = append(queryPars, entity.KeyWord)
			}

			continue

		} else {

			bufferWhere.WriteString(" ")
			bufferWhere.WriteString(entity.FieldName)
			bufferWhere.WriteString(" ")
			bufferWhere.WriteString(entity.Op)

			if entity.Op == "in" || entity.Op == "not in" {
				bufferWhere.WriteString(" (")
				lst := strings.Split(entity.KeyWord, ",")
				isfirstcol := true
				for _, s := range lst {
					if s == "" || strings.Trim(s, " ") == "" {
						continue
					}
					if !isfirstcol {
						bufferWhere.WriteString(",")
					}
					isfirstcol = false
					queryPars = append(queryPars, s)
					bufferWhere.WriteString("?")
				}
				bufferWhere.WriteString(" )")
			} else {
				// %3232%
				if entity.Op == "like" || entity.Op == "not like" {
					bufferWhere.WriteString("?")
				} else {
					bufferWhere.WriteString(" ? ")
				}
				if entity.Op == "<" || entity.Op == "<=" || entity.Op == ">" || entity.Op == ">=" {
					if strings.Contains(entity.SqlDatatype, "int") {
						queryPars = append(queryPars, commutil.ToInt64(entity.KeyWord))
					} else if strings.Contains(entity.SqlDatatype, "float") {
						queryPars = append(queryPars, commutil.ToFloat64(entity.KeyWord))
					} else if strings.Contains(entity.SqlDatatype, "decimal") {
						queryPars = append(queryPars, commutil.ToFloat64(entity.KeyWord))
					} else {
						//其它类型 例：date varchar
						queryPars = append(queryPars, entity.KeyWord)
					}

				} else {
					queryPars = append(queryPars, entity.KeyWord)
				}
			}
		}

	}
	if bufferWhere.Cap() > 0 {
		return strings.TrimRight(bufferWhere.String(), "and"), queryPars
	} else {
		return "", nil
	}
}

//先使用map存储需要支持分割查询的字段 , 待后面支持多组合条件查询 时再取消
var listMapSelect = map[string]interface{}{
	"billno": nil, "sourcebillno": nil, "realnumber": nil, "pid": nil,
	"csid": nil, "cscode": nil, "csname": nil,
	"costid": nil, "costname": nil, "costcode": nil,
	"deptid": nil, "deptcode": nil, "deptname": nil, "storecode": nil,
	"userid": nil, "username": nil, "usercode": nil,
	"centerid": nil, "centercode": nil, "centername": nil,
	"codeid": nil, "code": nil, "codename": nil,
}

// like = 比较符号 转成list查询 tran 是否转换
func GetFilterParamsToListSelect(field *sysmodel.QueryField, tran bool) (isToListSelect bool) {
	FieldName := strings.ToLower(field.FieldName)
	if field.Op == "" || field.KeyWord == "" || strings.Contains(FieldName, "memo") {
		return isToListSelect
	}
	//是否转 in 查询， 空格或,逗号拆分时
	if field.Op == "like" || field.Op == "=" {
		sep1 := strings.Contains(field.KeyWord, ",")
		sep2 := strings.Contains(field.KeyWord, " ")
		if !sep1 && !sep2 {
			if field.Op == "like" {
				if !strings.Contains(field.KeyWord, "%") {
					field.KeyWord = "%" + field.KeyWord + "%"
				}
			}
			return isToListSelect
		}
		if tran {
			sepStr := ","
			if sep2 {
				sepStr = " "
			}
			field.Op = "in"
			if sep2 {
				field.KeyWord = strings.ReplaceAll(field.KeyWord, sepStr, ",")
			}
			//替换 %  ， 空格替换为, 查询
			field.KeyWord = strings.ReplaceAll(field.KeyWord, "%", "")
			isToListSelect = true
		} else {
			// 列表页或拷贝页 打开界面，部分字段支持分割查询
			FieldName = strings.ReplaceAll(FieldName, "billmain.", "")
			FieldName = strings.ReplaceAll(FieldName, "detail.", "")
			FieldName = strings.ReplaceAll(FieldName, "basemain.", "")
			FieldName = strings.ReplaceAll(FieldName, "main.", "")
			if _, isok := listMapSelect[FieldName]; !isok {
				return isToListSelect
			}
			field.Op = "in"
			//替换 %  ， 空格替换为, 查询
			field.KeyWord = strings.ReplaceAll(field.KeyWord, "%", "")
			isToListSelect = true
		}
	}
	return isToListSelect
}

func row2mapStr(rows *sql.Rows, fields []string) (resultsMap map[string]string, err error) {
	result := make(map[string]string)
	scanResultContainers := make([]interface{}, len(fields))
	for i := 0; i < len(fields); i++ {
		var scanResultContainer interface{}
		scanResultContainers[i] = &scanResultContainer
	}
	if err := rows.Scan(scanResultContainers...); err != nil {
		return nil, err
	}

	for ii, key := range fields {
		rawValue := reflect.Indirect(reflect.ValueOf(scanResultContainers[ii]))
		// if row is null then as empty string
		if rawValue.Interface() == nil {
			result[key] = ""
			continue
		}

		if data, err := value2String(&rawValue); err == nil {
			result[key] = data
		} else {
			return nil, err
		}
	}
	return result, nil
}

func value2String(rawValue *reflect.Value) (str string, err error) {
	aa := reflect.TypeOf((*rawValue).Interface())
	vv := reflect.ValueOf((*rawValue).Interface())
	switch aa.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str = strconv.FormatInt(vv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str = strconv.FormatUint(vv.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		str = strconv.FormatFloat(vv.Float(), 'f', -1, 64)
	case reflect.String:
		str = vv.String()
	case reflect.Array, reflect.Slice:
		switch aa.Elem().Kind() {
		case reflect.Uint8:
			data := rawValue.Interface().([]byte)
			str = string(data)
			if str == "\x00" {
				str = "0"
			}
		default:
			err = fmt.Errorf("Unsupported struct type %v", vv.Type().Name())
		}
	// time type
	case reflect.Struct:
		if aa.ConvertibleTo(core.TimeType) {
			str = vv.Convert(core.TimeType).Interface().(time.Time).Format(time.RFC3339Nano)
		} else {
			err = fmt.Errorf("Unsupported struct type %v", vv.Type().Name())
		}
	case reflect.Bool:
		str = strconv.FormatBool(vv.Bool())
	case reflect.Complex128, reflect.Complex64:
		str = fmt.Sprintf("%v", vv.Complex())
	default:
		err = fmt.Errorf("Unsupported struct type %v", vv.Type().Name())
	}
	return
}

func SelectMaxPrimaryKey(userID string, pkCount int, tableName string) (maxPrimary int) {

	if commutil.IsNullOrEmpty(tableName) {
		panic("PublicTip_ParameterIsNull=============参数不能为空")
	}
	// //调用获得最大主键的存储过程，输出参数为表名
	// UP_SYS_GetBillId 为oracle数据库里获得最大主键的存储过程
	var procList []sysmodel.ProcPar

	// 入参 设值 参数名称 值
	var parIn = sysmodel.ProcPar{"TableName", "string", 100, false, tableName}
	// 入参 设值 参数名称 值
	var parIn1 = sysmodel.ProcPar{"pri_count", "int", 100, false, commutil.ToString(pkCount)}

	// 出参 参数名称 类型
	var parOut = sysmodel.ProcPar{"BillID", "int", 100, true, ""}

	// 出参 参数名称 类型
	var parOut1 = sysmodel.ProcPar{"R_esult", "string", 100, true, ""}

	procList = append(procList, parIn)
	procList = append(procList, parIn1)
	procList = append(procList, parOut)
	procList = append(procList, parOut1)

	enigne, _ := conn.GetConnection(userID, false)

	// 调用存储过程
	session := enigne.NewSession()
	out, _ := execProcInTran_OutParamMysql(session, userID, false, "UP_SYS_GetBillId", procList)

	if len(out) == 0 {
		panic("PublicTip_QueryDataFailed =================查询数据失败")
	} else {
		maxPrimary = commutil.ToInt(out["BillID"])
	}

	return maxPrimary
}

func QueryMap(userID string, IsMasterDB bool, keyStr string, valueStr string, sqlOrArgs ...interface{}) (map[string]string, error) {

	// 通用方法无需区分数据库类型
	tim := loghelper.BeginimeRecord() //记录开始时间
	engine, _ := conn.GetConnection(userID, IsMasterDB)
	results, err := engine.QueryString(sqlOrArgs...)

	if err != nil {
		var sqlError = concatErr(userID, err, sqlOrArgs...)
		loghelper.ByError(logtype.QueryErr, sqlError, userID)
		return nil, err
	}
	result := make(map[string]string, 0)
	for _, row := range results {
		keyValue := row[keyStr]
		value := row[valueStr]
		result[keyValue] = value
	}

	loghelper.EndTimeRecord(userID, "执行SQL", tim, sqlOrArgs...) //记录结束时间
	return result, nil
}

//sqlOrArgs 查询sql和参数， 返回map[string]map[string]string
//keyStr 作为map[string]的主键  map[string]string 是数据行
func QueryMapByKeyMap(userID string, IsMasterDB bool, keyStr string, sqlOrArgs ...interface{}) (map[string]map[string]string, error) {
	// 通用方法无需区分数据库类型
	tim := loghelper.BeginimeRecord() //记录开始时间
	engine, _ := conn.GetConnection(userID, IsMasterDB)
	session := engine.NewSession()
	defer session.Close()
	rows, err := session.QueryRows(sqlOrArgs...)
	if err != nil {
		var sqlError = concatErr(userID, err, sqlOrArgs...)
		loghelper.ByError(logtype.QueryErr, sqlError, userID)
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	compareFieldByMap := make(map[string]map[string]string)
	for rows.Next() {
		err = row2listmapStrByKeyCol(rows, columns, keyStr, compareFieldByMap)
		if err != nil {
			return nil, err
		}
	}

	loghelper.EndTimeRecord(userID, "执行SQL", tim, sqlOrArgs...) //记录结束时间
	return compareFieldByMap, nil
}

// 转换行数据map[key]map[string]string, 外层map key等于 keyCol 字段的值
func row2listmapStrByKeyCol(rs *xormplus_core.Rows, columns []string, keyCol string, ListMapByKey map[string]map[string]string) (err error) {
	scanResultContainers := make([]interface{}, len(columns))
	//vvv := vv.Elem()
	for i := range columns {
		//scanResultContainers[ii] = rs.db.reflectNew(vvv.Type().Elem()).Interface()
		var scanResultContainer interface{}
		scanResultContainers[i] = &scanResultContainer
	}

	err = rs.Rows.Scan(scanResultContainers...)
	if err != nil {
		return err
	}

	keybuf := &bytes.Buffer{}
	result := map[string]string{}
	for ii, key := range columns {
		vname := reflect.ValueOf(key).String()
		result[vname] = reflect.ValueOf(scanResultContainers[ii]).String()
		//vvv.SetMapIndex(vname, reflect.ValueOf(scanResultContainers[ii]).Elem())

		rawValue := reflect.Indirect(reflect.ValueOf(scanResultContainers[ii]))
		// if row is null then as empty string
		if rawValue.Interface() == nil {
			result[key] = ""
			continue
		}

		if data, err := value2String(&rawValue); err == nil {
			result[key] = data
			if key == keyCol {
				keybuf.WriteString(data)
			}
		} else {
			return err
		}
	}

	ListMapByKey[keybuf.String()] = result
	return nil
}

/**
根据 struct 结构自动插入表单
*/
func InsertStruct(userID string, bean interface{}) (bool, error) {
	engine, err := conn.GetConnection(userID, true)
	if err != nil {
		fmt.Errorf("获取数据库链接失败：%s", err.Error())
		return false, err
	}
	session := engine.NewSession()
	_, err = session.Insert(bean)
	if err != nil {
		fmt.Errorf("执行插入失败：%s", err.Error())
		return false, err
	}
	return true, nil
}

/**
根据 struct 结构更新数据
*/
func UpdateStruct(userID string, bean interface{}, query interface{}, args ...interface{}) (int64, error) {
	engine, err := conn.GetConnection(userID, true)
	if err != nil {
		fmt.Errorf("获取数据库链接失败：%s", err.Error())
		return 0, err
	}
	session := engine.NewSession()
	return session.Where(query, args...).Update(bean)
}

func InsertStructMulti(userID string, bean []interface{}) (bool, error) {
	engine, err := conn.GetConnection(userID, true)
	if err != nil {
		fmt.Errorf("获取数据库链接失败：%s", err.Error())
		return false, err
	}
	session := engine.NewSession()
	_, err = session.InsertMulti(bean)
	if err != nil {
		fmt.Errorf("批量执行插入失败：%s", err.Error())
		return false, err
	}
	return true, nil
}
