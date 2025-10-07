package rediscache

import (
	"strings"

	"github.com/luoliDark/base/redishelper/rediscache/model"

	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/redishelper"
	"github.com/luoliDark/base/util/commutil"
)

//特殊场景缓存读取
func GetListMapBySp(entId, Pid int, xmlName, keyValueStr string, refreType *model.Redis_Refreshtype) []map[string]string {

	//获取当前库
	dbIndex := commutil.ToInt(confighelper.GetIniConfig("redisdbindex", "cachedbindex"))

	//获取当前版本号
	currVerKey, _ := getPreKey(commutil.ToString(entId), commutil.ToString(Pid), refreType)
	versionKey := "version_" + currVerKey
	currVerNumber := redishelper.GetStringNew(versionKey, dbIndex)

	//获取缓存数据
	currVerKey = currVerKey + "v" + currVerNumber
	valueStr := xmlName + "_" + keyValueStr
	m := redishelper.GetListMap(currVerKey, dbIndex, valueStr)

	return m

}

// 获取listMap对象
func GetListMap(entId, Pid int, xmlName, keyValueStr string) []map[string]string {

	openNewCache := commutil.ToInt(confighelper.GetIniConfig("global", "opennewcache"))
	if commutil.ToBool(openNewCache) {

		//如果是特殊场景
		if xmlName != "refreshconfig" {

			tmpM := GetHashMap(entId, Pid, "refreshconfig", xmlName)
			if len(tmpM) > 0 {

				refreshType := model.Redis_Refreshtype{IsByEnt: commutil.ToInt(tmpM["isbyent"]), IsByPid: commutil.ToInt(tmpM["isbypid"]), CType: tmpM["ctype"]}
				return GetListMapBySp(entId, Pid, xmlName, keyValueStr, &refreshType)

			}

		}

		//获取当前版本号
		verKey := "version_currver_other"
		dbIndex := commutil.ToInt(confighelper.GetIniConfig("redisdbindex", "cachedbindex"))
		currVerNumber_ByOther := redishelper.GetStringNew(verKey, dbIndex)

		refreType := model.Redis_Refreshtype{CType: "other", IsByPid: 0, IsByEnt: 0}
		currVerKey, _ := getPreKey(commutil.ToString(entId), commutil.ToString(Pid), &refreType)

		currVerKey = currVerKey + "v" + currVerNumber_ByOther
		valueStr := xmlName + "_" + keyValueStr

		m := redishelper.GetListMap(currVerKey, dbIndex, valueStr)

		return m
	} else {

		//老版本
		return getListMap_oldVersion(entId, Pid, xmlName, keyValueStr)
	}

}

// 获取HashMap对象
func GetHashMapBySp(entId, Pid int, xmlName, keyValueStr string, refreshType *model.Redis_Refreshtype) map[string]string {

	//获取当前库
	dbIndex := commutil.ToInt(confighelper.GetIniConfig("redisdbindex", "cachedbindex"))

	//获取当前版本号
	currVerKey, _ := getPreKey(commutil.ToString(entId), commutil.ToString(Pid), refreshType)
	versionKey := "version_" + currVerKey
	currVerNumber := redishelper.GetStringNew(versionKey, dbIndex)

	//获取缓存数据
	currVerKey = currVerKey + "v" + currVerNumber
	valueStr := xmlName + "_" + keyValueStr
	m := redishelper.GetHashMap(currVerKey, dbIndex, valueStr)

	return m

}

