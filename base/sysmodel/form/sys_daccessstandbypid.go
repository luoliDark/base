package form

type Sys_daccessstandbypid struct {
	ID         int    `xorm:"id" json:"id"`
	Pid        int    `xorm:"pid" json:"pid"`
	ConfigCode string `xorm:"configcode" json:"configcode"`
	IsOpen     int    `xorm:"isopen" json:"isopen"`
	Memo       string `xorm:"memo" json:"memo"`
	Entid      string `xorm:"entid" json:"entid"`
	Issign     int    `xorm:"issign" json:"issign"`
}

func (*Sys_daccessstandbypid) TableName() string {
	return "sys_daccessstandbypid"
}
