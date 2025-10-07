/*
	jsz by 2020/2/1
*/

package rediscache

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/db/dbhelper"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/redishelper"
	"github.com/luoliDark/base/redishelper/rediscache/model"
	"github.com/luoliDark/base/sysmodel/logtype"
	"github.com/luoliDark/base/util/commutil"

	redigo "github.com/garyburd/redigo/redis"
	"github.com/gogf/gf/frame/g"
	"github.com/xormplus/xorm"
)

var lockCacheRefer2 sync.Map

var globalXmlData sync.Map //父级XML查询时所生成的主健值

//刷新入口
func LoadRedisMain(userid string, entId, Pid, cType string) (err error) {

	session, _ := conn.GetSession()
	defer session.Close()

	//缓存场景类型
	refreshType, err := getRefreshType(cType)
	if err != nil {
		loghelper.ByHighError(logtype.RedisCheckErr, err.Error(), userid)
		return err
	}

	//配置读取
	arrayConfigs, mapConfigs, err := getRefreshConfig(cType)
	if err != nil {
		loghelper.ByHighError(logtype.RedisCheckErr, err.Error(), userid)
		return err
	}

	//获取所有缓存SQL
	err, xmlCaches := getXmlList(arrayConfigs, refreshType)
	if err != nil {
		loghelper.ByHighError(logtype.RedisCheckErr, err.Error(), userid)
		return err
	}

	// 必要检查
	err = CheckCacheSQL(session, userid, xmlCaches)
	if err != nil {
		loghelper.ByHighError(logtype.RedisCheckErr, err.Error(), userid)
		return err
	}

	//获取新版本号
	dbIndex, currVerKey, currVerNumber := getNextVer(entId, Pid, refreshType)

	//加载到redis
	newPreKey := commutil.AppendStr(currVerKey, "v", currVerNumber)
	err = loadToRedis(session, xmlCaches, dbIndex, newPreKey, mapConfigs, refreshType, entId, Pid)
	if err != nil {
		loghelper.ByHighError(logtype.RedisCheckErr, err.Error(), userid)
		return err
	}

	//修改版本号为最新
	setCacheVer2(currVerKey, currVerNumber, dbIndex) //使版本号自动加1

	loghelper.ByInfo(logtype.Config, "结束初始化缓存 ", userid)

	//删除历史版本 清空1小时以前
	if CheckIsCanClearHistorCache2(currVerKey, dbIndex) {
		ClearHistoryVerCache2(currVerKey, currVerNumber, dbIndex) //异步清空
	}

	return nil

}

//查询配置
func getRefreshConfig(cType string) ([]model.Redis_RefreshConfig, map[string]*model.Redis_RefreshConfig, error) {

	session, _ := conn.GetSession()
	defer session.Close()

	Redis_RefreshConfigs := []model.Redis_RefreshConfig{}
	err := session.Where("ctype=? and IsOpen=1", cType).OrderBy("parentXmlName ASC,SortID ASC ").Find(&Redis_RefreshConfigs)
	if err != nil {
		return nil, nil, err
	}

	mapChches := make(map[string]*model.Redis_RefreshConfig)

	for _, config := range Redis_RefreshConfigs {

		tmpJson := commutil.ToJson(config)

		tmpByte := []byte(tmpJson)
		newObj := model.Redis_RefreshConfig{}
		err = json.Unmarshal(tmpByte, &newObj)

		mapChches[config.XmlName] = &newObj
	}

	return Redis_RefreshConfigs, mapChches, nil

}

//对所有流程进行重新加载
func LoadAllWfCache(entId string) error {

	session, _ := conn.GetSession()
	defer session.Close()

	sql := "SELECT distinct pid from  sys_wfflow where entid=? "
	lst, err := session.QueryString(sql, entId)
	if err != nil {
		return err
	}

	var sbErrs strings.Builder
	if len(lst) > 0 {

		for _, m := range lst {
			pid := m["pid"]
			err = LoadRedisMain("", entId, pid, "wf")
			if err != nil {
				sbErrs.WriteString("pid=")
				sbErrs.WriteString(pid)
				sbErrs.WriteString("缓存刷新失败:")
				sbErrs.WriteString(err.Error())
			}
		}

	}

	if sbErrs.Cap() > 0 {
		return errors.New(sbErrs.String())
	} else {
		return nil
	}

}

//查询类型
func getRefreshType(cType string) (*model.Redis_Refreshtype, error) {

	session, _ := conn.GetSession()
	defer session.Close()

	refreshType := model.Redis_Refreshtype{}
	_, err := session.Where("ctype=?", cType).Get(&refreshType)
	if err != nil {
		return &refreshType, err
	}

	return &refreshType, nil

}

