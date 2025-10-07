package rediscache

import (
	"strings"

	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/redishelper"
	"github.com/luoliDark/base/util/commutil"
)

// SetString 设置String对象
func SetString(key, val string) bool {
	//返回结果
	return SetStringExpire(key, val, 0)
}

// SetStringExpire 设置String对象
// timeout  秒
func SetStringExpire(key, val string, timeout int) (b bool) {
	key = strings.ToLower(key)

	//获取企业ID,（最新）版本号、dbIndex
	enterpriseID := confighelper.GetEnterpriseID()
	dbIndex, cacherVer := GetCacheVerID(enterpriseID)

	//拼接版本号到企业ID _v1..
	enterpriseID = commutil.AppendStr(enterpriseID, "_v", commutil.ToString(cacherVer))

	//从redis取值
	if timeout > 0 {
		b = redishelper.SetStringExpire(enterpriseID, dbIndex, key, val, timeout)
	} else {
		b = redishelper.SetString(enterpriseID, dbIndex, key, val)
	}
	//返回结果
	return b
}
