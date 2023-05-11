package form

type Sys_fpageaccess struct {
	ID     int    `xorm:"ID" json:"ID"`
	RoleID string `xorm:"RoleID" json:"RoleID"`
	Pid    int    `xorm:"Pid" json:"Pid"`
	EntId  int    `xorm:"entid" json:"entid"`
}

func (*Sys_fpageaccess) TableName() string {
	return "sys_fpageaccess"
}
