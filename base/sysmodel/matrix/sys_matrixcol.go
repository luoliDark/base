package matrix

type Sys_matrixcol struct {
	ColId      int    `xorm:"colid" json:"colid"` // 主健 一定要自增按此ID来作务拼接refkey顺序的
	MatrixId   string `xorm:"matrixid" json:"matrixid"`
	Sqlcol     string `xorm:"sqlcol" json:"sqlcol"`         // 关联字段名
	ColName    string `xorm:"colname" json:"colname"`       // 关联字段中文
	Op         string `xorm:"op" json:"op"`                 // 比较符号 例：> < = !=
	IsMultiple int    `xorm:"ismultiple" json:"ismultiple"` // 是否允许多选
	DataSource int    `xorm:"datasource" json:"datasource"` // 数据源PID
	DSGroupID  int    `xorm:"dsgroupid" json:"dsgroupid"`   // 字典类别ID 对应eb_infomain
	Entid      int    `xorm:"entid" json:"entid"`
	Ver        string `xorm:"ver" json:"ver"`
}

func (*Sys_matrixcol) TableName() string {
	return "sys_matrixcol"
}
