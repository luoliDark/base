package rediscache

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/db/dbhelper"

	"github.com/xormplus/xorm"
)

func CheckCache(userid string) error {
	//加载 xml文件
	file, err := os.Open(confighelper.LoadGoEnv() + "cache.xml") //jsz 临时代码因为一直报文件找不到
	if err != nil {
		return err
	}
	defer file.Close()
	//读取xml到变量data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	v := Result{}
	//反序列化成对象
	err = xml.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	err = CheckCacheSQL(nil, userid, v.Cache)
	if err != nil {
		return err
	}
	return nil
}

// CheckCacheSQL 检查缓存SQL ，避免刷新后造成缓存异常，
// 例如：发布后，但数据库未更新表结构，刷新缓存！ 造成部门缓存全部失效！
// 造成生成快捷付款单据生成时付款单异常，数据混乱 且恢复困难 ！
func CheckCacheSQL(session *xorm.Session, userid string, Caches []StrCache) (err error) {
	//遍历xml所有缓存配置SQL
	// list也检查 避免语句中也有字段也错误！
	if session == nil {
		session, _ = conn.GetSession()
		defer session.Close()
	}
	bf := bytes.Buffer{}
	for _, val := range Caches {
		// 1=2 会检查SQL正确性，字段是否存在等。 不会查询记录
		QuerySql := fmt.Sprintf("select 1 from (%v) as a where 1=2 ", val.QuerySql)
		_, err = dbhelper.QueryFirstByTran(session, userid, false, QuerySql)
		if err != nil {
			str := " SQL异常语句=" + QuerySql + "\n" + ", err=" + err.Error()
			bf.WriteString(str)
		}
	}
	if bf.Len() > 0 {
		return errors.New(bf.String())
	}
	return err
}
