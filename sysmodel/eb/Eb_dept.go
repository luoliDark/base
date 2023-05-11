package eb

type Eb_dept struct {
	DeptID       string `xorm:"deptid" json:"deptid"`
	DeptCode     string `xorm:"deptcode" json:"deptcode"`
	DeptName     string `xorm:"deptname" json:"deptname"`
	ManagerUID   string `xorm:"manageruid" json:"manageruid"`
	ManagerFGUID string `xorm:"manager_fguid" json:"manager_fguid"`
	CompID       string `xorm:"compid" json:"compid"`
	ParentID     string `xorm:"parentid" json:"parentid"`
	SaveSource   string `xorm:"savesource" json:"savesource"`
	ParentCode   string `xorm:"parentcode" json:"parentcode"`
	Entid        int    `xorm:"entid" json:"entid"`
	DidiDeptID   string `xorm:"didideptid" json:"didideptid"`
	IsDiscard    int    `xorm:"isdiscard" json:"isdiscard"`
	DisCardDate  string `xorm:"discard_date" json:"discard_date"`
}

func (*Eb_dept) TableName() string {
	return "eb_dept"
}
