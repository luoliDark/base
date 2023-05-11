package wf

import "time"

type Sys_wfwaiteapproveVo struct {
	WaiteID    string    `xorm:"waiteid pk" json:"waiteid"`
	FlowID     string    `xorm:"flowid" json:"flowid"`
	Pid        int       `xorm:"pid" json:"pid"`
	BillID     string    `xorm:"billid" json:"billid"`
	BillNO     string    `xorm:"billno" json:"billno"`
	SubmitDate time.Time `xorm:"submitdate" json:"submitdate"`
	SubmitUid  string    `xorm:"submituid" json:"submituid"`
	TotalMoney float32   `xorm:"totalmoney" json:"totalmoney"`
	NewGuid    string    `xorm:"newguid" json:"newguid"`
	BUid       string    `xorm:"b_uid" json:"b_uid"`
	BDeptID    string    `xorm:"b_deptid" json:"b_deptid"`
	BBillDate  time.Time `xorm:"b_billdate" json:"b_billdate"`
	BMemo      string    `xorm:"b_memo" json:"b_memo"`
	BCDefine1  string    `xorm:"b_cdefine1" json:"b_cdefine1"`
	BCDefine2  string    `xorm:"b_cdefine2" json:"b_cdefine2"`
	BCDefine3  string    `xorm:"b_cdefine3" json:"b_cdefine3"`
	BCDefine4  string    `xorm:"b_cdefine4" json:"b_cdefine4"`
	BCDefine5  string    `xorm:"b_cdefine5" json:"b_cdefine5"`
	UserName   string    `xorm:"username" json:"username"`
	B_DeptName string    `xorm:"b_deptname" json:"b_deptname"`
	BCsId      string    `xorm:"b_csid" json:"b_csid"`
	BCsName    string    `xorm:"b_csname" json:"b_csname"`
	EntId      int       `xorm:"entid" json:"entid"`
	StepID     string    `xorm:"stepid" json:"stepid"`
	StepAttr   string    `xorm:"stepattr" json:"stepattr"`
}

func (*Sys_wfwaiteapproveVo) TableName() string {
	return "sys_wfwaiteapprove"
}
