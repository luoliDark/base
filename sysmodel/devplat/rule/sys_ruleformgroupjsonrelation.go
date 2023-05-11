package rule

type Sys_ruleformgroupjsonrelation struct {
	Jsonrelationid string `xorm:"jsonrelationid" json:"jsonrelationid"`
	RuleGroupID    string `xorm:"rulegroupid" json:"rulegroupid"`
	Id             string `xorm:"id" json:"id"` // 规则结构id
	Sub            int    `xorm:"sub" json:"sub"`
	Type           int    `xorm:"type" json:"type"`
	Select         int    `xorm:"select" json:"select"`
	Checked        int    `xorm:"checked" json:"checked"`
	Parent         string `xorm:"parent" json:"parent"`
	Relate         string `xorm:"relate" json:"relate"`
	Ruledetailid   string `xorm:"ruledetailid" json:"ruledetailid"` // 规则明细id
	Entid          int    `xorm:"entid" json:"entid"`               // 规则明细id
}

func (*Sys_ruleformgroupjsonrelation) TableName() string {
	return "sys_ruleformgroupjsonrelation"
}
