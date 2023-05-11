package form

type Sys_FormTable struct {
	Pid          int    `xorm:"pid" json:"pid"`
	GridId       int    `xorm:"gridid" json:"gridid"`
	SqlTableName string `xorm:"sqltablename" json:"sqltablename"`
	Name         string `xorm:"name" json:"name"`
	GridName     string `xorm:"gridname" json:"gridname"`
	TableBM      string `xorm:"tablebm" json:"tablebm"`
	IsMain       int    `xorm:"ismain" json:"ismain"`
}

func (*Sys_FormTable) TableName() string {
	return "Sys_FormTable"
}
