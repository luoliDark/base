package wf

/**
 * @Author: hub
 * @Date: 2020/3/18 23:19
 * @describe:单据审批代理
 */
type Sys_wfconsignpid struct {
	Id                int    `xorm:"id" json:"id"`
	ConsinID          int    `xorm:"consinid" json:"consinid"`
	Pid               int    `xorm:"pid" json:"pid"`
	AppUid            string `xorm:"appuid" json:"appuid"`
	Create_uid        string `xorm:"create_uid" json:"create_uid"`
	Create_Date       string `xorm:"create_date" json:"create_date"`
	ConsignTargentUid string `xorm:"consigntargentuid" json:"consigntargentuid"`
	IsByWhere         int    `xorm:"isbywhere" json:"isbywhere"`
	EntId             int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfconsignpid) TableName() string {
	return "Sys_wfconsignpid"
}
