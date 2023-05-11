package access

type Sys_daccessmaproleonly struct {
	Id         int    `xorm:"id" json:"id"`
	RoleId     int    `xorm:"roleid" json:"roleid"`
	Pid        int    `xorm:"pid" json:"pid"`
	DataSource int    `xorm:"datasource" json:"datasource"`
	PrimaryKey string `xorm:"primarykey" json:"primarykey"`
	Ver        string `xorm:"ver" json:"ver"`
}

func (*Sys_daccessmaproleonly) TableName() string {
	return "sys_daccessmaproleonly"
}
