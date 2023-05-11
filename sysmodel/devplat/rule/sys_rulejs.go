package rule

type Sys_rulejs struct {
	ObjectID     string `xorm:"objectid pk" json:"objectid"`
	RuleGroupID  string `xorm:"rulegroupid" json:"rulegroupid"`
	Pid          int    `xorm:"pid" json:"pid"`
	GridId       int    `xorm:"gridid" json:"gridid"`
	SqlCol       string `xorm:"sqlcol" json:"sqlcol"`
	ColName      string `xorm:"colname" json:"colname"`
	IsReadonly   int    `xorm:"isreadonly" json:"isreadonly"`
	IsRequired   int    `xorm:"isrequired" json:"isrequired"`
	IsHide       int    `xorm:"ishide"    json:"ishide"`
	IsHideGrid   int    `xorm:"ishidegrid"    json:"ishidegrid"`
	IsReExecuted int    `xorm:"isreexecuted"    json:"isreexecuted"`
	RuleCode     string `xorm:"rulecode"    json:"rulecode"`
	EntId        int    `xorm:"entid" json:"entid"`
	NewValue     string `xorm:"newvalue" json:"newvalue"`
}

func (*Sys_rulejs) TableName() string {
	return "sys_rulejs"
}
