//jsz by 2020.2.2 用于对redis进行赋值设定操作

package redishelper

import (
	"base/loghelper"
	"base/sysmodel/logtype"
	"base/util/commutil"

	"github.com/garyburd/redigo/redis"
)

//SetHashMap set map对象到redis中 针对interface对象
func SetHash(enterpriseID string, dbIndex int, key string, m map[string]interface{}) bool {
	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//拼接企业ID
	if enterpriseID != "" {
		key = commutil.AppendStr(enterpriseID, "_", key)
	}

	//加入redis
	for k, v := range m {
		//加入缓冲区
		err := c.Send("hmset", key, k, v)
		if err != nil {
			loghelper.ByHighError("redis操作", "设置hashkey失败"+err.Error(), "")

		}
	}
	//从缓冲区发送到redis
	c.Flush()

	return true
}

// SetHashMap set map对象到redis中
func SetHashMap(enterpriseID string, dbIndex int, key string, m map[string]string) bool {

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//拼接企业ID
	if enterpriseID != "" {
		key = commutil.AppendStr(enterpriseID, "_", key)
	}

	//加入redis
	for k, v := range m {
		err := c.Send("hmset", key, k, v)
		if err != nil {
			loghelper.ByHighError(logtype.RedisLoadErr, "设置hashkey失败"+err.Error(), "")

		}
	}
	//从缓冲区发送到redis
	c.Flush()

	return true

}

// 加载List对象到redis中
func SetList(enterpriseID string, dbIndex int, key string, arr []string) bool {

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//拼接企业ID
	if enterpriseID != "" {
		key = commutil.AppendStr(enterpriseID, "_", key)
	}

	//加入redis
	for _, v := range arr {
		err := c.Send("rpush", key, v)
		if err != nil {
			loghelper.ByHighError(logtype.RedisLoadErr, "设置List失败"+err.Error(), "")

		}
	}

	//从缓冲区发送redis
	c.Flush()

	return true
}

// 加载String对象到redis中
func SetStringExpire(enterpriseID string, dbIndex int, key string, val string, timeout int) bool {

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//拼接企业ID
	if enterpriseID != "" {
		key = commutil.AppendStr(enterpriseID, "_", key)
	}

	_, err := c.Do("SET", key, val, "EX", timeout)
	if err != nil {
		loghelper.ByHighError(logtype.RedisLoadErr, "设置String失败"+err.Error(), "")
		return false
	}

	return true

}

// 加载String对象到redis中
func SetString(enterpriseID string, dbIndex int, key string, str string) bool {

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	//拼接企业ID
	if enterpriseID != "" {
		key = commutil.AppendStr(enterpriseID, "_", key)
	}

	_, err := c.Do("SET", key, str)
	if err != nil {
		loghelper.ByHighError(logtype.RedisLoadErr, "设置String失败"+err.Error(), "")
		return false
	}

	return true

}

//模糊查找删除
// DelKeyPreStr 表示删除以某字符串开头的全部数据，包含list,map,string
func DeleleByLike(dbIndex int, DelKeyPreStr string) bool {

	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()

	val, err := redis.Strings(c.Do("KEYS", DelKeyPreStr+"*"))
	if err != nil {
		loghelper.ByError(logtype.RedisClearErr, err.Error(), "")
	}

	c.Send("MULTI")
	for i, _ := range val {
		c.Send("DEL", val[i])
	}
	c.Do("EXEC")

	return true
}

// 删除redis 数据by key
func DeleteKey(dbIndex int, delKey string) bool {
	if mapPool[dbIndex] == nil {
		Init(dbIndex)
	}

	//获取连接
	c := GetConn(dbIndex)

	//记得销毁本次链连接
	defer c.Close()
	c.Send("DEL", delKey)
	return true
}
