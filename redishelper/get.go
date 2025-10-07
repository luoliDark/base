//jsz by 2020.2.2 用于对redis进行取值设定操作

package redishelper

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gogf/gf/frame/g"
	"github.com/luoliDark/base/util/commutil"
)

// 检查是否存在某key
func IsExists(preKeyStr string, dbIndex int, key string) bool {

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//拼接企业ID
	if preKeyStr != "" {
		key = commutil.AppendStr(preKeyStr, "_", key)
	}

	//判断key是否存在
	is_key_exit, err := redis.Bool(c.Do("EXISTS", key))
	if err != nil {
		is_key_exit = false
	}
	return is_key_exit
}

// 获取SET 类型的字符串
func GetStringNew(key string, dbIndex int) string {

	if g.IsEmpty(key) {
		return ""
	}

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//从redis获取
	var re string
	val, err := redis.String(c.Do("GET", key))
	if err != nil {
		re = ""
	} else {
		re = val
	}
	return re
}

// 获取SET 类型的字符串
func GetString(preKeyStr string, dbIndex int, key string) string {

	if g.IsEmpty(key) {
		return ""
	}

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//拼接企业ID
	if preKeyStr != "" {
		key = commutil.AppendStr(preKeyStr, "_", key)
	}

	//从redis获取
	var re string
	val, err := redis.String(c.Do("GET", key))
	if err != nil {
		re = ""
	} else {
		re = val
	}
	return re
}

// 获取HashMap 类型的字符串
func GetHashMap(preKeyStr string, dbIndex int, key string) map[string]string {

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//拼接企业ID
	if preKeyStr != "" {
		key = commutil.AppendStr(preKeyStr, "_", key)
	}

	//从redis获取
	v, err := redis.StringMap(c.Do("HGETALL", key))
	if err != nil {
		v = make(map[string]string)
	}

	return v
}

// 获取ListMap 类型的字符串
func GetListMap(preKeyStr string, dbIndex int, key string) []map[string]string {

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//拼接企业ID
	if preKeyStr != "" {
		key = commutil.AppendStr(preKeyStr, "_", key)
	}

	var lstMap []map[string]string

	//从redis获取
	values, _ := redis.Values(c.Do("lrange", key, "0", "-1"))

	for _, v := range values {

		//获取List中存的value，该value就是hashmap对应的key
		var key string
		key = string(v.([]byte))

		//从redis获取
		v, err := redis.StringMap(c.Do("HGETALL", key))
		if err != nil {
			v = make(map[string]string)
		}

		lstMap = append(lstMap, v)

	}

	return lstMap
}

// 获取ListMap 类型的字符串
func GetList(preKeyStr string, dbIndex int, key string) []string {

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//拼接企业ID
	if preKeyStr != "" {
		key = commutil.AppendStr(preKeyStr, "_", key)
	}

	var lst []string

	//从redis获取
	lst, _ = redis.Strings(c.Do("lrange", key, "0", "-1"))

	return lst
}

// 检查是否存在以xxx开头的ｋｅｙ
func IsExistsByLike(dbIndex int, KeyPreStr string) bool {

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	val, err := redis.Strings(c.Do("KEYS", KeyPreStr+"*"))
	if err != nil {
		return false
	} else {
		if len(val) > 0 {
			return true
		} else {
			return false
		}
	}

}
