package createtable

import (
	"base/base/db/dbhelper"
	"base/base/sysmodel"
	"base/base/util/commutil"
	"bytes"
	"strings"
)

//表单管理字段确定好以后 创建表结构
//uniqueCol 为唯一索引或联合唯一健 多个字段以逗号分隔
// 增加 IsForm 是否是业务表单类型表创建，是的话，自动增加Year 字段，存储年份，实现数据分片存储
func createTable_sqlserver(userID string, tableName string, pkCol string, pkIsIdentity bool, uniqueCol string,
	fields []sysmodel.SqlField, IsForm bool) (success bool) {
	//
	if tableName == "" || len(fields) == 0 {
		// 表名为空或者没有字段
		return false
	}

	var listSql []string
	buffer := new(bytes.Buffer)
	buffer.WriteString("CREATE TABLE ")
	buffer.WriteString(tableName)

	var sqlColDefine = concatSqlServerFields(tableName, pkCol, pkIsIdentity, fields, IsForm)
	buffer.WriteString(sqlColDefine) // 字段定义

	listSql = append(listSql, buffer.String())
	// 是否有唯一键
	if strings.ReplaceAll(uniqueCol, " ", "") != "" {
		// 包含唯一键
		idx_buffer := new(bytes.Buffer)
		idx_buffer.WriteString("create unique index ")
		idx_buffer.WriteString(strings.ToLower(tableName))
		idx_buffer.WriteString("_")
		idx_buffer.WriteString(strings.ToLower(uniqueCol))
		idx_buffer.WriteString(" on ")
		idx_buffer.WriteString(tableName)
		idx_buffer.WriteString(" (")
		idx_buffer.WriteString(uniqueCol)
		idx_buffer.WriteString(")")
		listSql = append(listSql, idx_buffer.String())
	}
	// 执行批量sql
	success, _ = dbhelper.ExecMoreSql(userID, true, listSql)
	return success
}

