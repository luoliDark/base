package conn

import (
	"bytes"
	"fmt"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/sysmodel/logtype"
	"github.com/luoliDark/base/util/commutil"
	_ "github.com/luoliDark/base/util/commutil"
	"github.com/xormplus/xorm"
)

const CONF_SEC_GLOBAL = "global"
const CONF_CURRDB = "currdb"
const CONF_DBSERVER = "server"
const CONF_DBUSERNAME = "username"
const CONF_DBPASSWORD = "password"
const CONF_DBPORT = "port"
const CONF_DBNAME = "database"
const DBTYPE_MYSQL = "mysql"
const DBTYPE_OCR = "ocrdb"
const DBTYPE_SQLSERVER = "mssql"

//采用单例
var EngineMap map[string]*xorm.Engine
var DBType string

func init() {
	// 初始化数据库连接对象
	// 初始化连接过程，需要把所有的数据库连接都初始化好，如slave 和 master 然后存储在map里面
	InitConnections()

}

func InitConnections() {
	if EngineMap == nil {
		EngineMap = make(map[string]*xorm.Engine)
	}

	// 读取配置文件
	// 当前数据库类型
	DBType = confighelper.GetIniConfig(CONF_SEC_GLOBAL, CONF_CURRDB)
	if DBType == "" {
		DBType = DBTYPE_MYSQL
	}

	//初始化主库连接
	initConn(DBType, DBType)

	//初始化营收稽核 外部原始数据库
	//initConn(DBType, "saleexdb")

	//初始化营收稽核 按天汇总的 数据库
	//initConn(DBType, "salesumdb")

	//初始化 商旅平台 数据库
	//initConn(DBType, "easytradb")

	initConn(DBType, "busfadb")
}

//初始化连接
func initConn(DBType string, dbName string) {

	server := confighelper.GetIniConfig(dbName, CONF_DBSERVER)
	port := confighelper.GetIniConfig(dbName, CONF_DBPORT)
	username := confighelper.GetIniConfig(dbName, CONF_DBUSERNAME)
	password := confighelper.GetIniConfig(dbName, CONF_DBPASSWORD)
	database := confighelper.GetIniConfig(dbName, CONF_DBNAME)

	if server == "" || username == "" || password == "" || database == "" {
		fmt.Println("连接数据库" + dbName + "失败，服务器地址或账号密码未设置")
		return
	}

	var currDbUrl string
	isdev := commutil.ToBool(confighelper.GetIniConfig("global", "isdeving"))

	switch strings.ToLower(DBType) {
	case DBTYPE_MYSQL:
		currDbUrl = concatMysql(server, port, username, password, database)
		db, err := xorm.NewEngine(DBTYPE_MYSQL, currDbUrl)
		if err != nil {
			loghelper.ByHighError(logtype.GetConnErr, err.Error(), "")
		}

		if isdev {
			//开发模式
			db.ShowSQL(true)
		}

		//连接池设置
		db.SetMaxIdleConns(30)
		db.SetMaxOpenConns(200)

		EngineMap[dbName] = db

	case DBTYPE_SQLSERVER:
		currDbUrl = concatSqlServer(server, port, username, password, database)
		db, err := xorm.NewEngine(DBType, currDbUrl)
		if err != nil {
			loghelper.ByHighError(logtype.GetConnErr, err.Error(), "")
		}
		//todo
		if isdev {
			//开发模式
			db.ShowSQL(true)
		}
		//连接池设置
		db.SetMaxIdleConns(30)
		db.SetMaxOpenConns(200)
		EngineMap[dbName] = db
	default:
		loghelper.ByHighError(logtype.GetConnErr, "请检查config.ini中currdb 配置", "")
		panic("获取连接对象失败,请检查config.ini中currdb 配置")
	}
}

