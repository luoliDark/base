package rediscache

import (
	"base/base/confighelper"
	"base/base/redishelper"
	"base/base/util/commutil"
	"strings"
)

// DelCache 删除缓存
func DelCache(key string) bool {
	key = strings.ToLower(key)
	//获取企业ID,（最新）版本号、dbIndex
	enterpriseID := confighelper.GetEnterpriseID()
	dbIndex, cacherVer := GetCacheVerID(enterpriseID)
	//拼接版本号到企业ID _v1..
	enterpriseID = commutil.AppendStr(enterpriseID, "_v", commutil.ToString(cacherVer))
	return redishelper.DeleteKey(dbIndex, key)
}
