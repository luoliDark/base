package common

import (
	"paas/base/db/conn"
	"paas/base/sso/ssologin/model"
	"paas/base/util/commutil"
)

//插入登录日志
func InsertLoginLog(userId, logMsg, sid, appterminal string) {
	defer commutil.CatchError()
	eng, _ := conn.GetConnection(userId, false)
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