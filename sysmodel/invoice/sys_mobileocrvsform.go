package invoice

type Sys_mobileocrvsform struct {
	Id          int    `xorm:"id" json:"id"`
	Invoicetype string `xorm:"invoicetype" json:"invoicetype"`
	Pid         int    `xorm:"pid" json:"pid"`
	Memo        string `xorm:"memo" json:"memo"`
	CostId      string `xorm:"costid" json:"costid"`
	GridId      int    `xorm:"gridid" json:"gridid"`
}

func (*Sys_mobileocrvsform) TableName() string {
	return "sys_mobileocrvsform"
}
