package form

type Sys_fpage struct {
	Pid               int    `xorm:"pid" json:"pid"`
	Pname             string `xorm:"pname" json:"pname"`
	PrimaryKey        string `xorm:"primarykey" json:"primarykey"`
	PrimaryCode       string `xorm:"primarycode" json:"primarycode"`
	NameCol           string `xorm:"namecol" json:"namecol"`
	SqlTableName      string `xorm:"sqltablename" json:"sqltablename"`
	TableBM           string `xorm:"tablebm" json:"tablebm"`
	EditFromSql       string `xorm:"editfromsql" json:"editfromsql"`
	ListFromSql       string `xorm:"listfromsql" json:"listfromsql"`
	WhereSql          string `xorm:"wheresql" json:"wheresql"`
	QXWhereSql        string `xorm:"qxwheresql" json:"qxwheresql"`
	OrderSql          string `xorm:"ordersql" json:"ordersql"`
	ConfigWhereFormat string `xorm:"configwhereformat" json:"configwhereformat"`
}

func (*Sys_fpage) TableName() string {
	return "sys_fpage"
}
