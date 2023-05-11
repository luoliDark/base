package wf

/**
 * @describe:
 *
 * @Author: YiXin
 * @Date: 2021/3/21
 */

type Sys_wfwaiteapprove_user struct {
	WaiteID    string  `xorm:"waiteid" json:"waiteid"`
	FlowID     string  `xorm:"flowid" json:"flowid"`
	Pid        int     `xorm:"pid" json:"pid"`
	VerId      int     `xorm:"verid" json:"verid"`
	BillID     string  `xorm:"billid" json:"billid"`
	BillNO     string  `xorm:"billno" json:"billno"`
	SubmitDate string  `xorm:"submitdate" json:"submitdate"`
	SubmitUid  string  `xorm:"submituid" json:"submituid"`
	TotalMoney float64 `xorm:"totalmoney" json:"totalmoney"`
	BUid       string  `xorm:"b_uid" json:"b_uid"`
	BDeptID    string  `xorm:"b_deptid" json:"b_deptid"`
	BBillDate  string  `xorm:"b_billdate" json:"b_billdate"`
	BMemo      string  `xorm:"b_memo" json:"b_memo"`
	Cdefine1   string  `xorm:"cdefine1" json:"cdefine1"`
	Cdefine2   string  `xorm:"cdefine2" json:"cdefine2"`
	Cdefine3   string  `xorm:"cdefine3" json:"cdefine3"`
	Cdefine4   string  `xorm:"cdefine4" json:"cdefine4"`
	Cdefine5   string  `xorm:"cdefine5" json:"cdefine5"`
	UserName   string  `xorm:"username" json:"username"`
	BDeptname  string  `xorm:"b_deptname" json:"b_deptname"`
	ApproveUid string  `xorm:"approve_uid" json:"approve_uid"`
	Pname      string  `xorm:"pname" json:"pname"`
	Newguid    string  `xorm:"newguid" json:"newguid"`
	Entid      int     `xorm:"entid" json:"entid"`
	BCsid      string  `xorm:"b_csid" json:"b_csid"`
	BCsname    string  `xorm:"b_csname" json:"b_csname"`
}

func (*Sys_wfwaiteapprove_user) TableName() string {
	return "sys_wfwaiteapprove_user"
}
