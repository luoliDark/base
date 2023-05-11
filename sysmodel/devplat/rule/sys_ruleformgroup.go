package rule

type Sys_ruleformgroup struct {
	RuleGroupID string `xorm:"rulegroupid pk" json:"rulegroupid"`
	RuleTypeID  string `xorm:"ruletypeid" json:"ruletypeid"`
	Pid         int    `xorm:"pid" json:"pid"`
	ObjectID    string `xorm:"objectid" json:"objectid"`
	RuleFormat  string `xorm:"ruleformat" json:"ruleformat"`
	RuleText    string `xorm:"ruletext" json:"ruletext"`
	RuleJson    string `xorm:"rulejson" json:"rulejson"`
	EntId       int    `xorm:"entid" json:"entid"`
	CheckType   string `xorm:"checktype" json:"checktype"`
}

func (*Sys_ruleformgroup) TableName() string {
	return "sys_ruleformgroup"
}
