package wf

/**
 * @Author: hub
 * @Date: 2020/3/18 23:19
 * @describe:审批代理
 */
type Sys_wfconsign struct {
	ConsinID          int    `xorm:"consinid" json:"consinid"`
	AppUid            string `xorm:"appuid" json:"appuid"`                       //审批人/授权人
	ConsignTargentUid string `xorm:"consigntargentuid" json:"consigntargentuid"` // 被授权人,可以审批授权人的单据
	ConsignStatus     int    `xorm:"consignstatus" json:"consignstatus"`         // 1 代理中 0 已关闭
	ConsignBeginTime  string `xorm:"consignbegintime" json:"consignbegintime"`
	ConsignEndTime    string `xorm:"consignendtime" json:"consignendtime"`
	IsByPid           int    `xorm:"isbypid" json:"isbypid"`
	Memo              string `xorm:"memo" json:"memo"`
	Create_Uid        string `xorm:"create_uid" json:"create_uid"`
	Create_Date       string `xorm:"create_date" json:"create_date"`
	Update_Uid        string `xorm:"update_uid" json:"update_uid"`
	Update_Date       string `xorm:"update_date" json:"update_date"`
	IsDisCard         int    `xorm:"isdiscard" json:"isdiscard"`
	DisCard_Uid       string `xorm:"discard_uid" json:"discard_uid"`
	DisCard_Date      string `xorm:"discard_date" json:"discard_date"`
	EntId             int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfconsign) TableName() string {
	return "Sys_wfconsign"
}
