package dbhelper

import (
	"base/base/db/conn"
	"base/base/loghelper"
	"base/base/sysmodel"
	"base/base/sysmodel/logtype"
	"base/base/util/commutil"
	"base/base/util/jsonutil"
	"bytes"
	"errors"
	"strings"

	"github.com/xormplus/xorm"
)

// QueryPaging 执行SQL分页查询，可传递参数，并返回[]map[string]string
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
// pageIndex 第一页传0
func queryPaging_Sqlserver(userID string, IsMasterDB bool, orderBy string, pageIndex int, pageSize int, sqlOrArgs ...interface{}) (lstMap []map[string]string, err error) {
	// QueryPaging 执行SQL分页查询，
	if len(sqlOrArgs) < 1 {
		return nil, errors.New("参数为空")
	}
	if orderBy == "" {
		// order by 语句为空
		return nil, errors.New("order by 为空")
	}
	db, err := conn.GetConnection(userID, IsMasterDB)
	if err != nil {

		var sqlError = concatErr(userID, err, "连接数据库失败")
		loghelper.ByError(logtype.GetConnErr, sqlError, userID)
		return nil, err
	}
	slice := concatSqlserverPageSql(orderBy, pageIndex, pageSize, sqlOrArgs)
	result, err := db.QueryString(slice...)
	if err != nil {

		var sqlError = concatErr(userID, err, slice...)
		loghelper.ByError(logtype.QueryErr, sqlError, userID)
		return nil, err
	}
	return result, nil
}

// QueryPaging 执行SQL分页查询，可传递参数，并返回[]map[string]string
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
// pageIndex 第一页传0
func queryPaging_SqlserverTran(session *xorm.Session, userID string, IsMasterDB bool, orderBy string, pageIndex int, pageSize int, sqlOrArgs ...interface{}) (lstMap []map[string]string, err error) {
	// QueryPaging 执行SQL分页查询，
	if len(sqlOrArgs) < 1 {
		return nil, errors.New("参数为空")
	}
	if orderBy == "" {
		// order by 语句为空
		return nil, errors.New("order by 为空")
	}

	slice := concatSqlserverPageSql(orderBy, pageIndex, pageSize, sqlOrArgs)
	result, err := session.QueryString(slice...)
	if err != nil {

		var sqlError = concatErr(userID, err, slice...)
		loghelper.ByError(logtype.QueryErr, sqlError, userID)
		return nil, err
	}
	return result, nil
}

// QueryPaging 执行SQL分页查询，可传递参数，并返回[]map[string]string
// userID 用户ID
// IsMasterDB 是否从masterDB读取，否则会自动选slave 库读取
// SqlAndArgs 表示SQL语句和参数（参数为可选项 例 :  select * from xx where a=? ,123 就需要两个参数
func queryPagingByJson_Sqlserver(userID string, IsMasterDB bool, orderBy string, pageIndex int, pageSize int, sqlOrArgs ...interface{}) (strJson string, err error) {

	// QueryPaging 执行SQL分页查询，
	// QueryPaging 执行SQL分页查询，
	if len(sqlOrArgs) < 1 {
		return "", errors.New("参数为空")
	}
	db, err := conn.GetConnection(userID, IsMasterDB)
	if err != nil {

		return "", err
	}
	// 将替换后的sql 重新组装成列表供xorm 调用
	slice := concatSqlserverPageSql(orderBy, pageIndex, pageSize, sqlOrArgs)
	result, err := db.QueryString(slice...)
	if err != nil {

		var sqlError = concatErr(userID, err, slice...)
		loghelper.ByError(logtype.QueryErr, sqlError, userID)
		return "", err
	}
	json, err := jsonutil.ObjToJson(result)
	if err != nil {
		return "", err
	}
	return json, nil
}

