/*
	jsz by 2020/2/1
*/

package rediscache

import (
	"base/base/confighelper"
	"base/base/db/conn"
	"base/base/db/dbhelper"
	"base/base/loghelper"
	"base/base/redishelper"
	"base/base/sysmodel/logtype"
	"base/base/util/commutil"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	redigo "github.com/garyburd/redigo/redis"
	"github.com/gogf/gf/frame/g"
	"github.com/xormplus/xorm"
)

type StrCache struct {
	DBname         string `xml:"DBname,attr"`
	Name           string `xml:"Name,attr"`
	Type           string `xml:"Type,attr"`
	aspect         string `xml:"aspect,attr"`
	KeyField       string `xml:"KeyField,attr"`
	GroupField     string `xml:"GroupField,attr"`
	HashKeyPreName string `xml:"HashKeyPreName,attr"`
	Memo           string `xml:"Memo,attr"`
	QuerySql       string `xml:"QuerySql"`
}

type Result struct {
	XMLName xml.Name `xml:"CacheInfo"`
	Cache   []StrCache
}

// 加载xml中SQL语句查询数据到redis缓存
func LoadRedisCache(userid string) (err error) {
	if isok, err := updateFlagHandle(1); !isok {
		return err
	}
	defer UpdateFlagTheRuning()
	//加载 xml文件
	file, err := os.Open(confighelper.LoadGoEnv() + "cache.xml") //jsz 临时代码因为一直报文件找不到
	if err != nil {
		loghelper.ByHighError(logtype.RedisLoadErr, err.Error(), userid)
		return err
	}
	defer file.Close()
	//读取xml到变量data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		loghelper.ByHighError(logtype.RedisLoadErr, err.Error(), userid)
		return err
	}
	v := Result{}
	//反序列化成对象
	err = xml.Unmarshal(data, &v)
	if err != nil {
		loghelper.ByHighError(logtype.RedisLoadErr, err.Error(), userid)
		return err
	}

	session := conn.GetSession(true)
	defer session.Close()
	// 必要检查
	err = CheckCacheSQL(session, userid, &v)
	if err != nil {
		loghelper.ByHighError(logtype.RedisCheckErr, err.Error(), userid)
		return err
	}

	loghelper.ByInfo(logtype.Config, "开始初始化缓存", userid)
	//企业ID
	enterpriseID := confighelper.GetEnterpriseID()
	//获取企业id 及 缓存dbindex 及版本号
	dbIndex, cacheVer := GetCacheVerID(enterpriseID)
	if cacheVer > 0 {
		cacheVer = cacheVer + 1 //版本号加1
		enterpriseID = commutil.AppendStr(enterpriseID, "_v", commutil.ToString(cacheVer))
	}
	//redis 连接对象
	c := redishelper.GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()
	//遍历xml所有缓存配置SQL
	for _, val := range v.Cache {

		if strings.ToLower(val.Type) == "hash" {
			b := loadCacheByHash(session, userid, enterpriseID, dbIndex, val, c)
			if !b {
				loghelper.ByHighError(logtype.RedisLoadErr, val.QuerySql, userid)
			}
		} else if strings.ToLower(val.Type) == "listmap" {
			b := loadCacheByList(session, userid, enterpriseID, dbIndex, val, c)
			if !b {
				loghelper.ByHighError(logtype.RedisLoadErr, val.QuerySql, userid)
			}
		} else {
			loghelper.ByHighError(logtype.RedisLoadErr, "未指定有效的type "+val.QuerySql, userid)
		}
		fmt.Println(val.Name, "结束.....")
	}

	//从缓冲区发送所有redis到服务器
	err2 := c.Flush()
	if err2 != nil {
		loghelper.ByHighError(logtype.RedisLoadErr, "发送到缓冲区失败", userid)
		return err2
	}

	enid := confighelper.GetEnterpriseID() //从新获取企业id是因为

	//获取上一次加载缓存的时候，如超过20分钟以上则删除历史
	if CheckIsCanClearHistorCache(enid, dbIndex) {
		go ClearHistoryVerCache(enid, cacheVer-1) //异步清空
	}

	//保存最新版本号，及本次刷新时间
	setCacheVer(enid, cacheVer) //使版本号自动加1
	loghelper.ByInfo(logtype.Config, "结束初始化缓存", userid)

	return nil
}

// 将查询出数据以hase方式加载到redis
func loadCacheByHash(session *xorm.Session, userid string, enterpriseID string, dbIndex int, cacheModel StrCache, c redigo.Conn) bool {

	m, err := dbhelper.QueryByTran(session, userid, true, cacheModel.QuerySql)
	if err != nil {
		loghelper.ByError(logtype.RedisLoadErr, err.Error(), "")
		return false
	}

	for _, val := range m {

		//生成hashkey 值
		var hashkey string

		hashkey = commutil.AppendStr(cacheModel.HashKeyPreName, "_", commutil.ToString(val[cacheModel.KeyField]))

		//写入redis
		//插入缓冲区
		for k, v := range val {
			err := c.Send("hmset", commutil.AppendStr(enterpriseID, "_", hashkey), k, v)
			if err != nil {
				loghelper.ByHighError(logtype.RedisLoadErr, "设置hashkey失败"+err.Error(), "")
			}
		}

	}

	return true
}

