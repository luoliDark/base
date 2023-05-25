// jsz by 2020.2.7 用于创建数据表或修改表结构信息

package createtable

import (
	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/sysmodel/logtype"
)

//表单管理字段确定好以后 创建表结构
//uniqueCol 为唯一索引或联合唯一健 多个字段以逗号分隔
func CreateTable(userID string, tableName string, pkCol string, pkIsIdentity bool, uniqueCol string, fields []sysmodel.SqlField, IsForm bool, dbname string) bool {

	switch confighelper.GetCurrdb() {
	case conn.DBTYPE_MYSQL:
		return createTable_mysql(userID, tableName, pkCol, pkIsIdentity, uniqueCol, fields, IsForm, dbname)
	case conn.DBTYPE_SQLSERVER:
		return createTable_sqlserver(userID, tableName, pkCol, pkIsIdentity, uniqueCol, fields, IsForm)
	default:
		loghelper.ByHighError(logtype.CreateTableErr, "currdb配置错误", userID)
		panic("建表错误currdb配置错误")
	}
}

//检查表是否存在
func TableIsExists(userID string, tableName string, dbName string) bool {
	switch confighelper.GetCurrdb() {
	case conn.DBTYPE_MYSQL:
		return tableIsExists_mysql(userID, tableName, dbName)
	case conn.DBTYPE_SQLSERVER:
		return tableIsExists_sqlserver(userID, tableName)
	default:
		loghelper.ByHighError(logtype.CreateTableErr, "currdb配置错误", userID)
		panic("建表错误currdb配置错误")
	}
}

//检查字段是否存在
func ColIsExists(userID string, tableName string, colName string) bool {
	switch confighelper.GetCurrdb() {
	case conn.DBTYPE_MYSQL:
		return colIsExists_mysql(userID, tableName, colName)
	case conn.DBTYPE_SQLSERVER:
		return colIsExists_sqlserver(userID, tableName, colName)
	default:
		loghelper.ByHighError(logtype.CreateTableErr, "currdb配置错误", userID)
		panic("建表错误currdb配置错误")
	}
}

//删除表结构
func DropTable(userID string, tableName string) bool {
	switch confighelper.GetCurrdb() {
	case conn.DBTYPE_MYSQL:
		return dropTable_mysql(userID, tableName)
	case conn.DBTYPE_SQLSERVER:
		return dropTable_sqlserver(userID, tableName)
	default:
		loghelper.ByHighError(logtype.CreateTableErr, "currdb配置错误", userID)
		panic("建表错误currdb配置错误")
	}
}

//增加字段
func AddField(userID string, tableName string, fields []sysmodel.SqlField, dbname string) bool {
	switch confighelper.GetCurrdb() {
	case conn.DBTYPE_MYSQL:
		return addField_mysql(userID, tableName, fields, dbname)
	case conn.DBTYPE_SQLSERVER:
		return addField_sqlserver(userID, tableName, fields, dbname)
	default:
		loghelper.ByHighError(logtype.CreateTableErr, "currdb配置错误", userID)
		panic("建表错误currdb配置错误")
	}
}

//删除字段
func DropField(userID string, tableName string, fields []sysmodel.SqlField) bool {
	switch confighelper.GetCurrdb() {
	case conn.DBTYPE_MYSQL:
		return dropField_mysql(userID, tableName, fields)
	case conn.DBTYPE_SQLSERVER:
		return dropField_sqlserver(userID, tableName, fields)
	default:
		loghelper.ByHighError(logtype.CreateTableErr, "currdb配置错误", userID)
		panic("建表错误currdb配置错误")
	}
}

//修改字段 (注：只修改数据类型、MEMO说明、长度）
func AlterField(userID string, tableName string, fields []sysmodel.SqlField) bool {
	switch confighelper.GetCurrdb() {
	case conn.DBTYPE_MYSQL:
		return alterField_mysql(userID, tableName, fields)
	case conn.DBTYPE_SQLSERVER:
		return alterField_sqlserver(userID, tableName, fields)
	default:
		loghelper.ByHighError(logtype.CreateTableErr, "currdb配置错误", userID)
		panic("建表错误currdb配置错误")
	}
}

//检查主健是否自增列
func PKIsIdentity(tableName string, pkCol string) bool {
	switch confighelper.GetCurrdb() {
	case conn.DBTYPE_MYSQL:
		return pKIsIdentity_mysql(tableName, pkCol)
	case conn.DBTYPE_SQLSERVER:
		return pKIsIdentity_sqlserver(tableName, pkCol)
	default:
		loghelper.ByHighError(logtype.CreateTableErr, "currdb配置错误", "")
		panic("建表错误currdb配置错误")
	}
}

//获取存储过程参数列表
func GetProcParList(userID string, procNmae string) []sysmodel.ProcPar {
	switch confighelper.GetCurrdb() {
	case conn.DBTYPE_MYSQL:
		return getProcParList_mysql(userID, procNmae)
	case conn.DBTYPE_SQLSERVER:
		return getProcParList_sqlserver(userID, procNmae)
	default:
		loghelper.ByHighError(logtype.CreateTableErr, "currdb配置错误", userID)
		panic("建表错误currdb配置错误")
	}
}

//获取某张表的结构信息
func GetSqlTableStruct(userID string, sqlTableName string, dbName string) sysmodel.SqlTableStruct {
	switch confighelper.GetCurrdb() {
	case conn.DBTYPE_MYSQL:
		return getSqlTableStruct_mysql(userID, sqlTableName, dbName)
	case conn.DBTYPE_SQLSERVER:
		return getSqlTableStruct_sqlserver(userID, sqlTableName)
	default:
		loghelper.ByHighError(logtype.CreateTableErr, "currdb配置错误", userID)
		panic("建表错误currdb配置错误")
	}
}

//返回某字段的结构信息 例：字段名，数据类型，长度等
func GetColStruct(userID string, sqlTableName string, colName string) sysmodel.SqlField {
	switch confighelper.GetCurrdb() {
	case conn.DBTYPE_MYSQL:
		return getColStruct_mysql(userID, sqlTableName, colName)
	case conn.DBTYPE_SQLSERVER:
		return getColStruct_sqlserver(userID, sqlTableName, colName)
	default:
		loghelper.ByHighError(logtype.CreateTableErr, "currdb配置错误", userID)
		panic("建表错误currdb配置错误")
	}
}
