package wf

type Sys_wfapprovedbyuser struct {
	LogID             int     `xorm:"logid" json:"logid"`
	WaiteID           string  `xorm:"waiteid" json:"waiteid"`
	FlowID            string  `xorm:"flowid" json:"flowid"`
	Pid               int     `xorm:"pid" json:"pid"`
	VerId             int     `xorm:"-" json:"verid"` //仅用于返回前端。
	BillID            string  `xorm:"billid" json:"billid"`
	BillNo            string  `xorm:"billno" json:"billno"`
	AppUsers          string  `xorm:"AppUsers" json:"AppUsers"`
	StepID            string  `xorm:"stepid" json:"stepid"`
	ResultType        string  `xorm:"resulttype" json:"resulttype"`
	AppOpinion        string  `xorm:"appopinion" json:"appopinion"`
	FlowStatus        int     `xorm:"flowstatus" json:"flowstatus"`
	ActionName        string  `xorm:"actionname" json:"actionname"`
	NewGuid           string  `xorm:"newguid" json:"newguid"`
	AppTerminal       string  `xorm:"appterminal" json:"appterminal"`
	ApproveUid        string  `xorm:"approve_uid" json:"approve_uid"`
	ApproveDate       string  `xorm:"approve_date" json:"approve_date"`
	BMemo             string  `xorm:"b_memo" json:"b_memo"`
	TotalMoney        float64 `xorm:"totalmoney" json:"totalmoney"`
	BDeptName         string  `xorm:"b_deptname" json:"b_deptname"`
	BillDate          string  `xorm:"billdate" json:"billdate"`
	BDeptId           string  `xorm:"b_deptid" json:"b_deptid"`
	BCompId           string  `xorm:"b_compid" json:"b_compid"`
	BCompName         string  `xorm:"b_compname" json:"b_compname"`
	Submitusername    string  `xorm:"submitusername" json:"submitusername"`
	SubmitUid         string  `xorm:"submituid" json:"submituid"`
	CsId              string  `xorm:"csid" json:"csid"`
	CsName            string  `xorm:"csname" json:"csname"`
	Pname             string  `xorm:"pname" json:"pname"`
	EntId             int     `xorm:"entid" json:"entid"`
	DetailDeptAllName string  `xorm:"detaildeptallname" json:"detaildeptallname"`
	DetailCostAllName string  `xorm:"detailcostallname" json:"detailcostallname"`
}

func (*Sys_wfapprovedbyuser) TableName() string {
	return "sys_wfapprovedbyuser"
}
