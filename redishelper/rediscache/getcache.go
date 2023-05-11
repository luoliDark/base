package rediscache

import (
	"strings"

	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/redishelper"
	"github.com/luoliDark/base/util/commutil"
)

// 获取listMap对象
func GetListMap(key string) []map[string]string {

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

// 获取HashMap对象
func GetHashMap(key string) map[string]string {

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

	redismap := GetHashMap("sys_langdetail_" + strings.ToLower(sourceObjID) + "_" + strings.ToLower(languageCode) + "_" + strings.ToLower(sourceDetailID))

	if redismap == nil || len(redismap) <= 0 {
		return result
	}
	return commutil.ToString(redismap)
}