//获取某段XML的SQL
func getXmlList(arrayConfigs []model.Redis_RefreshConfig, refreshType *model.Redis_Refreshtype) (error, []StrCache) {

	//加载 xml文件
	file, err := os.Open(confighelper.LoadGoEnv() + "cache.xml") //jsz 临时代码因为一直报文件找不到
	if err != nil {
		loghelper.ByHighError(logtype.RedisLoadErr, err.Error(), "")
		return err, nil
	}
	defer file.Close()
	//读取xml到变量data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		loghelper.ByHighError(logtype.RedisLoadErr, err.Error(), "")
		return err, nil
	}
	v := Result{}
	//反序列化成对象
	err = xml.Unmarshal(data, &v)
	if err != nil {
		loghelper.ByHighError(logtype.RedisLoadErr, err.Error(), "")
		return err, nil
	}

	caches := []StrCache{}

	if strings.ToLower(refreshType.CType) == "other" {

		//其它XML
		for _, cache := range v.Cache {

			if commutil.ToBool(cache.IsOther) {
				caches = append(caches, cache)
			}
		}

	} else {

		//指定场景XML
		for _, config := range arrayConfigs {

			for _, cache := range v.Cache {

				if cache.Name == config.XmlName {
					caches = append(caches, cache)
				}
			}

		}
	}

	return nil, caches

}

//加载数据到redis
func loadToRedis(session *xorm.Session, Caches []StrCache,
	dbIndex int, keyPreStr string,
	configNames map[string]*model.Redis_RefreshConfig, refreshType *model.Redis_Refreshtype, entId, Pid string) error {

	//redis 连接对象
	c := redishelper.GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//遍历xml所有缓存配置SQL
	for _, val := range Caches {

		fmt.Println(val.Name, "开始.....")

		config := configNames[val.Name] //该XML的缓存配置

		if strings.ToLower(val.Type) == "hash" {

			b := loadCacheByHash2(session, keyPreStr, dbIndex, val, c, config, refreshType, entId, Pid)
			if !b {
				loghelper.ByHighError(logtype.RedisLoadErr, val.QuerySql, "")
			}
		} else if strings.ToLower(val.Type) == "listmap" {
			b := loadCacheByList2(session, keyPreStr, dbIndex, val, c, config, refreshType, entId, Pid)
			if !b {
				loghelper.ByHighError(logtype.RedisLoadErr, val.QuerySql, "")
			}
		} else {
			loghelper.ByHighError(logtype.RedisLoadErr, "未指定有效的type "+val.QuerySql, "")
		}

		fmt.Println(val.Name, "结束.....")
	}

	//从缓冲区发送所有redis到服务器
	err2 := c.Flush()
	if err2 != nil {
		loghelper.ByHighError(logtype.RedisLoadErr, "发送到缓冲区失败", "")
		return err2
	}

	return nil

}

func getPreKey(entId, Pid string, refreshType *model.Redis_Refreshtype) (currVerKey, lastVerKey string) {

	currVerKey = "currver_" + refreshType.CType //例：wf form ...  当前最新版本
	lastVerKey = "lastver_" + refreshType.CType //最后一次刷新的版本号

	//如果按企业设置版本
	if commutil.ToBool(refreshType.IsByEnt) {
		currVerKey += "_" + entId
		lastVerKey += "_" + entId
	}
	if commutil.ToBool(refreshType.IsByPid) {
		currVerKey += "_" + Pid
		lastVerKey += "_" + Pid
	}

	return currVerKey, lastVerKey

}