//初始化连接
//func initConn(DBType string) {
//	section := DBType
//	server := confighelper.GetIniConfig(section, CONF_DBSERVER)
//	port := confighelper.GetIniConfig(section, CONF_DBPORT)
//	username := confighelper.GetIniConfig(section, CONF_DBUSERNAME)
//	password := confighelper.GetIniConfig(section, CONF_DBPASSWORD)
//	database := confighelper.GetIniConfig(section, CONF_DBNAME)
//
//	if server == "" || username == "" || password == "" || database == "" {
//		fmt.Println("连接数据库" + section + "失败，服务器地址或账号密码未设置")
//		return
//	}
//
//	var currDbUrl string
//	isdev := commutil.ToBool(confighelper.GetIniConfig("global", "isdeving"))
//
//	switch strings.ToLower(DBType) {
//	case DBTYPE_MYSQL, DBTYPE_OCR:
//		currDbUrl = concatMysql(server, port, username, password, database)
//		db, err := xorm.NewEngine(DBTYPE_MYSQL, currDbUrl)
//		if err != nil {
//			loghelper.ByHighError(logtype.GetConnErr, err.Error(), "")
//		}
//
//		if isdev {
//			//开发模式
//			db.ShowSQL(true)
//		}
//
//		//连接池设置
//		db.SetMaxIdleConns(30)
//		db.SetMaxOpenConns(200)
//
//		if DBType == DBTYPE_MYSQL {
//			EngineMap[DBTYPE_MYSQL] = db
//		} else {
//			EngineMap[DBTYPE_OCR] = db
//		}
//	case DBTYPE_SQLSERVER:
//		currDbUrl = concatSqlServer(server, port, username, password, database)
//		db, err := xorm.NewEngine(DBType, currDbUrl)
//		if err != nil {
//			loghelper.ByHighError(logtype.GetConnErr, err.Error(), "")
//		}
//		//todo
//		if isdev {
//			//开发模式
//			db.ShowSQL(true)
//		}
//		//连接池设置
//		db.SetMaxIdleConns(30)
//		db.SetMaxOpenConns(200)
//		EngineMap[DBTYPE_SQLSERVER] = db
//	default:
//		loghelper.ByHighError(logtype.GetConnErr, "请检查config.ini中currdb 配置", "")
//		panic("获取连接对象失败,请检查config.ini中currdb 配置")
//	}
//}

// 拼接连接mysql 字符串
// server  ip地址
// port   端口，为空则使用默认
// username 用户账号
// password  密码
// database 数据库名称
func concatMysql(server string, port string, username string, password string, database string) string {
	buffer := new(bytes.Buffer)
	buffer.WriteString(username)
	buffer.WriteString(":")
	buffer.WriteString(password)
	buffer.WriteString("@tcp(")
	buffer.WriteString(server)
	if port != "" {
		buffer.WriteString(":" + port)
	}
	buffer.WriteString(")/")
	buffer.WriteString(database)
	buffer.WriteString("?charset=utf8&multiStatements=true")
	return buffer.String()
}

// 拼接 sqlserver 连接字符串
//参考格式 ：sqlserver://username:password@host:port?param1=value&param2=value
//sqlserver://sa:mypass@localhost:1234?database=数据库名&connection+timeout=30 //本地主机上的端口1234
func concatSqlServer(server string, port string, username string, password string, database string) string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("sqlserver://")
	buffer.WriteString(username)
	buffer.WriteString(":")
	buffer.WriteString(password)
	buffer.WriteString("@")
	buffer.WriteString(server)
	if port != "" {
		buffer.WriteString(":" + port)
	}
	buffer.WriteString("?database=")
	buffer.WriteString(database)
	buffer.WriteString("&charset=utf8&multiStatements=true")
	return buffer.String()
}

