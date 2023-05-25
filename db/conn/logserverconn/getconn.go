package logserverconn

import (
	"bytes"

	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/util/commutil"
	_ "github.com/luoliDark/base/util/commutil"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

//采用单例 (注：对象保存的为万恶的指针对象)
var EngineMySql *xorm.Engine

// 初始化数据库连接对象
func InitConnections() {

	// 当前数据库类型

	server := confighelper.GetIniConfig("logserver", "server")
	port := confighelper.GetIniConfig("logserver", "port")
	username := confighelper.GetIniConfig("logserver", "username")
	password := confighelper.GetIniConfig("logserver", "password")
	database := confighelper.GetIniConfig("logserver", "database")

	if server == "" || username == "" || password == "" || database == "" {
		panic("连接数据库失败，服务器地址或账号密码未设置")
	}

	var currDbUrl string
	isdev := commutil.ToBool(confighelper.GetIniConfig("global", "isdeving"))

	if EngineMySql == nil {
		currDbUrl = concatMysql(server, port, username, password, database)
		db, err := xorm.NewEngine("mysql", currDbUrl)
		if err != nil {
			panic("日志DB初始化错误,请检查config.ini")
		}
		EngineMySql = db
		if isdev {
			//开发模式
			db.ShowSQL(true)
		}
	}
}

// GetDBConnection 获取连接对象
func GetDB() (db *xorm.Engine, err error) {

	//初始化对象
	if EngineMySql == nil {
		// 连接串是空的  重新执行数据库连接获取
		InitConnections()
	}

	return EngineMySql, err
}

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

//参考格式 ：sqlserver://username:password@host:port?param1=value&param2=value
//sqlserver://sa:mypass@localhost:1234?database=master&connection+timeout=30 //本地主机上的端口1234
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
	buffer.WriteString("?charset=utf8&multiStatements=true")
	return buffer.String()
}

func concatOracle(server string, port string, username string, password string, database string) string {

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
