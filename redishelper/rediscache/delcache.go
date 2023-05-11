package rediscache

import (
	"strings"

	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/redishelper"
	"github.com/luoliDark/base/util/commutil"
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
