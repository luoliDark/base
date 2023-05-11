package wf

/**
 * @Author: weiyg
 * @Date: 2020/4/12 16:04
 * @describe:
 */

type Sys_wfwaitelooksend struct {
	ID          int    `xorm:"id" json:"id"`
	Pid         int    `xorm:"pid" json:"pid"`
	BillID      string `xorm:"billid" json:"billid"`
	LookUid     string `xorm:"lookuid" json:"lookuid"`
	ApproveDate string `xorm:"approve_date" json:"approve_date"`
	MsgTitle    string `xorm:"msgtitle" json:"msgtitle"`
	MsgBody     string `xorm:"msgbody" json:"msgbody"`
	IsViewed    int    `xorm:"isviewed" json:"isviewed"`
	ViewTime    string `xorm:"viewtime" json:"viewtime"`
	EntId       int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfwaitelooksend) TableName() string {
	return "sys_wfwaitelooksend"
}
