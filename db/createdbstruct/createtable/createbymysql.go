package createtable

import (
	"bytes"
	"strings"

	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/db/dbhelper"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/util/commutil"

	"github.com/xormplus/xorm"
)

//表单管理字段确定好以后 创建表结构
//uniqueCol 为唯一索引或联合唯一健 多个字段以逗号分隔
// 增加 IsForm 是否是业务表单类型表创建，是的话，自动增加Year 字段，存储年份，实现数据分片存储
func createTable_mysql(userID string, tableName string, pkCol string, pkIsIdentity bool, uniqueCol string,
	fields []sysmodel.SqlField, IsForm bool, dbname string) (success bool) {
	//
	if tableName == "" || len(fields) == 0 {
		// 表名为空或者没有字段
		return false
	}

	var listSql []string
	buffer := new(bytes.Buffer)
	buffer.WriteString("CREATE TABLE ")
	buffer.WriteString(tableName)
	buffer.WriteString(" (")

	var sqlColDefine = concatMysqlFields(tableName, pkCol, pkIsIdentity, fields, IsForm)
	buffer.WriteString(sqlColDefine) // 字段定义

	buffer.WriteString(")")
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
	success, _ = dbhelper.ExecMoreSql(userID, true, listSql, dbname)
	return success
}

// 拼接创建表的字段
func concatMysqlFields(TableName string, PrimaryKey string, PKIsIdentify bool, fields []sysmodel.SqlField, IsForm bool) string {
	buffer := new(bytes.Buffer)

	var colName = ""
	var sqlDataType = ""
	var isContainsFK = false
	fkBuffer := new(bytes.Buffer)
	for index, sqlCol := range fields {
		colName = strings.ReplaceAll(sqlCol.ColName, " ", "")
		sqlDataType = sqlCol.DataType.DataType
		buffer.WriteString(colName)
		buffer.WriteString(" ")
		buffer.WriteString(sysmodel.SqlDataType.GetDatatype(sqlCol.DataType))
		if sqlCol.DefaultValue != "" && sqlCol.ColName != PrimaryKey {
			// 有默认值
			buffer.WriteString(" default ")
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
		}
		// 是否主键
		if strings.ToLower(PrimaryKey) == strings.ToLower(colName) {
			// 当前是主键
			if PKIsIdentify {
				buffer.WriteString(" auto_increment primary key ")
			} else {
				buffer.WriteString(" primary key ")
			}
		}
		// 是否允许为空
		if sqlCol.IsNotNull {
			// 必填属性
			buffer.WriteString(" not null ")
		}
		if sqlCol.ColMemo != "" {
			memo := sqlCol.ColMemo
			// 有描述信息
			if strings.Contains(memo, "'") {
				memo = strings.ReplaceAll(memo, "'", "''")
			}
			buffer.WriteString(" comment '")
			buffer.WriteString(memo)
			buffer.WriteString("' ")
		}
		if sqlCol.IsFK && sqlCol.FK_TableName != "" && sqlCol.FK_TablePrimaryKey != "" {
			// 有外键
			isContainsFK = true
			fkBuffer.WriteString("CONSTRAINT FK_")
			fkBuffer.WriteString(TableName)
			fkBuffer.WriteString("_")
			fkBuffer.WriteString(colName)
			fkBuffer.WriteString(" FOREIGN KEY (")
			fkBuffer.WriteString(colName)
			fkBuffer.WriteString(") REFERENCES ")
			fkBuffer.WriteString(sqlCol.FK_TableName)
			fkBuffer.WriteString("(")
			fkBuffer.WriteString(sqlCol.FK_TablePrimaryKey)
			fkBuffer.WriteString(")")
		}
		if index < len(fields)-1 {
			buffer.WriteString(",")
		}
	}
	if isContainsFK {
		buffer.WriteString(",")
		buffer.WriteString(fkBuffer.String())
	}
	return buffer.String()
}