//获取下一个版本号
func getNextVer(entId, Pid string, refreshType *model.Redis_Refreshtype) (dbIndex int, currVerKey string, currVerNumber int) {

	///获取企业id 及 缓存dbindex 及版本号
	//缓存版本号1，因意外导致只刷新部分缓存未刷新完，版本号2未更新完成。下次刷新还会获取到刷新异常的版本号2，导致同一个key会有重复两份或多份数据。
	//解决方式：记录每次开始刷新的版本号，版本号开始刷新，不管是否刷新完成，下次刷新不能再使用需要获取新版本号

	//获取版本号前缀
	currVerKey, lastVerKey := getPreKey(entId, Pid, refreshType)

	//当前版本号
	dbIndex, currVerNumber = GetCacheVerID2(currVerKey)
	if currVerNumber > 0 {
		currVerNumber = currVerNumber + 1 //版本号加1
	}

	//最新一次刷新记录
	dbIndex, lastVerNumber := GetCacheVerID2(lastVerKey)
	if commutil.ToInt(lastVerNumber) >= currVerNumber {
		currVerNumber = commutil.ToInt(lastVerNumber) + 1 //检查版本号是否已经使用，已使用则在上次已使用版本号基础上再加1
	}

	//先记录最新版本号，不管是否刷新成功，防止刷新失败版本号错乱
	redishelper.SetStringNew(dbIndex, lastVerKey, commutil.ToString(currVerNumber))

	//再保存个空key,用于删除时根据此key过虑
	tmpKey := currVerKey + "v" + commutil.ToString(currVerNumber) + "_"
	redishelper.SetStringNew(dbIndex, tmpKey, commutil.GetNowTime())

	//新的 key前缀
	return dbIndex, currVerKey, currVerNumber

}

//拼接过虑条件
func replaceQuerySql(querySql string, config *model.Redis_RefreshConfig, refreshType *model.Redis_Refreshtype, entId, Pid string) string {

	var filterSql string
	if g.IsEmpty(config.ParentXmlName) {

		//有下级SQL要用本级ID做为过虑条件
		if commutil.ToBool(refreshType.IsByEnt) {
			if g.IsEmpty(config.TableBM) {
				filterSql += " and  entid=" + entId
			} else {
				filterSql += " and " + config.TableBM + ".entid=" + entId
			}
		}

		if commutil.ToBool(refreshType.IsByPid) {
			if g.IsEmpty(config.TableBM) {
				filterSql += " and  pid=" + Pid
			} else {
				filterSql += " and " + config.TableBM + ".pid=" + Pid
			}
		}

	} else if !g.IsEmpty(config.ParentXmlName) && !g.IsEmpty(config.FilterStr) {

		//用父级的主健过虑本级
		parentIdStr, ok := globalXmlData.Load(config.ParentXmlName)
		parentIdStr2 := commutil.ToString(parentIdStr)
		if ok {
			filterSql += " and " + config.FilterStr + " in (" + parentIdStr2 + ")"
		}

	}

	var resSql string
	if !g.IsEmpty(filterSql) {
		indx := strings.Index(querySql, "1=1")

		selectStr := querySql[0 : indx+3]

		orderBy := querySql[indx+3:]

		resSql = selectStr + " " + filterSql + " " + orderBy
	} else {
		resSql = querySql
	}

	return resSql

}

//将值保存到全局对象用于对下级数据过虎用
func savePkToGlobalMap(cacheModel StrCache, config *model.Redis_RefreshConfig, cacheData []map[string]string) error {

	if commutil.ToBool(config.IsHasChild) && len(cacheData) > 0 {

		key := config.XmlName
		var sbValStr strings.Builder

		for i, datum := range cacheData {

			if i > 0 {
				sbValStr.WriteString(",")
			}
			sbValStr.WriteString("'")
			sbValStr.WriteString(datum[cacheModel.KeyField])
			sbValStr.WriteString("'")

		}

		globalXmlData.Store(key, sbValStr.String())
	}

	return nil

}

// 将查询出数据以hase方式加载到redis
func loadCacheByHash2(session *xorm.Session, keyPreStr string,
	dbIndex int, cacheModel StrCache, c redigo.Conn,
	config *model.Redis_RefreshConfig, refreshType *model.Redis_Refreshtype, entId, Pid string) bool {

	//增加PID等过虑条件
	querySql := cacheModel.QuerySql
	if refreshType.CType != "other" {
		querySql = replaceQuerySql(cacheModel.QuerySql, config, refreshType, entId, Pid)
	}

	m, err := dbhelper.QueryByTran(session, "", true, querySql)
	if err != nil {
		loghelper.ByError(logtype.RedisLoadErr, err.Error(), "")
		return false
	}

	//将值保存到全局对象用于对下级数据过虎用
	if refreshType.CType != "other" {
		err = savePkToGlobalMap(cacheModel, config, m)
		if err != nil {
			loghelper.ByError(logtype.RedisLoadErr, err.Error(), "")
			return false
		}
	}

	for _, val := range m {

		//生成hashkey 值
		var hashkey string

		hashkey = commutil.AppendStr(cacheModel.HashKeyPreName, "_", commutil.ToString(val[cacheModel.KeyField]))

		//写入redis
		//插入缓冲区
		for k, v := range val {
			err := c.Send("hmset", commutil.AppendStr(keyPreStr, "_", hashkey), k, v)
			if err != nil {
				loghelper.ByHighError(logtype.RedisLoadErr, "设置hashkey失败"+err.Error(), "")
			}
		}

	}

	return true
}

