package wf

/**
 * @Author: yix
 * @Date: 2020/4/27 11:53:50
 * @describe:流程结构体
 */
type Sys_wfflow struct {
	FlowID             string `xorm:"flowid" json:"flowid"`
	FlowCode           string `xorm:"flowcode" json:"flowcode"`
	FlowName           string `xorm:"flowname" json:"flowname"`
	Pid                int    `xorm:"pid" json:"pid"`
	AccessLevel        string `xorm:"accesslevel" json:"accesslevel"`
	UsingExistFlowID   int    `xorm:"usingexistflowid" json:"usingexistflowid"`
	IsOpen             int    `xorm:"isopen" json:"isopen"`
	AutoAppSql         string `xorm:"autoappsql" json:"autoappsql"`
	AutoAppText        string `xorm:"autoapptext" json:"autoapptext"`
	AutoAppStr         string `xorm:"autoappstr" json:"autoappstr"`
	IsChildFlow        int    `xorm:"ischildflow" json:"ischildflow"`
	IsFreeFlow         int    `xorm:"isfreeflow" json:"isfreeflow"`
	RefStepID          string `xorm:"refstepid" json:"refstepid"`
	IsOpenEmail        int    `xorm:"isopenemail" json:"isopenemail"`
	IsOpenWeChat       int    `xorm:"isopenwechat" json:"isopenwechat"`
	IsOpenDingtalk     int    `xorm:"isopendingtalk" json:"isopendingtalk"`
	ConfigWhereFormat  string `xorm:"configwhereformat" json:"configwhereformat"`
	ConfigWhereFromSql string `xorm:"configwherefromsql" json:"configwherefromsql"`
	IsDiscard          int    `xorm:"isdiscard" json:"isdiscard"`
	UIJson             string `xorm:"uijson" json:"uijson"`
	EntId              int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfflow) TableName() string {
	return "sys_wfflow"
}
