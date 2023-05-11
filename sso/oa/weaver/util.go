package weaver

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"paas/base/db/dbhelper"
)

var ssokey string

func initSsolKey() {
	ssokey, _ = dbhelper.QueryFirstCol("", false, "select ssokey from sys_weaveroaconfig where isopen = 1 ")
	if ssokey == "" {
		panic("单点登录配置未开启，请联系管理员！")
	}
}

func GetToken(loginid string, stamp string) string {
	if ssokey == "" {
		initSsolKey()
	}
	return EncodeToString(ssokey + loginid + stamp)
}

func StringToMD5(data string) string {
	bytes := []byte(data)
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

func EncodeToString(data string) string {
	bytes := []byte(data)
	h := sha1.New()
	h.Write(bytes)
	t := h.Sum(nil)
	return hex.EncodeToString(t)
}
