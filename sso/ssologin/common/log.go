package common

import (
	"github.com/luoliDark/base/db/conn"
	ssomodel "github.com/luoliDark/base/sso/ssologin/model"
	"github.com/luoliDark/base/util/commutil"
)

//插入登录日志
func InsertLoginLog(userId, logMsg, sid, appterminal string) {

	eng, _ := conn.GetDB()
	log := ssomodel.Sys_LogCheckLogin{
		Logid:       commutil.GetUUID(),
		LogMsg:      logMsg,
		Sid:         sid,
		Appterminal: appterminal,
		UserId:      userId,
		InsertDate:  commutil.GetNowTime(),
	}
	_, _ = eng.Insert(log)
}
