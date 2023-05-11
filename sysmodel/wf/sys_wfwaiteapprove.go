package wf

type Sys_wfwaiteapprove struct {
	WaiteID           string  `xorm:"waiteid pk" json:"waiteid"`
	FlowID            string  `xorm:"flowid" json:"flowid"`
	Pid               int     `xorm:"pid" json:"pid"`
	BillID            string  `xorm:"billid" json:"billid"`
	BillNO            string  `xorm:"billno" json:"billno"`
	SubmitDate        string  `xorm:"submitdate" json:"submitdate"`
	SubmitUid         string  `xorm:"submituid" json:"submituid"`
	TotalMoney        float64 `xorm:"totalmoney" json:"totalmoney"`
	NewGuid           string  `xorm:"newguid" json:"newguid"`
	BUid              string  `xorm:"b_uid" json:"b_uid"`
	BDeptID           string  `xorm:"b_deptid" json:"b_deptid"`
	BBillDate         string  `xorm:"b_billdate" json:"b_billdate"`
	BMemo             string  `xorm:"b_memo" json:"b_memo"`
	RealNumber        string  `xorm:"realnumber" json:"realnumber"`
	BCDefine1         string  `xorm:"b_cdefine1" json:"b_cdefine1"`
	BCDefine2         string  `xorm:"b_cdefine2" json:"b_cdefine2"`
	BCDefine3         string  `xorm:"b_cdefine3" json:"b_cdefine3"`
	BCDefine4         string  `xorm:"b_cdefine4" json:"b_cdefine4"`
	BCDefine5         string  `xorm:"b_cdefine5" json:"b_cdefine5"`
	UserName          string  `xorm:"username" json:"username"`
	B_DeptName        string  `xorm:"b_deptname" json:"b_deptname"`
	BCsId             string  `xorm:"b_csid" json:"b_csid"`
	BCsName           string  `xorm:"b_csname" json:"b_csname"`
	EntId             int     `xorm:"entid" json:"entid"`
	IsResubmit        int     `xorm:"isresubmit" json:"isresubmit"`
	DetailDeptAllName string  `xorm:"detaildeptallname" json:"detaildeptallname"`
	DetailCostAllName string  `xorm:"detailcostallname" json:"detailcostallname"`
	BCompId           string  `xorm:"b_compid" json:"b_compid"`
	BCompName         string  `xorm:"b_compname" json:"b_compname"`
	//EXWFWaiteId       string  `xorm:"exwfwaiteid" json:"exwfwaiteid"`
}

func (*Sys_wfwaiteapprove) TableName() string {
	return "sys_wfwaiteapprove"
}