// 获取HashMap对象
func GetHashMap(entId, Pid int, xmlName, keyValueStr string) map[string]string {

	openNewCache := commutil.ToInt(confighelper.GetIniConfig("global", "opennewcache"))
	if commutil.ToBool(openNewCache) {

		//如果是特殊场景
		if xmlName != "refreshconfig" {

			tmpM := GetHashMap(entId, Pid, "refreshconfig", xmlName)
			if len(tmpM) > 0 {
				refreshType := model.Redis_Refreshtype{IsByEnt: commutil.ToInt(tmpM["isbyent"]), IsByPid: commutil.ToInt(tmpM["isbypid"]), CType: tmpM["ctype"]}
				return GetHashMapBySp(entId, Pid, xmlName, keyValueStr, &refreshType)
			}

		}

		//获取当前版本号
		verKey := "version_currver_other"
		dbIndex := commutil.ToInt(confighelper.GetIniConfig("redisdbindex", "cachedbindex"))
		currVerNumber_ByOther := redishelper.GetStringNew(verKey, dbIndex)

		refreType := model.Redis_Refreshtype{CType: "other", IsByPid: 0, IsByEnt: 0}
		currVerKey, _ := getPreKey(commutil.ToString(entId), commutil.ToString(Pid), &refreType)

		currVerKey = currVerKey + "v" + currVerNumber_ByOther
		valueStr := xmlName + "_" + keyValueStr

		m := redishelper.GetHashMap(currVerKey, dbIndex, valueStr)

		return m
	} else {

		//老版本
		return getHashMap_OldVersion(entId, Pid, xmlName, keyValueStr)
	}

}

// 获取listMap对象
func getListMap_oldVersion(entId, Pid int, xmlName, keyValueStr string) []map[string]string {

	key := xmlName + "_" + keyValueStr

	key = strings.ToLower(key)

	var lstMap []map[string]string

	//获取企业ID,（最新）版本号、dbIndex
	enterpriseID := confighelper.GetEnterpriseID()
	dbIndex, cacherVer := GetCacheVerID(enterpriseID)

	//拼接版本号到企业ID _v1..
	enterpriseID = commutil.AppendStr(enterpriseID, "_v", commutil.ToString(cacherVer))

	//从redis获取list对象
	lstMap = redishelper.GetListMap(enterpriseID, dbIndex, key)

	//返回结果
	return lstMap

}

// 获取HashMap对象  老版本代码
func getHashMap_OldVersion(entId, Pid int, xmlName, keyValueStr string) map[string]string {

	key := xmlName + "_" + keyValueStr

	key = strings.ToLower(key)

	//获取企业ID,（最新）版本号、dbIndex
	enterpriseID := confighelper.GetEnterpriseID()
	dbIndex, cacherVer := GetCacheVerID(enterpriseID)

	//拼接版本号到企业ID _v1..
	enterpriseID = commutil.AppendStr(enterpriseID, "_v", commutil.ToString(cacherVer))

	//从redis取值
	m := redishelper.GetHashMap(enterpriseID, dbIndex, key)

	//返回结果
	return m

}

// 获取String对象
func GetString(key string) string {

	key = strings.ToLower(key)

	//获取企业ID,（最新）版本号、dbIndex
	enterpriseID := confighelper.GetEnterpriseID()
	dbIndex, cacherVer := GetCacheVerID(enterpriseID)

	//拼接版本号到企业ID _v1..
	enterpriseID = commutil.AppendStr(enterpriseID, "_v", commutil.ToString(cacherVer))

	//从redis取值
	s := redishelper.GetString(enterpriseID, dbIndex, key)

	//返回结果
	return s

}

// 获取String对象
func GetList(key string) []string {

	key = strings.ToLower(key)

	//获取企业ID,（最新）版本号、dbIndex
	enterpriseID := confighelper.GetEnterpriseID()
	dbIndex, cacherVer := GetCacheVerID(enterpriseID)

	//拼接版本号到企业ID _v1..
	enterpriseID = commutil.AppendStr(enterpriseID, "_v", commutil.ToString(cacherVer))

	//从redis取值
	lst := redishelper.GetList(enterpriseID, dbIndex, key)

	//返回结果
	return lst

}

func GetLanguageText(sourceObjID string, languageCode string, sourceDetailID string) (result string) {

	if commutil.IsNullOrEmpty(sourceObjID) || commutil.IsNullOrEmpty(sourceDetailID) {
		return result
	}
	if commutil.IsNullOrEmpty(languageCode) {
		languageCode = "zh"
	}

	redismap := GetHashMap(0, 0, "sys_langdetail", strings.ToLower(sourceObjID)+"_"+strings.ToLower(languageCode)+"_"+strings.ToLower(sourceDetailID))

	if redismap == nil || len(redismap) <= 0 {
		return result
	}
	return commutil.ToString(redismap)
}