// GetConnStr  获取当前连接字符串
func GetConnStr() string {

	// 读取配置文件
	DBType = confighelper.GetIniConfig(CONF_SEC_GLOBAL, CONF_CURRDB)

	// 当前数据库类型
	//DBType, _ = conf.GetValue(CONF_SEC_GLOBAL, CONF_CURRDB)
	if DBType == "" {
		DBType = DBTYPE_MYSQL
	}
	section := DBType
	server := confighelper.GetIniConfig(section, CONF_DBSERVER)
	port := confighelper.GetIniConfig(section, CONF_DBPORT)
	username := confighelper.GetIniConfig(section, CONF_DBUSERNAME)
	password := confighelper.GetIniConfig(section, CONF_DBPASSWORD)
	database := confighelper.GetIniConfig(section, CONF_DBNAME)

	if server == "" || username == "" || password == "" || database == "" {
		fmt.Println("连接数据库" + section + "失败，服务器地址或账号密码未设置")
		return ""
	}

	var currDbUrl string

	switch strings.ToLower(DBType) {
	case DBTYPE_MYSQL:
		currDbUrl = concatMysql(server, port, username, password, database)

	case DBTYPE_SQLSERVER:
		currDbUrl = concatSqlServer(server, port, username, password, database)

	default:
		loghelper.ByHighError(logtype.GetConnErr, "请检查config.ini中currdb 配置", "")
		panic("获取连接对象失败,请检查config.ini中currdb 配置")
	}

	return currDbUrl
}

// GetConnection 获取连接对象
func GetDB() (db *xorm.Engine, err error) {

	//初始化对象
	if EngineMap == nil || EngineMap[DBType] == nil {
		// 连接串是空的  重新执行数据库连接获取
		InitConnections()
	}
	db = EngineMap[DBType]
	//??? 缺少连接监控程序，防止开发人员连接获取出去以后不进行关闭，GC或mysql 也是长时间不地其关闭
	//注：连接监控时需要获取当前调用本方法的上层go 代码方法，这样才能知道是什么代码没关闭连接

	return db, err
}

// GetDBConnection 获取连接对象
func GetSession(ismasterdb bool) (db *xorm.Session) {
	eng, err := GetConnection("", ismasterdb)
	if err != nil {
		panic(err)
	} else {
		return eng.NewSession()
	}
}

// GetDBConnection 获取连接对象
func GetEasyTraDb() (db *xorm.Engine, err error) {

	//初始化对象
	if EngineMap == nil || EngineMap["easytradb"] == nil {
		// 连接串是空的  重新执行数据库连接获取
		InitConnections()
	}
	db = EngineMap["easytradb"]
	//??? 缺少连接监控程序，防止开发人员连接获取出去以后不进行关闭，GC或mysql 也是长时间不地其关闭
	//注：连接监控时需要获取当前调用本方法的上层go 代码方法，这样才能知道是什么代码没关闭连接

	return db, err
}

// GetDBConnection 获取连接对象
func GetOcrDB() (db *xorm.Engine, err error) {

	//初始化对象
	if EngineMap == nil || EngineMap["ocrdb"] == nil {
		// 连接串是空的  重新执行数据库连接获取
		InitConnections()
	}
	db = EngineMap["ocrdb"]
	//??? 缺少连接监控程序，防止开发人员连接获取出去以后不进行关闭，GC或mysql 也是长时间不地其关闭
	//注：连接监控时需要获取当前调用本方法的上层go 代码方法，这样才能知道是什么代码没关闭连接

	return db, err
}

// 获取数据库连接对象，主从方式获取，根据传递的userid 确定从数据库连接
// Engine调用一个SQL都会创建新的Session，完毕又会关闭Session， 多个SQL语句场景，非常消耗性能
// 多个SQL语句场景，建议使用 GetSession 获取Session ，注意需要手动关闭 建议使用该方法关闭：defer session.Close()
func GetConnection(userid string, ismasterdb bool) (db *xorm.Engine, err error) {
	if ismasterdb {
		// 主数据库连接
		return GetDB()
	} else {
		// 从数据库连接   暂不实现，理论上应该读取EngineMap里的数据库连接
		return GetDB()
	}
}

/**
业财融合数据库链接
*/
func GetBusFaDb() (db *xorm.Engine, err error) {

	//初始化对象
	if EngineMap == nil || EngineMap["busfadb"] == nil {
		// 连接串是空的  重新执行数据库连接获取
		InitConnections()
	}
	db = EngineMap["busfadb"]
	//??? 缺少连接监控程序，防止开发人员连接获取出去以后不进行关闭，GC或mysql 也是长时间不地其关闭
	//注：连接监控时需要获取当前调用本方法的上层go 代码方法，这样才能知道是什么代码没关闭连接

	return db, err
}
