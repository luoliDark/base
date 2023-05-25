package db

import (
	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/util/commutil"
)

const ApiLog = "ApiLog"         //外部调用API LOG
const DeleteLog = "DeleteLog"   //删除操作 LOG
const SaveLog = "SaveLog"       //保存日志 LOG
const EventLog = "EventLog"     //系统事件执行LOG
const VoucherLog = "VoucherLog" //凭证相关LOG
const PayLog = "PayLog"         //付款相关LOG
const ExWfLog = "ExWfLog"       //外部工作流LOG
const BudLog = "BudLog"         //外部工作流LOG

//DB LOG 对象
type DBLogEntity struct {
	Pid        int    //单据类型
	PrimaryKey string //主健
	IsSuccess  bool   //是否操作成功
	LogMsg     string //消息
	OpUserId   string //操作人
	LogTime    string //日志时间
	LogType    string //消息类型
	Action     string //操作类型 save,allpass,appreturn,apppass,delete,submit
}

type insertEntity struct {
	LogID      string `xorm:"logid" json:"logid"`
	Pid        int    `xorm:"pid" json:"pid"`
	PrimaryKey string `xorm:"primarykey" json:"primarykey"`
	IsSuccess  int    `xorm:"issuccess" json:"issuccess"`
	LogMsg     string `xorm:"logmsg" json:"logmsg"`
	LogTime    string `xorm:"logtime" json:"logtime"`
	LogType    string `xorm:"logtype" json:"logtype"`
	Action     string `xorm:"action" json:"action"`
	OpUserId   string `xorm:"opuserid" json:"opuserid"`
}

func InsertLog(entity *DBLogEntity) {

	var sqlTable string
	switch entity.LogType {
	case EventLog:
		sqlTable = "log_sysevent"
	case SaveLog:
		sqlTable = "log_syssave"
	case DeleteLog:
		sqlTable = "log_sysdelete"
	case ApiLog:
		sqlTable = "log_sysapi"
	case VoucherLog:
		sqlTable = "log_voucher"
	case PayLog:
		sqlTable = "log_pay"
	case ExWfLog:
		sqlTable = "log_exwf"
	case BudLog:
		sqlTable = "log_bud"
	}

	insertE := insertEntity{
		LogID:      commutil.GetUUID(),
		Pid:        entity.Pid,
		PrimaryKey: entity.PrimaryKey,
		LogMsg:     entity.LogMsg,
		LogTime:    entity.LogTime,
		LogType:    entity.LogType,
		Action:     entity.Action,
		OpUserId:   entity.OpUserId,
	}
	if entity.IsSuccess {
		insertE.IsSuccess = 1
	}

	go func(entity *insertEntity, sqlTable string) {

		session, err := conn.GetSession()
		if err == nil {
			defer session.Close()
			_, _ = session.Table(sqlTable).Insert(entity)
		}

	}(&insertE, sqlTable)

}
