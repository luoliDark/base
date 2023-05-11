package matrix

type Sys_matrixaccessbyrule struct {
	AccId         int    `xorm:"accid" json:"accid"`     // 主健
	RangeId       string `xorm:"rangeid" json:"rangeid"` // 对应数据范围的ID
	MatrixId      string `xorm:"matrixid" json:"matrixid"`
	Pid           int    `xorm:"pid" json:"pid"`
	GridId        int    `xorm:"gridid" json:"gridid"`
	FromSqlCol    string `xorm:"fromsqlcol" json:"fromsqlcol"`
	DataSource    int    `xorm:"datasource" json:"datasource"`
	DataSourceCol string `xorm:"datasourcecol" json:"datasourcecol"`
	Entid         int    `xorm:"entid" json:"entid"`
	Ver           string `xorm:"ver" json:"ver"`
}

func (*Sys_matrixaccessbyrule) TableName() string {
	return "sys_matrixaccessbyrule"
}
