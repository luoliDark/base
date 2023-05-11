package confighelper

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Unknwon/goconfig"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

//全局变量
var rootPath string
var rootFilePath string
var cfg *goconfig.ConfigFile
var enterpriseid string
var currdb string
var cachedbindex int = -1
var sessiondbindex int = -1
var eventdbindex int = -1
var otherdbindex int = -1
var databaseName string
var isdeving int = -1
var deskey string

func LoadGoEnvPath(dir string) string {
	//return 	strings.Replace(LoadGoEnv(),"config",dir,-1)
	return LoadGoEnv() + dir + "/"
}

//获取环境变量中配置文件存放路径
func LoadGoEnv() string {
	if g.IsEmpty(rootPath) {

		path, err := GetCurrentPath()

		//如果有错不判定文件是否有存在
		if err == nil {
			_, err = os.Stat(path)
		}

		//如果文件不存在从变量读取
		if err != nil {
			fmt.Println(fmt.Sprintf("读取默认配置路径 %v 失败：%v, ", path, err))
			path = os.Getenv("goenv")

			if path == "" {
				errmsg := "没有配置goenv环境变量，该变量用于保存config文件目录，请先配置"
				fmt.Println(errmsg)
				//panic(errmsg)
				path = "/Users/luoli/go/src/BpmServer/config/"
			}
		}

		rootPath = path
	}

	return rootPath
}

//获取数据库名称
func GetDatabaseName() string {
	if databaseName == "" {
		databaseName = GetIniConfig("mysql", "database")
	}
	return databaseName
}

// 获取企业ID
func GetEnterpriseID() string {
	if enterpriseid == "" {
		enterpriseid = GetIniConfig("global", "enterpriseid")
	}
	return enterpriseid
}

//获取当前DB类型
func GetCurrdb() string {
	if currdb == "" {
		currdb = GetIniConfig("global", "currdb")
	}
	return currdb
}

//获取是否开发模式
func GetIsdeving() bool {
	if isdeving == -1 {
		isdeving = toInt(GetIniConfig("global", "isdeving"))
	}
	return toBool(isdeving)
}

//获取当前 加密用的 deskey
func GetDesKey() string {
	if deskey == "" {
		deskey = GetIniConfig("global", "deskey")
	}
	return deskey
}

//获取缓存DB索引
func GetCacheDbIndex() int {
	if cachedbindex == -1 {
		cachedbindex = toInt(GetIniConfig("redisdbindex", "cachedbindex"))
	}
	return toInt(cachedbindex)
}

//获取登录用户session DB索引
func GetSessionDbIndex() int {
	if sessiondbindex == -1 {
		sessiondbindex = toInt(GetIniConfig("redisdbindex", "sessiondbindex"))
	}
	return sessiondbindex
}

//获取事什 DB索引
func GetEventDbIndex() int {
	if eventdbindex == -1 {
		eventdbindex = toInt(GetIniConfig("redisdbindex", "eventdbindex"))
	}
	return eventdbindex
}

//获取其它配置DB索引
func GetOtherDbIndex() int {
	if otherdbindex == -1 {
		otherdbindex = toInt(GetIniConfig("redisdbindex", "otherdbindex"))
	}
	return otherdbindex
}

// 获取config.ini中的配置数据
func GetIniConfig(configName string, key string) string {

	var err error

	if cfg == nil {
		//从当前目录获取配置文件
		currPath := LoadGoEnv()
		currPath += "conf.ini" //ini文件目录 config/config.ini
		cfg, err = goconfig.LoadConfigFile(currPath)
	}

	if err != nil {
		panic(err.Error())
	}

	value, err := cfg.GetValue(configName, key)
	if err != nil {
		fmt.Println("获取ini失败", "未找到"+configName)
		return ""
	}

	return value
}

//获取程序文件路径

func GetEmailConfig() (user, pwd, server, port string, ok bool) {
	user = GetIniConfig("email", "user")
	pwd = GetIniConfig("email", "pwd")
	server = GetIniConfig("email", "server")
	port = GetIniConfig("email", "port")
	ok = true
	if user == "" || pwd == "" || server == "" || port == "" {
		ok = false
	}

	return
}

func GetDingDingConfig() (agentId, appKey, appSecret, accesstokenurl string, ok bool) {
	agentId = GetIniConfig("dingding", "agentId")
	appKey = GetIniConfig("dingding", "appKey")
	appSecret = GetIniConfig("dingding", "appSecret")
	accesstokenurl = GetIniConfig("dingding", "accesstokenurl")
	ok = true
	if agentId == "" || appKey == "" || appSecret == "" {
		ok = false
	}
	return
}

// 转为Int型  注：根据操作系统是32 或64 自动转换
func toInt(str interface{}) int {

	if str == nil || g.IsEmpty(str) {
		return 0
	}

	var re int
	switch vv := str.(type) {
	case string:
		re, _ = strconv.Atoi(vv)
	case int64:
		re = gconv.Int(str)
	case float64:
		re = gconv.Int(str)
	default:
		tmp, ok := str.(int)
		if !ok {
			return 0
		} else {
			re = tmp
		}
	}

	return re

}

//转换类型
func toBool(value interface{}) bool {

	if value == nil {
		return false
	}

	result := false
	switch value := value.(type) {
	case int, int8, int16, int32, int64:

		if value == 1 {
			result = true
		} else {
			result = false
		}

	case string:
		if value == "true" || value == "1" {
			result = true
		} else {
			result = false
		}
	default:
		result = false
	}

	return result
}