// 将查询出数据以List方式加载到redis
func loadCacheByList(session *xorm.Session, userid string, enterpriseID string, dbIndex int, cacheModel StrCache, c redigo.Conn) bool {

	m, err := dbhelper.QueryByTran(session, userid, true, cacheModel.QuerySql)
	if err != nil {
		loghelper.ByError(logtype.RedisLoadErr, err.Error(), "")
		return false
	}

	var currKey string                                     //当前key
	groupArr := strings.Split(cacheModel.GroupField, ",")  //分组字段
	keyFieldArr := strings.Split(cacheModel.KeyField, ",") //list对应值的取值字段
	var lst []string                                       //要放放redis的本组数据

	for index, val := range m {

		//获取一组数据的名称 例sys_fapgefield_229001
		currKeyTmp := cacheModel.Name //拼接上名称 例：sys_fpagefield
		for _, groupField := range groupArr {
			currKeyTmp = commutil.AppendStr(currKeyTmp, "_", commutil.ToString(val[groupField]))
		}

		if currKey == "" {
			//第一组
			currKey = currKeyTmp
		} else if currKey != currKeyTmp {

			//如果出现名称不一样说明是第二组数据了
			//将上一组数据放入redis
			//插入缓冲区
			for _, v := range lst {

				err := c.Send("rpush", commutil.AppendStr(enterpriseID, "_", currKey), v)

				if err != nil {
					loghelper.ByHighError(logtype.RedisLoadErr, "设置List失败"+err.Error(), "")

				}
			}

			//修改当前组名为最新的组
			currKey = currKeyTmp

			//清空list对象，用于插入本组的数据
			lst = lst[0:0] //从0切到0 就表示把空值赋给lst对象  以达到清空的目的
		}

		//获取keyfiedl对应的值，拼作为listVlue加入到list对象中
		currValueTmp := commutil.AppendStr(enterpriseID, "_", cacheModel.HashKeyPreName) //value未来是作为map的Mapkey值,因为pagefield等是公用，所以增加此变量
		for _, keyField := range keyFieldArr {
			currValueTmp = commutil.AppendStr(currValueTmp, "_", commutil.ToString(val[keyField]))
		}

		lst = append(lst, currValueTmp)

		//最后一组数据
		if currKey == currKeyTmp && index == len(m)-1 {
			//插入缓冲区
			for _, v := range lst {

				err := c.Send("rpush", commutil.AppendStr(enterpriseID, "_", currKey), v)

				if err != nil {
					loghelper.ByHighError(logtype.RedisLoadErr, "设置List失败"+err.Error(), "")

				}
			}
		}
	}

	return true

}

// 版本ID，DB库索引
func GetCacheVerID(enterpriseID string) (dbIndex int, cacheVer int) {

	//var dbIndex int
	dbIndex = commutil.ToInt(confighelper.GetIniConfig("redisdbindex", "cachedbindex"))

	//var cacheVer int
	// 注：版本号需要根据企业ID进行获取，因为每个企业都有自己的缓存
	tmpVer := redishelper.GetString(enterpriseID, commutil.ToInt(confighelper.GetCacheDbIndex()), "cachever")
	if tmpVer == "" {
		cacheVer = 1
	} else {
		cacheVer = commutil.ToInt(tmpVer)
	}

	return dbIndex, cacheVer
}

// 设置缓存版本为最新，下次单据加载时会拿到最新代码
// 注：保存信息为 企业id,verid，当前时间
func setCacheVer(enterpriseID string, verID int) bool {

	key := "cachever"
	dbIndex := commutil.ToInt(confighelper.GetCacheDbIndex())

	//缓存刷新时间
	redishelper.SetString(enterpriseID, dbIndex, key, commutil.ToString(verID))
	ntime := time.Now().Format(commutil.Time_Fomat01)
	redishelper.SetString(enterpriseID, dbIndex, "cacheRefreshTime", ntime)

	//记录版本标志用来做清空历史记录用
	redishelper.SetString(enterpriseID, dbIndex, "_v"+commutil.ToString(verID)+"isrefresh", ntime)

	return true
}

// 根据时间检查是否可以清除历史缓存
func CheckIsCanClearHistorCache(enterpriseID string, dbIndex int) bool {
	//检查上次缓存刷新时间是否已超过一小时
	key := "cacheRefreshTime"
	refreTime := redishelper.GetString(enterpriseID, dbIndex, key)

	if refreTime == "" {
		return false
	}

	//上次刷新时间
	time1, err := time.Parse(commutil.Time_Fomat01, refreTime)
	if err != nil {
		return false
	}

	//当前时间

	now, _ := time.Parse(commutil.Time_Fomat01, time.Now().Format(commutil.Time_Fomat01))
	diff := now.Sub(time1)
	flag := diff.Minutes()

	if flag > 20 {
		return true //如果超过20分钟
	} else {
		return false
	}

}

// 清除历史版本失效缓存
// 异步清除某企业1小时以前的历史缓存
// beginClearCacheVer 表示开始清除的版本号 需要用当前版本减1
func ClearHistoryVerCache(enterpriseID string, beginClearCacheVer int) bool {
	defer commutil.CatchError()
	//获取当前版本 并减1
	dbIndex, _ := GetCacheVerID(enterpriseID)

	//循环递减历史版本号
	for i := beginClearCacheVer; i >= 1; i-- {
		var keyPreStr string
		// like 24 会把 241 版本的缓存也删掉，需要加上 _ 下划线,指定 24 缓存的版本
		keyPreStr = commutil.AppendStr(enterpriseID, "_v", commutil.ToString(i)+"_")

		//检查该历史版本是否存在数据，如果存在则模糊查找清除
		re := redishelper.GetString(enterpriseID, dbIndex, "_v"+commutil.ToString(i)+"isrefresh")
		if !g.IsEmpty(re) {
			//clear
			redishelper.DeleleByLike(dbIndex, keyPreStr)
		}
	}

	return false
}
