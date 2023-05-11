package dataaccess

type Sys_daccesstop struct {
	DAuthID int    `xorm:"dauthid pk" json:"dauthid"`
	UserID  string `xorm:"userid" json:"userid"`
	RoleID  string `xorm:"roleid" json:"roleid"`
	CompID  string `xorm:"compid" json:"compid"`
	DeptID  string `xorm:"deptid" json:"deptid"`
	JobID   string `xorm:"jobid" json:"jobid"`
	EntId   string `xorm:"entid" json:"entid"`
}

func (*Sys_daccesstop) TableName() string {
	return "sys_daccesstop"
}
