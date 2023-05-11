package save

type Sys_saveupdatecolfilter struct {
	Id                 int    `xorm:"id" json:"id"`
	RuleId             string `xorm:"ruleid" json:"ruleid"`
	GridId             int    `xorm:"gridid" json:"gridid"`
	IsMain             bool   `xorm:"ismain" json:"ismain"`
	CompareCol         string `xorm:"comparecol" json:"comparecol"`
	Op                 string `xorm:"op" json:"op"`
	CompareStaticValue string `xorm:"comparestaticvalue" json:"comparestaticvalue"`
	CompareRuleMath    string `xorm:"comparerulemath" json:"comparerulemath"`
}

func (*Sys_saveupdatecolfilter) TableName() string {
	return "sys_saveupdatecolfilter"
}
