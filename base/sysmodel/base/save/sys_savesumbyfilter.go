package save

type Sys_savesumbyfilter struct {
	MathId    string `xorm:"mathid" json:"mathid"`
	Pid       int    `xorm:"pid" json:"pid"`
	GridId    int    `xorm:"gridid" json:"gridid"`
	FilterStr string `xorm:"filterstr" json:"filterstr"` // 过虑条件 例：Costid in (1,2,3,4)
	MoneyCol  string `xorm:"moneycol" json:"moneycol"`   // 需要合计的字段 例：ybmoney
	Keyword   string `xorm:"keyword" json:"keyword"`     // 聚合关健字 例：sum
}

func (*Sys_savesumbyfilter) TableName() string {
	return "sys_savesumbyfilter"
}
