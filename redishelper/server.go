package redishelper

import (
	"fmt"
	"paas/base/confighelper"
	"paas/base/loghelper"
	"paas/base/util/commutil"
	"time"
)
import "github.com/garyburd/redigo/redis"

//全局变量，用来保存16个库的连接对象
var mapPool map[int]*redis.Pool
var redis_host string
var redis_port int
var password string

//获取服务器配置信息
func GetServerInfo() (string, int, string) {

	if redis_host == "" {
		redis_host = confighelper.GetIniConfig("redisserver", "ip")
		redis_port = commutil.ToInt(confighelper.GetIniConfig("redisserver", "port"))
		password = confighelper.GetIniConfig("redisserver", "password")
	}
	return redis_host, redis_port, password

}

//获取一个实例连接
func GetConn(dbIndex int) redis.Conn {
	if mapPool == nil {
		loghelper.ByHighError("获取redis连接严重错误", "mapPool对象为空，找不到"+commutil.ToString(dbIndex)+"的对象", "")
		return nil
	}

	pol, ok := mapPool[dbIndex]
	if ok {
		return pol.Get()
	} else {
		loghelper.ByHighError("获取redis连接严重错误", "mapPool对象没有指定索引"+commutil.ToString(dbIndex)+"的对象", "")
	}
	return nil
}

// 初始化连接对象
func Init(dbindex int) {

	//服务器信息
	redis_host, redis_port, password := GetServerInfo()

	//首次创建map
	if mapPool == nil {
		mapPool = make(map[int]*redis.Pool)
	}

	//初始化过后不再初始化
	if mapPool[dbindex] != nil {
		return
	}

	//待优化 网上推荐用 gopkg.in/redis.v5

	//创建连接对象
	pool := &redis.Pool{

		//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxIdle: commutil.ToInt(confighelper.GetIniConfig("redisserver", "maxidle")),
		//MaxActive：最大的激活连接数，表示同时最多有N个连接
		MaxActive: commutil.ToInt(confighelper.GetIniConfig("redisserver", "maxactive")),
		//IdleTimeout：最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: 180 * time.Second,
		//当连接超出数量先之后，是否等待到空闲连接释放
		Wait: true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", redis_host, redis_port))
			if err != nil {
				return nil, err
			}
			//设置密码
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			//选择某个库
			if _, err := c.Do("SELECT", dbindex); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}

	//加入map
	mapPool[dbindex] = pool

}