//查询首行首列
func queryFirstCol_Sqlserver(userID string, IsMasterDB bool, sqlOrArgs ...interface{}) (result string, err error) {

	sql := commutil.ToString(sqlOrArgs[0])
	buffer := new(bytes.Buffer)
	buffer.WriteString("set rowcount 1;")
	buffer.WriteString(sql)
	buffer.WriteString("; set rowcount 0;")
	sqlOrArgs[0] = buffer.String()

	db, err := conn.GetConnection(userID, IsMasterDB)
	if err != nil {

		return "", err
	}

	rows, err := db.QueryString(sqlOrArgs...)
	if err != nil {
		loghelper.ByError(logtype.QueryErr, commutil.AppendStr(err.Error(), sqlOrArgs[0]), userID)
		return "", err
	}

	if len(rows) > 0 {
		for _, v := range rows[0] {
			return v, nil
		}
	} else {
		return "", nil
	}
	return "", nil
}

//查询首行首列
func queryFirstCol_SqlserverTran(session *xorm.Session, userID string, IsMasterDB bool, sqlOrArgs ...interface{}) (result string, err error) {

	sql := commutil.ToString(sqlOrArgs[0])
	buffer := new(bytes.Buffer)
	buffer.WriteString("set rowcount 1;")
	buffer.WriteString(sql)
	buffer.WriteString("; set rowcount 0;")
	sqlOrArgs[0] = buffer.String()

	rows, err := session.QueryString(sqlOrArgs...)
	if err != nil {
		loghelper.ByError(logtype.QueryErr, commutil.AppendStr(err.Error(), sqlOrArgs[0]), userID)
		return "", err
	}

	if len(rows) > 0 {
		for _, v := range rows[0] {
			return v, nil
		}
	} else {
		return "", nil
	}
	return "", nil
}

// sql:  set rowcount pagesize
// select * from (select row_number() over (order by orderbysql) as rownumber,* from (select xxx,xxx from fromsql where xxx )a )b where rownumber > start
// set rowcount 0
func concatSqlserverPageSql(OrderBy string, PageIndex int, PageSize int, sqlOrAgrs ...interface{}) (slice []interface{}) {
	var start = PageIndex * PageSize
	buffer := new(bytes.Buffer)
	buffer.WriteString("set rowcount ")
	buffer.WriteString(commutil.ToString(PageSize))
	buffer.WriteString(";")
	buffer.WriteString(" select * from (")
	buffer.WriteString("select row_number() over (order by " + OrderBy + ") as rownumber,* from (")
	buffer.WriteString(commutil.ToString(sqlOrAgrs[0]))
	buffer.WriteString(")a")
	buffer.WriteString(")b where rownumber>" + commutil.ToString(start))
	buffer.WriteString("; set rowcount 0;")

	// 将替换后的sql 重新组装成列表供xorm 调用
	slice = append(slice, buffer.String())
	slice = append(slice, sqlOrAgrs[1:]...)
	return slice
}

