package wf

type Sys_myreturn struct {
	LogID             int     `xorm:"logid" json:"logid"`
	Pid               int     `xorm:"pid" json:"pid"`
	Pname             string  `xorm:"pname" json:"pname"`
	BillID            string  `xorm:"billid" json:"billid"`
	BillNo            string  `xorm:"billno" json:"billno"`
	AppUsers          string  `xorm:"AppUsers" json:"AppUsers"`
	UserId            string  `xorm:"userid" json:"userid"`
	BillDate          string  `xorm:"billdate" json:"billdate"`
	B_Memo            string  `xorm:"b_memo" json:"b_memo"`
	TotalMoney        float64 `xorm:"totalmoney" json:"totalmoney"`
	CreateDate        string  `xorm:"create_date" json:"create_date"`
	BDeptID           string  `xorm:"b_deptid" json:"b_deptid"`
	BDeptName         string  `xorm:"b_deptname" json:"b_deptname"`
	BCompId           string  `xorm:"b_compid" json:"b_compid"`
	BCompName         string  `xorm:"b_compname" json:"b_compname"`
	B_CsName          string  `xorm:"b_csname" json:"b_csname"`
	UserName          string  `xorm:"username" json:"username"`
	Entid             int     `xorm:"entid" json:"entid"`
	DetailDeptAllName string  `xorm:"detaildeptallname" json:"detaildeptallname"`
	DetailCostAllName string  `xorm:"detailcostallname" json:"detailcostallname"`
}

func (*Sys_myreturn) TableName() string {
	return "sys_myreturn"
}