// 拼接定义字段的sql 非主键，无约束
func concatMysqlFieldDefine(sqlCol sysmodel.SqlField) string {
	buffer := new(bytes.Buffer)
	colName := strings.ReplaceAll(sqlCol.ColName, " ", "")
	sqlDataType := sqlCol.DataType.DataType
	buffer.WriteString(colName)
	buffer.WriteString(" ")
	buffer.WriteString(sysmodel.SqlDataType.GetDatatype(sqlCol.DataType))
	if sqlCol.DefaultValue != "" {
		// 有默认值
		buffer.WriteString(" default ")
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
	}

	// 是否允许为空
	if sqlCol.IsNotNull {
		// 必填属性
		buffer.WriteString(" not null ")
	}
	if sqlCol.ColMemo != "" {
		memo := sqlCol.ColMemo
		// 有描述信息
		if strings.Contains(memo, "'") {
			memo = strings.ReplaceAll(memo, "'", "''")
		}
		buffer.WriteString(" comment '")
		buffer.WriteString(memo)
		buffer.WriteString("' ")
	}
	return buffer.String()
}

//检查表是否存在
func tableIsExists_mysql(userID string, tableName string, dbName string) bool {
	var engine *xorm.Engine
	engine, _ = conn.GetDBConnection(userID, false, dbName)
	// 查询表是否存在
	/*tName, _ := dbhelper.QueryFirstCol(userID, false, "show tables like ? ", tableName)*/
	tName, _ := engine.QueryString("show tables like '" + tableName + "'")
	if len(tName) > 0 {
		return true
	}
	return false
}

//检查字段是否存在
func colIsExists_mysql(userID string, tableName string, colName string) bool {
	//
	var query = "select count(*) from information_schema.columns where table_name = ? and column_name = ?"
	count, _ := dbhelper.QueryFirstCol(userID, false, query, tableName, colName)
	if count != "" {
		return true
	}
	return false
}

//删除表结构
func dropTable_mysql(userID string, tableName string) bool {
	//
	success, _ := dbhelper.ExecSql(userID, true, "drop table if EXISTS ?", tableName)
	return success
}

//增加字段
func addField_mysql(userID string, tableName string, fields []sysmodel.SqlField, dbname string) bool {
	sqlList := make([]string, len(fields))
	for index, field := range fields {
		sqlList[index] = commutil.AppendStr("alter table ", tableName, " add column ", concatMysqlFieldDefine(field))
	}
	success, _ := dbhelper.ExecMoreSql(userID, true, sqlList, dbname)
	return success
	return false
}

//删除字段
func dropField_mysql(userID string, tableName string, fields []sysmodel.SqlField) bool {
	//
	sqlList := make([]string, len(fields))
	for index, field := range fields {
		sqlList[index] = commutil.AppendStr("alter table ", tableName, " drop column ", field.ColName)
	}
	success, _ := dbhelper.ExecMoreSql(userID, true, sqlList, "")
	return success
}

//修改字段 (注：只修改数据类型、MEMO说明、长度）
func alterField_mysql(userID string, tableName string, fields []sysmodel.SqlField) bool {
	//
	sqlList := make([]string, len(fields))
	for index, field := range fields {
		sqlList[index] = commutil.AppendStr("alter table ", tableName, " modify ", concatMysqlFieldDefine(field))
	}
	success, _ := dbhelper.ExecMoreSql(userID, true, sqlList, "")
	return success
	return false
}

//检查主健是否自增列
func pKIsIdentity_mysql(tableName string, pkCol string) bool {
	//
	var str = "SELECT COUNT(*) FROM `information_schema`.`columns` WHERE TABLE_NAME=? and extra='auto_increment' and COLUMN_name=?"
	value, _ := dbhelper.QueryFirstCol("admin", false, str, tableName, pkCol)
	if value != "" && value != "0" {
		return true
	}
	return false
}

//获取存储过程参数列表
func getProcParList_mysql(userID string, procNmae string) []sysmodel.ProcPar {
	//toto
	query := "select pam.parameter_mode, pam.parameter_name name,pam.data_type,pam.dtd_identifier," +
		"replace(replace(replace(pam.dtd_identifier,data_type, ''), '(',''), ')', '') length " +
		"from information_schema.parameters pam " +
		"where pam.specific_schema=DATABASE() and pam.specific_name=? order by ordinal_position"
	procPars, _ := dbhelper.Query(userID, false, query, procNmae)
	// 拼接返回值
	// 遍历赋值
	if procPars != nil && len(procPars) > 0 {
		ListPar := make([]sysmodel.ProcPar, len(procPars))
		var isOut = false
		for index, par := range procPars {
			if "OUT" == par["parameter_mode"] {
				// 是否是出参
				isOut = true
			}
			ListPar[index] = sysmodel.ProcPar{
				ProcName:   par["name"],
				DataType:   par["dtd_identifier"],
				DataLength: commutil.ToInt(par["length"]),
				IsOutPut:   isOut,
			}
		}
		return ListPar
	}

	return nil
}

