package bbs

/**
 * @Author: weiyg
 * @Date: 2020/4/29 8:49
 * @describe:
 */

type Bbsmsgbypid struct {
	MsgType           string   `xorm:"msgtype" json:"msgtype"`
	MsgId             string   `xorm:"msgid" json:"msgid"`
	EntId             string   `xorm:"entid" json:"entid"`
	FormEntId         string   `xorm:"-" json:"-"`
	Pid               int      `xorm:"pid" json:"pid"`
	PName             string   `xorm:"pname" json:"pname"`
	BillNo            string   `xorm:"billno" json:"billno"`
	BillId            string   `xorm:"billid" json:"billid"`
	BillVer           int      `xorm:"billver" json:"billver"`
	Title             string   `xorm:"title" json:"title"`
	Body              string   `xorm:"body" json:"body"`
	SendUid           string   `xorm:"send_uid" json:"send_uid"`
	SendDeptId        string   `xorm:"-" json:"-"`
	UserName          string   `xorm:"username" json:"username"`
	BDeptname         string   `xorm:"b_deptname" json:"b_deptname"`
	TotalMoney        string   `xorm:"totalmoney" json:"totalmoney"`
	SendDate          string   `xorm:"send_date" json:"send_date"`
	IsReply           int      `xorm:"isreply" json:"isreply"`
	ReplyMsgId        int      `xorm:"replymsgid" json:"replymsgid"`
	SendUserIds       []string `xorm:"-" json:"-"`
	DetailCostAllName string   `xorm:"detailcostallname" json:"detailcostallname"`
}

func (*Bbsmsgbypid) TableName() string {
	return "bbsmsgbypid"
}