// 拼接创建表的字段
func concatSqlServerFields(TableName string, PrimaryKey string, PKIsIdentify bool, fields []sysmodel.SqlField, IsForm bool) string {
	buffer := new(bytes.Buffer)

	var colName = ""
	var sqlDataType = ""
	buffer.WriteString("(")
	memoBuffer := new(bytes.Buffer)

	for index, sqlCol := range fields {
		colName = strings.ReplaceAll(sqlCol.ColName, " ", "")
		sqlDataType = sqlCol.DataType.DataType
		buffer.WriteString(colName)
		buffer.WriteString(" ")
		buffer.WriteString(sysmodel.SqlDataType.GetDatatype(sqlCol.DataType))
		if sqlCol.DefaultValue != "" {
			// 有默认值
			buffer.WriteString(" default (")
			if strings.ToLower(sqlDataType) == "varchar" || strings.ToLower(sqlDataType) == "char" {
				// 字符型或字符串类型
				buffer.WriteString("'")
				if strings.Contains(sqlCol.DefaultValue, "'") {
					buffer.WriteString(strings.ReplaceAll(sqlCol.DefaultValue, "'", "''"))
				} else {
					buffer.WriteString(sqlCol.DefaultValue)
				}
				buffer.WriteString("' ")
			} else {
				buffer.WriteString(sqlCol.DefaultValue)
			}
			buffer.WriteString(") ")
		}
		// 是否主键
		if strings.ToLower(PrimaryKey) == strings.ToLower(colName) {
			// 当前是主键
			if PKIsIdentify {
				buffer.WriteString(" identity primary key ")
			} else {
				buffer.WriteString(" primary key ")
			}
		}
		// 是否允许为空
		if sqlCol.IsNotNull {
			// 必填属性
			buffer.WriteString(" not null ")
		}
		// sqlserver 版本添加描述有区别需要单独调用 proc 来实现
		if sqlCol.IsFK && sqlCol.FK_TableName != "" && sqlCol.FK_TablePrimaryKey != "" {
			// 有外键
			buffer.WriteString(" references ")
			buffer.WriteString(sqlCol.FK_TableName)
			buffer.WriteString("(")
			buffer.WriteString(sqlCol.FK_TablePrimaryKey)
			buffer.WriteString(")")

		}
		if sqlCol.ColMemo != "" {
			memo := sqlCol.ColMemo
			// 有描述信息
			if strings.Contains(memo, "'") {
				memo = strings.ReplaceAll(memo, "'", "''")
			}

			memoBuffer.WriteString("exec sys.sp_addextendedproperty @name=N'MS_Description',@value='" + memo + "',")
			memoBuffer.WriteString("@level0type=N'SCHEMA',@level0name=N'dbo',@level1type=N'TABLE',@level1name='" + TableName + "',")
			memoBuffer.WriteString("@level2type=N'COLUMN',@level2name='" + colName + "';")

		}
		if index < len(fields)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(");")
	if memoBuffer.Cap() > 0 {
		buffer.WriteString(memoBuffer.String())
	}
	// 添加描述
	return buffer.String()
}

// 拼接定义字段的sql 非主键，无约束
func concatSqlServerFieldDefine(sqlCol sysmodel.SqlField) string {
	buffer := new(bytes.Buffer)
	colName := strings.ReplaceAll(sqlCol.ColName, " ", "")
	sqlDataType := sqlCol.DataType.DataType
	buffer.WriteString(colName)
	buffer.WriteString(" ")
	buffer.WriteString(sysmodel.SqlDataType.GetDatatype(sqlCol.DataType))
	if sqlCol.DefaultValue != "" {
		// 有默认值
		buffer.WriteString(" default (")
		if strings.ToLower(sqlDataType) == "varchar" || strings.ToLower(sqlDataType) == "char" {
			// 字符型或字符串类型
			buffer.WriteString("'")
			if strings.Contains(sqlCol.DefaultValue, "'") {
				buffer.WriteString(strings.ReplaceAll(sqlCol.DefaultValue, "'", "''"))
			} else {
				buffer.WriteString(sqlCol.DefaultValue)
			}
			buffer.WriteString("' ")
		} else {
			buffer.WriteString(sqlCol.DefaultValue)
		}
		buffer.WriteString(") ")
	}

	// 是否允许为空
	if sqlCol.IsNotNull {
		// 必填属性
		buffer.WriteString(" not null ")
	}
	return buffer.String()
}

//检查表是否存在
func tableIsExists_sqlserver(userID string, tableName string) bool {
	// 查询表是否存在
	tName, _ := dbhelper.QueryFirstCol(userID, false, "show tables like ? ", tableName)
	if len(tName) > 0 {
		return true
	}
	return false
}

//检查字段是否存在
func colIsExists_sqlserver(userID string, tableName string, colName string) bool {
	//
	var query = "select count(*) from information_schema.columns where table_name = ? and column_name = ?"
	count, _ := dbhelper.QueryFirstCol(userID, false, query, tableName, colName)
	if count != "" {
		return true
	}
	return false
}

//删除表结构
func dropTable_sqlserver(userID string, tableName string) bool {
	//
	success, _ := dbhelper.ExecSql(userID, true, "drop table if EXISTS ?", tableName)
	return success
}

//增加字段
func addField_sqlserver(userID string, tableName string, fields []sysmodel.SqlField) bool {
	sqlList := make([]string, len(fields))
	for index, field := range fields {

		var sql = commutil.AppendStr("alter table ", tableName, " add ", concatSqlServerFieldDefine(field), ";")
		if field.ColMemo != "" {
			memo := field.ColMemo
			// 有描述信息
			if strings.Contains(memo, "'") {
				memo = strings.ReplaceAll(memo, "'", "''")
			}
			memoBuffer := new(bytes.Buffer)
			memoBuffer.WriteString("exec sys.sp_addextendedproperty @name=N'MS_Description',@value='" + memo + "',")
			memoBuffer.WriteString("@level0type=N'SCHEMA',@level0name=N'dbo',@level1type=N'TABLE',@level1name='" + tableName + "',")
			memoBuffer.WriteString("@level2type=N'COLUMN',@level2name='" + field.ColName + "';")
			sql = commutil.AppendStr(sql, memoBuffer.String())
		}
		sqlList[index] = sql
	}
	success, _ := dbhelper.ExecMoreSql(userID, true, sqlList)
	return success
	return false
}

//删除字段
func dropField_sqlserver(userID string, tableName string, fields []sysmodel.SqlField) bool {
	//
	sqlList := make([]string, len(fields))
	for index, field := range fields {
		sqlList[index] = commutil.AppendStr("alter table ", tableName, " drop column ", field.ColName)
	}
	success, _ := dbhelper.ExecMoreSql(userID, true, sqlList)
	return success
}

//修改字段 (注：只修改数据类型、MEMO说明、长度）
func alterField_sqlserver(userID string, tableName string, fields []sysmodel.SqlField) bool {
	//
	sqlList := make([]string, len(fields))
	for index, field := range fields {
		sqlList[index] = commutil.AppendStr("alter table ", tableName, " alter column ", concatSqlServerFieldDefine(field))
	}
	success, _ := dbhelper.ExecMoreSql(userID, true, sqlList)
	return success
	return false
}

//检查主健是否自增列
func pKIsIdentity_sqlserver(tableName string, pkCol string) bool {
	//
	var str = "if Exists(Select top 1 1 from sysobjects Where objectproperty(id, 'TableHasIdentity') = 1 and upper(name) = upper(?)) select 1 else select 0"
	value, _ := dbhelper.QueryFirstCol("admin", false, str, tableName, pkCol)
	if value == "1" {
		return true
	}
	return false
}

//获取存储过程参数列表
func getProcParList_sqlserver(userID string, procNmae string) []sysmodel.ProcPar {
	//toto
	sql := "select s.name as parname,s.isoutparam,s.length,t.name typename  from syscolumns s left join systypes t on t.xtype=s.xtype" +
		" where s.id =(select id from sysobjects where name=?) and t.name!='sysname' "
	procPars, _ := dbhelper.Query(userID, false, sql, procNmae)
	// 拼接返回值
	// 遍历赋值
	if procPars != nil && len(procPars) > 0 {
		ListPar := make([]sysmodel.ProcPar, len(procPars))
		var isOut = false
		for index, par := range procPars {
			if "1" == par["isoutparam"] {
				// 是否是出参
				isOut = true
			}
			ListPar[index] = sysmodel.ProcPar{
				ProcName:   par["parname"],
				DataType:   par["typename"],
				DataLength: commutil.ToInt(par["length"]),
				IsOutPut:   isOut,
			}
		}
		return ListPar
	}

	return nil
}

//获取某张表的结构信息
func getSqlTableStruct_sqlserver(userID string, sqlTableName string) (m sysmodel.SqlTableStruct) {
	//toto
	str := "SELECT col.name AS column_name," +
		" t.name AS datatype," +
		" CASE WHEN col.length=4 then null else col.length end AS length," +
		" ISNULL(COLUMNPROPERTY(col.id, col.name, 'Scale'), 0) AS decimals, " +
		" CASE WHEN COLUMNPROPERTY(col.id, col.name, 'IsIdentity') = 1 THEN '1' ELSE '0' END AS isidentity ," +
		" CASE WHEN EXISTS ( SELECT 1  FROM dbo.sysindexes si INNER JOIN dbo.sysindexkeys sik ON si.id = sik.id AND si.indid = sik.indid" +
		" INNER JOIN dbo.syscolumns sc ON sc.id = sik.id" +
		" AND sc.colid = sik.colid" +
		" INNER JOIN dbo.sysobjects so ON so.name = si.name" +
		" AND so.xtype = 'PK'" +
		" WHERE    sc.id = col.id" +
		" AND sc.colid = col.colid ) THEN '1'" +
		" ELSE '0' END AS primarykey," +
		" CASE WHEN col.isnullable = 1 THEN 0 ELSE '1' END AS isnotnull," +
		" ISNULL(comm.text, '') AS defaultvalue " +
		" FROM dbo.syscolumns col LEFT  JOIN dbo.systypes t ON col.xtype = t.xusertype" +
		" inner JOIN dbo.sysobjects obj ON col.id = obj.id" +
		" AND obj.xtype = 'U'" +
		" AND obj.status >= 0" +
		" LEFT  JOIN dbo.syscomments comm ON col.cdefault = comm.id" +
		" LEFT  JOIN sys.extended_properties ep ON col.id = ep.major_id" +
		" AND col.colid = ep.minor_id" +
		" AND ep.name = 'MS_Description' " +
		" LEFT  JOIN sys.extended_properties epTwo ON obj.id = epTwo.major_id" +
		" AND epTwo.minor_id = 0" +
		" AND epTwo.name = 'MS_Description'" +
		" WHERE obj.name =? ORDER BY col.colorder"
	values, _ := dbhelper.Query(userID, false,
		str,
		sqlTableName)
	m.TableName = sqlTableName
	if values != nil && len(values) > 0 {
		m.ColList = make([]sysmodel.SqlField, len(values))
		var sqlField sysmodel.SqlField
		for index, col := range values {
			sqlField = sysmodel.SqlField{}
			if col["primarykey"] == "1" { // 主键约束
				m.PrimaryKey = col["column_name"]
				sqlField.IsPrimaryKey = true
			}
			if col["length"] == "" { // 字段长度
				sqlField.DataLength = 0
			} else {
				sqlField.DataLength = commutil.ToInt(col["length"])
			}
			sqlField.DataType = sysmodel.SqlDataType{DataType: col["datatype"], Length: sqlField.DataLength}
			sqlField.ColName = col["column_name"]
			sqlField.DefaultValue = col["defaultvalue"] //默认值
			if col["isnotnull"] == "1" {
				sqlField.IsNotNull = true
			} else {
				sqlField.IsNotNull = false
			}
			m.ColList[index] = sqlField
		}
		return m
	}

	return m
}

//返回某字段的结构信息 例：字段名，数据类型，长度等
func getColStruct_sqlserver(userID string, sqlTableName string, colName string) sysmodel.SqlField {
	var f sysmodel.SqlField
	str := "SELECT col.name AS column_name," +
		" t.name AS datatype," +
		" CASE WHEN col.length=4 then null else col.length end AS length," +
		" ISNULL(COLUMNPROPERTY(col.id, col.name, 'Scale'), 0) AS decimals, " +
		" CASE WHEN COLUMNPROPERTY(col.id, col.name, 'IsIdentity') = 1 THEN '1' ELSE '0' END AS isidentity ," +
		" CASE WHEN EXISTS ( SELECT 1  FROM dbo.sysindexes si INNER JOIN dbo.sysindexkeys sik ON si.id = sik.id AND si.indid = sik.indid" +
		" INNER JOIN dbo.syscolumns sc ON sc.id = sik.id" +
		" AND sc.colid = sik.colid" +
		" INNER JOIN dbo.sysobjects so ON so.name = si.name" +
		" AND so.xtype = 'PK'" +
		" WHERE    sc.id = col.id" +
		" AND sc.colid = col.colid ) THEN '1'" +
		" ELSE '0' END AS primarykey," +
		" CASE WHEN col.isnullable = 1 THEN 0 ELSE '1' END AS isnotnull," +
		" ISNULL(comm.text, '') AS defaultvalue " +
		" FROM dbo.syscolumns col LEFT  JOIN dbo.systypes t ON col.xtype = t.xusertype" +
		" inner JOIN dbo.sysobjects obj ON col.id = obj.id" +
		" AND obj.xtype = 'U'" +
		" AND obj.status >= 0" +
		" LEFT  JOIN dbo.syscomments comm ON col.cdefault = comm.id" +
		" LEFT  JOIN sys.extended_properties ep ON col.id = ep.major_id" +
		" AND col.colid = ep.minor_id" +
		" AND ep.name = 'MS_Description' " +
		" LEFT  JOIN sys.extended_properties epTwo ON obj.id = epTwo.major_id" +
		" AND epTwo.minor_id = 0" +
		" AND epTwo.name = 'MS_Description'" +
		" WHERE obj.name =? and col.name=? ORDER BY col.colorder"
	values, _ := dbhelper.Query(userID, false,
		str,
		sqlTableName, colName)
	if values != nil && len(values) > 0 {
		var sqlField sysmodel.SqlField
		sqlField = sysmodel.SqlField{}
		col := values[0]
		if col["primarykey"] == "1" { // 主键约束
			sqlField.IsPrimaryKey = true
		}
		if col["length"] == "" { // 字段长度
			sqlField.DataLength = 0
		} else {
			sqlField.DataLength = commutil.ToInt(col["length"])
		}
		sqlField.DataType = sysmodel.SqlDataType{DataType: col["datatype"], Length: sqlField.DataLength}
		sqlField.ColName = col["column_name"]
		sqlField.DefaultValue = col["defaultvalue"] //默认值
		if col["isnotnull"] == "1" {
			sqlField.IsNotNull = true
		} else {
			sqlField.IsNotNull = false
		}
		return sqlField
	}
	//todo
	return f
}
