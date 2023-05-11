package form

type Sys_myreq struct {
	LogID             int     `xorm:"logid" json:"logid"`
	Pid               int     `xorm:"pid" json:"pid"`
	Pname             string  `xorm:"pname" json:"pname"`
	BillID            string  `xorm:"billid" json:"billid"`
	BillNo            string  `xorm:"billno" json:"billno"`
	UserId            string  `xorm:"userid" json:"userid"`
	UserName          string  `xorm:"username" json:"username"`
	BillDate          string  `xorm:"billdate" json:"billdate"`
	B_Memo            string  `xorm:"b_memo" json:"b_memo"`
	AppUsers          string  `xorm:"AppUsers" json:"AppUsers"`
	TotalMoney        float64 `xorm:"totalmoney" json:"totalmoney"`
	Create_Date       string  `xorm:"create_date" json:"create_date"`
	B_DeptID          string  `xorm:"b_deptid" json:"b_deptid"`
	B_DeptName        string  `xorm:"b_deptname" json:"b_deptname"`
	EntId             int     `xorm:"entid" json:"entid"`
	PayStatus         int     `xorm:"paystatus" json:"paystatus"`
	DetailDeptAllName string  `xorm:"detaildeptallname" json:"detaildeptallname"`
	DetailCostAllName string  `xorm:"detailcostallname" json:"detailcostallname"`
}

func (*Sys_myreq) TableName() string {
	return "Sys_myreq"
}
