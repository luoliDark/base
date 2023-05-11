package form

type Sys_daccessstandbyglobal struct {
	GlobalID   int    `xorm:"globalid" json:"globalid"`
	ConfigCode string `xorm:"configcode" json:"configcode"`
	IsOpen     int    `xorm:"isopen" json:"isopen"`
	Memo       string `xorm:"memo" json:"memo"`
	Entid      int    `xorm:"entid" json:"entid"`
	Issign     int    `xorm:"issign" json:"issign"`
}

func (*Sys_daccessstandbyglobal) TableName() string {
	return "sys_daccessstandbyglobal"
}