// 调用此存储过程用于返回出参，必须包含result出参,如果没有则不符合规范，入参可以为空
// 调用示例 execProc_OutParamsqlserver("11",  ==userid
// true,  == IsMaster
// "UP_SYS_GetBillId",  == ProcName
// []string{"billid","result"},  == OutArgs
// "EB_Role",1)	== ProcInArgs
func execProc_OutParamSqlserver(userID string, IsMasterDB bool, ProcName string, ProcParList []sysmodel.ProcPar) (result map[string]string, err error) {
	if ProcName == "" {
		return nil, errors.New("存储过程名为空")
	}
	if len(ProcParList) < 1 {
		return nil, errors.New("存储过程出参至少需要包含result")
	}
	queryStr, err := concatSqlserverExeProc(ProcName, ProcParList)
	if err != nil { // 执行proc 参数不对
		return nil, err
	}
	db, _ := conn.GetConnection(userID, IsMasterDB)
	res, err := db.QueryString(queryStr)
	if err != nil {

		var sqlError = concatErr(userID, err, queryStr)
		loghelper.ByError(logtype.ExecProcErr, sqlError, userID)
		return nil, err
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, errors.New("未获取到存储过程 " + ProcName + " 返回值")
}

// 拼接sqlserver 存储过程 调用sql
func concatSqlserverExeProc(ProcName string, ProcParList []sysmodel.ProcPar) (callSql string, err error) {
	// 第一个 调用 call proc 语句
	inBuffer := new(bytes.Buffer)
	inBuffer.WriteString("exec ")
	inBuffer.WriteString(ProcName)
	inBuffer.WriteString(" ")

	// 第二个查询result
	outBuffer := new(bytes.Buffer)
	outBuffer.WriteString("select ")

	outParBuffer := new(bytes.Buffer)

	// 遍历出参
	var isResultout = false

	for _, procPar := range ProcParList {
		if !procPar.IsOutPut {
			// 入参
			inBuffer.WriteString("'")
			inBuffer.WriteString(procPar.ProcValue)
			inBuffer.WriteString("',")
		} else {
			// 出参
			inBuffer.WriteString("@")
			inBuffer.WriteString(procPar.ProcName)
			inBuffer.WriteString("=@")
			inBuffer.WriteString(procPar.ProcName)
			inBuffer.WriteString(" output,")
			if strings.ToLower(procPar.ProcName) == "result" {
				// 存在result 出参
				isResultout = true
			}
			// 查询结果
			outBuffer.WriteString("@" + procPar.ProcName)
			outBuffer.WriteString(" as " + procPar.ProcName + ",")

			// 拼接 declare
			outParBuffer.WriteString("declare @")
			outParBuffer.WriteString(procPar.ProcName)
			outParBuffer.WriteString(" ")
			outParBuffer.WriteString(procPar.DataType)
			outParBuffer.WriteString(";")
		}
	}
	if !isResultout {
		return "", errors.New("没有传递result 出参")
	}
	// 最终sql
	result := new(bytes.Buffer)
	// 拼接 declare
	result.WriteString(outParBuffer.String())
	// 转换call sql
	var str1 = inBuffer.String()
	str1 = str1[:len(str1)-1] // 去除,
	// 拼接 call sql
	result.WriteString(str1)
	result.WriteString(";")
	// 结果查询 select sql
	var str2 = outBuffer.String()
	str2 = str2[:len(str2)-1]

	result.WriteString(str2)
	result.WriteString(";")
	return result.String(), nil
}

// 执行sqlserver 存储过程返回出参以及结果集
func execProc_ResultSetValue_Sqlserver(userID string, IsMasterDB bool, ProcName string,
	ProcParList []sysmodel.ProcPar) (resultSets []map[string]string, outValues map[string]string, err error) {
	if ProcName == "" {
		return nil, nil, errors.New("存储过程名为空")
	}
	if len(ProcParList) < 1 {
		return nil, nil, errors.New("存储过程出参至少需要包含result")
	}
	queryStr, err := concatSqlserverExeProc(ProcName, ProcParList)
	if err != nil { // 执行proc 参数不对
		return nil, nil, err
	}
	db, _ := conn.GetConnection(userID, IsMasterDB)
	dbConn := db.DB().DB
	defer dbConn.Close()
	rows, err := dbConn.Query(queryStr)
	if err != nil {
		var sqlError = concatErr(userID, err, queryStr)
		loghelper.ByError(logtype.ExecProcErr, sqlError, userID)
		return nil, nil, err
	}
	// 首先返回结果集 只允许有一个结果集返回
	for rows.Next() {
		fields, _ := rows.Columns()
		result, _ := row2mapStr(rows, fields)
		resultSets = append(resultSets, result)
	}

	// 然后才是出参
	if rows.NextResultSet() {
		for rows.Next() {
			fields, _ := rows.Columns()
			outValues, _ = row2mapStr(rows, fields)
		}
	}
	dbConn.Close()
	return resultSets, outValues, nil
}

//指定事务中执行S执行存储过程， 存储过程必须有一个名为result 的output输出参数
func execProcInTran_OutParamSqlServer(session *xorm.Session, userID string, IsMasterDB bool, ProcName string, ProcParList []sysmodel.ProcPar) (result map[string]string, err error) {
	if ProcName == "" {
		return nil, errors.New("存储过程名为空")
	}
	if len(ProcParList) < 1 {
		return nil, errors.New("存储过程出参至少需要包含result")
	}
	queryStr, err := concatSqlserverExeProc(ProcName, ProcParList)
	if err != nil { // 执行proc 参数不对
		return nil, err
	}

	res, err := session.QueryString(queryStr)
	if err != nil {
		var sqlError = concatErr(userID, err, queryStr)
		loghelper.ByError(logtype.ExecProcErr, sqlError, userID)
		return nil, err
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, errors.New("未获取到存储过程 " + ProcName + " 返回值")
}
