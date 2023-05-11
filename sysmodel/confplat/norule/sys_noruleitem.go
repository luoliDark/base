package norule

//规则元素表
type Sys_noruleitem struct {
	RuleItemID       string `xorm:"ruleitemid" json:"ruleitemid"`
	RuleItemName     string `xorm:"ruleitemname" json:"ruleitemname"`
	RefPid           int    `xorm:"refpid" json:"refpid"`
	RefCol           string `xorm:"refcol" json:"refcol"`
	RefDataSoruceCol string `xorm:"refdatasorucecol" json:"refdatasorucecol"`
}

func (*Sys_noruleitem) TableName() string {
	return "sys_noruleitem"
}