//获取某张表的结构信息
func getSqlTableStruct_mysql(userID string, sqlTableName string, dbName string) (m sysmodel.SqlTableStruct) {
	var engine *xorm.Engine
	engine, _ = conn.GetDBConnection(userID, false, dbName)
	//toto
	values, _ := engine.QueryInterface("select * from information_schema.columns where table_schema=DATABASE() and table_name=? order by ordinal_position", sqlTableName)
	/*values, _ := dbhelper.Query(userID, false,
	"select * from information_schema.columns where table_schema=DATABASE() and table_name=? order by ordinal_position",
	sqlTableName)*/
	m.TableName = sqlTableName
	if values != nil && len(values) > 0 {
		m.ColList = make([]sysmodel.SqlField, len(values))
		var sqlField sysmodel.SqlField
		for index, col := range values {
			sqlField = sysmodel.SqlField{}
			if col["COLUMN_KEY"] == "PRI" { // 主键约束
				m.PrimaryKey = commutil.ToString(col["COLUMN_NAME"])
				sqlField.IsPrimaryKey = true
			}
			if col["CHARACTER_MAXIMUM_LENGTH"] == "" { // 字段长度
				sqlField.DataLength = 0
			} else {
				sqlField.DataLength = commutil.ToInt(col["CHARACTER_MAXIMUM_LENGTH"])
			}
			sqlField.DataType = sysmodel.SqlDataType{DataType: commutil.ToString(col["DATA_TYPE"]), Length: sqlField.DataLength}
			if col["NUMERIC_SCALE"] == "" { // 字段长度
				sqlField.DataType.DecimalLength = 0
			} else {
				sqlField.DataType.DecimalLength = commutil.ToInt(col["NUMERIC_SCALE"])
			}
			sqlField.ColName = commutil.ToString(col["COLUMN_NAME"])
			sqlField.DefaultValue = commutil.ToString(col["COLUMN_DEFAULT"]) //默认值
			if col["IS_NULLABLE"] == "NO" {
				sqlField.IsNotNull = true
			} else {
				sqlField.IsNotNull = false
			}
			sqlField.ColMemo = commutil.ToString(col["COLUMN_COMMENT"])
			m.ColList[index] = sqlField
		}
		return m
	}

	return m
}

//返回某字段的结构信息 例：字段名，数据类型，长度等
func getColStruct_mysql(userID string, sqlTableName string, colName string) sysmodel.SqlField {
	var f sysmodel.SqlField
	values, _ := dbhelper.Query(userID, false,
		"select * from information_schema.columns where table_schema=DATABASE() and table_name=? and column_name=? order by ordinal_position",
		sqlTableName, colName)
	if values != nil && len(values) > 0 {
		var sqlField sysmodel.SqlField
		sqlField = sysmodel.SqlField{}
		col := values[0]
		if col["COLUMN_KEY"] == "PRI" { // 主键约束
			sqlField.IsPrimaryKey = true
		}
		if col["CHARACTER_MAXIMUM_LENGTH"] == "" { // 字段长度
			sqlField.DataLength = 0
		} else {
			sqlField.DataLength = commutil.ToInt(col["CHARACTER_MAXIMUM_LENGTH"])
		}
		sqlField.DataType = sysmodel.SqlDataType{DataType: col["DATA_TYPE"], Length: sqlField.DataLength}
		sqlField.ColName = col["COLUMN_NAME"]
		sqlField.DefaultValue = col["COLUMN_DEFAULT"] //默认值
		if col["IS_NULLABLE"] == "NO" {
			sqlField.IsNotNull = true
		} else {
			sqlField.IsNotNull = false
		}
		sqlField.ColMemo = col["COLUMN_COMMENT"]
		return sqlField
	}
	//todo
	return f
}
