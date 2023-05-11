package norule

//单号规则组成明细
type Sys_noruledetail struct {
	Did         string          `xorm:"did" json:"did"`
	RuleType    string          `xorm:"ruletype" json:"ruletype"`     // Static    固定字符
	StaticChar  string          `xorm:"staticchar" json:"staticchar"` // 例：PTBX
	RuleItemID  string          `xorm:"ruleitemid" json:"ruleitemid"`
	NoId        string          `xorm:"noid" json:"noid"`
	AutoNumber  int             `xorm:"autonumber" json:"autonumber"`
	PrefixField int             `xorm:"prefixfield" json:"prefixfield"`
	SortID      int             `xorm:"sortid" json:"sortid"`
	GetMainCol  string          `xorm:"getmaincol" json:"getmaincol"`
	Isrequired  int             `xorm:"isrequired" json:"isrequired"`
	Noruleitem  *Sys_noruleitem `xorm:"-" json:"noruleitem"`
}

func (*Sys_noruledetail) TableName() string {
	return "sys_noruledetail"
}