// 将查询出数据以List方式加载到redis
func loadCacheByList2(session *xorm.Session, keyPreStr string,
	dbIndex int, cacheModel StrCache, c redigo.Conn,
	config *model.Redis_RefreshConfig, refreshType *model.Redis_Refreshtype, entId, Pid string) bool {

	//增加PID等过虑条件
	querySql := cacheModel.QuerySql
	if refreshType.CType != "other" {
		querySql = replaceQuerySql(cacheModel.QuerySql, config, refreshType, entId, Pid)
	}

	m, err := dbhelper.QueryByTran(session, "", true, querySql)
	if err != nil {
		loghelper.ByError(logtype.RedisLoadErr, err.Error(), "")
		return false
	}

	//将值保存到全局对象用于对下级数据过虎用
	if refreshType.CType != "other" {
		err = savePkToGlobalMap(cacheModel, config, m)
		if err != nil {
			loghelper.ByError(logtype.RedisLoadErr, err.Error(), "")
			return false
		}
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

				err := c.Send("rpush", commutil.AppendStr(keyPreStr, "_", currKey), v)

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
		currValueTmp := commutil.AppendStr(keyPreStr, "_", cacheModel.HashKeyPreName) //value未来是作为map的Mapkey值,因为pagefield等是公用，所以增加此变量
		for _, keyField := range keyFieldArr {
			currValueTmp = commutil.AppendStr(currValueTmp, "_", commutil.ToString(val[keyField]))
		}

		lst = append(lst, currValueTmp)

		//最后一组数据
		if currKey == currKeyTmp && index == len(m)-1 {
			//插入缓冲区
			for _, v := range lst {

				err := c.Send("rpush", commutil.AppendStr(keyPreStr, "_", currKey), v)

				if err != nil {
					loghelper.ByHighError(logtype.RedisLoadErr, "设置List失败"+err.Error(), "")

				}
			}
		}
	}

	return true

}

// 版本ID，DB库索引
func GetCacheVerID2(preKeyStr string) (dbIndex int, cacheVer int) {

	//var dbIndex int
	dbIndex = commutil.ToInt(confighelper.GetIniConfig("redisdbindex", "cachedbindex"))

	//var cacheVer int
	// 注：版本号需要根据企业ID进行获取，因为每个企业都有自己的缓存
	tmpVer := redishelper.GetStringNew(preKeyStr, commutil.ToInt(confighelper.GetCacheDbIndex()))
	if tmpVer == "" {
		cacheVer = 1
	} else {
		cacheVer = commutil.ToInt(tmpVer)
	}

	return dbIndex, cacheVer
}

// 设置缓存版本为最新，下次单据加载时会拿到最新代码
// 注：保存信息为 企业id,verid，当前时间
func setCacheVer2(preKey string, currVerNumber, dbIndex int) bool {

	//记录最新版本号
	redishelper.SetStringNew(dbIndex, "version_"+preKey, commutil.ToString(currVerNumber))

	ntime := time.Now().Format(commutil.Time_Fomat01)

	lastUpdateKey := "LastRefreshTime_" + preKey

	redishelper.SetStringNew(dbIndex, lastUpdateKey, ntime)

	return true
}

// 根据时间检查是否可以清除历史缓存
func CheckIsCanClearHistorCache2(preKeyStr string, dbIndex int) bool {

	//检查上次缓存刷新时间是否已超过一小时
	key := "LastRefreshTime_" + preKeyStr
	refreTime := redishelper.GetStringNew(key, dbIndex)

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
func ClearHistoryVerCache2(currVerKey string, currVerNumber, dbIndex int) bool {
	defer commutil.CatchError()

	beginClearCacheVer := currVerNumber - 1 //版本减 1

	if beginClearCacheVer <= 1 {
		return false
	}

	//循环递减历史版本号
	for i := beginClearCacheVer; i >= 1; i-- {
		var keyPreStr string
		// like 24 会把 241 版本的缓存也删掉，需要加上 _ 下划线,指定 24 缓存的版本
		keyPreStr = commutil.AppendStr(currVerKey, "v", commutil.ToString(i)+"_")

		//检查该历史版本是否存在数据，如果存在则模糊查找清除
		re := redishelper.GetStringNew(keyPreStr, dbIndex)

		if !g.IsEmpty(re) {
			redishelper.DeleleByLike(dbIndex, keyPreStr)
		}
	}

	return false
}
