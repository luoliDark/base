package form

type Sys_daccesspidlevel struct {
	ID              int    `xorm:"id" json:"id"`
	Pid             int    `xorm:"pid" json:"pid"`
	DataAccessLevel string `xorm:"dataaccesslevel" json:"dataaccesslevel"`
	EntId           int    `xorm:"entid" json:"entid"`
}

func (*Sys_daccesspidlevel) TableName() string {
	return "sys_daccesspidlevel"
}
