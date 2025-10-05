package uservsent

import (
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/luoliDark/base/db/dbhelper"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/redishelper/rediscache"
	"github.com/luoliDark/base/sysmodel/logtype"
)

type Eb_enterprise struct {
	EntID              string `xorm:"entid" json:"entid"`
	EntName            string `xorm:"entname" json:"entname"`
	RegDate            string `xorm:"regdate" json:"regdate"`
	UserQty            int    `xorm:"userqty" json:"userqty"`
	Logo               string `xorm:"logo" json:"logo"`
	Entcode            string `xorm:"entcode" json:"entcode"`
	MenuVer            int    `xorm:"menuver" json:"menuver"`
	Rejister           string `xorm:"rejister" json:"rejister"`
	Rejphone           string `xorm:"rejphone" json:"rejphone"`
	RejAdd             string `xorm:"rejadd" json:"rejadd"`
	FormEntId          int    `xorm:"formentid" json:"formentid"`
	Scan_way           string `xorm:"scan_way" json:"scan_way"`
	IsOpenRefreshLevel string `xorm:"isopenrefreshlevel" json:"isopenrefreshlevel"`
	IsOpenEmailSend    string `xorm:"isopenemailsend" json:"isopenemailsend"`
	IsOpenDingDingSend string `xorm:"isopendingdingsend" json:"isopendingdingsend"`
	IsOpenWeChat       string `xorm:"isopenwechat" json:"isopenwechat"`
	IsOpenWeaverOASend string `xorm:"isopenweaveroasend" json:"isopenweaveroasend"`
	IsPrintImage       int    `xorm:"isprintimage" json:"isprintimage"`
	IsOpenWaitTaskSend int    `xorm:"isopenwaittasksend" json:"isopenwaittasksend"`
	WaitTaskTmplateId  string `xorm:"waittasktmplateid" json:"waittasktmplateid"`
	IsWFFindNextByUser int    `xorm:"iswffindnextbyuser" json:"iswffindnextbyuser"` //查询上级领导按人员档案的所属上级查找
}

func (*Eb_enterprise) TableName() string {
	return "eb_enterprise"
}

//更新版本号
func UpdateMenuVer(entid string) error {
	sql := "update Eb_enterprise set MenuVer=MenuVer+1 where entid=?"
	_, err := dbhelper.ExecSql("", true, sql, entid)
	if err != nil {
		loghelper.ByHighError(logtype.MenuVerErr, err.Error(), "")
		return err
	}
	return nil
}

// 检查该表是否存在平台数据+企业数据共用
func IsSharedTable(tableName string) bool {

	m := rediscache.GetHashMap(0, 0, "sys_pubsqltable", tableName)
	if len(m) > 0 {
		tb := m["sqltable"]
		if !g.IsEmpty(tb) {
			return true
		}
	}

	return false

}

//检查是否是平台系统表，平台系统表不参与企业过虑
func IsSysTable(tableName string) bool {
	re := false
	tableName = strings.ToLower(tableName)
	switch tableName {
	case "sys_info":
		re = true
	case "sys_infomain":
		re = true
	case "sys_fbtn":
		re = true
	case "sys_fsubsystem":
		re = true
	case "sys_fmodel":
		re = true
	case "sys_fpage":
		re = true
	case "sys_fpageid":
		re = true
	case "sys_fpagefield":
		re = true
	case "sys_fgrid":
		re = true
	case "sys_fgridfield":
		re = true
	}
	return re
}
