package releaseform

type Sys_FListTab struct {
	TabID    string `xorm:"tabid" json:"tabid"`
	TabName  string `xorm:"tabname" json:"tabname"`
	Pid      int    `xorm:"pid" json:"pid"`
	WhereSql string `xorm:"wheresql" json:"wheresql"`
	IsOpen   int    `xorm:"isopen" json:"isopen"`
	Memo     string `xorm:"memo" json:"memo"`
}

func (*Sys_FListTab) TableName() string {
	return "sys_flisttab"
}
