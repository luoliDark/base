package save

type Sys_saveupdatecol struct {
	RuleId            string `xorm:"ruleid" json:"ruleid"`
	Pid               int    `xorm:"pid" json:"pid"`
	GridId            int    `xorm:"gridid" json:"gridid"`
	IsMain            bool   `xorm:"ismain" json:"ismain"`
	SqlCol            string `xorm:"sqlcol" json:"sqlcol"`
	StaticValue       string `xorm:"staticvalue" json:"staticvalue"`
	RuleMath          string `xorm:"rulemath" json:"rulemath"`
	Memo              string `xorm:"memo" json:"memo"`
	Entid             int    `xorm:"entid" json:"entid"`
	DataSourcce       int    `xorm:"datasourcce" json:"datasourcce"`
	DataSourcceRefCol string `xorm:"datasourccerefcol" json:"datasourccerefcol"`
}

func (*Sys_saveupdatecol) TableName() string {
	return "sys_saveupdatecol"
}
