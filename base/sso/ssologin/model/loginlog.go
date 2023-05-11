package ssomodel

import (
	"base/base/db/conn"
	"time"
)

// 登录日志表
// create by zhongxinjian 2020年3月7日20:34:17
type Log_logininfo struct {
	LogID        int       `xorm:"logid bigint auto_increment primary key" json:"logid"`
	LoginName    string    `xorm:"loginname" json:"loginname"`
	LoginPwd     string    `xorm:"loginpwd" json:"loginpwd"`
	LoginTime    time.Time `xorm:"timestamp logintime" json:"logintime"`
	LoginIp      string    `xorm:"loginip" json:"loginip"`
	IsSuccess    int       `xorm:"issuccess int(1)" json:"issuccess"`
	Browser      string    `xorm:"browser" json:"browser"`
	Os           string    `xorm:"os" json:"os"`
	LoginMsg     string    `xorm:"loginmsg" json:"loginmsg"`
	LanguageCode string    `xorm:"languagecode" json:"languagecode"`
	LoginServer  string    `xorm:"loginserver" json:"loginserver"`
	IsMobile     int8      `xorm:"ismobile int(1)" json:"ismobile"`
}

// 记录登录日志
func (log Log_logininfo) RecLoginLog(userid string) {
	engine, _ := conn.GetConnection(userid, true)
	if _, err := engine.Insert(log); err != nil {

	}
}
