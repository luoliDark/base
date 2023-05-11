package matrix

type Sys_matrixdata_v struct {
	WhereId      int    `xorm:"whereid" json:"whereid"` // 主健
	MatrixId     string `xorm:"matrixid" json:"matrixid"`
	RangeId      string `xorm:"rangeid" json:"rangeid"`           // 对应数据范围的ID
	Sqlcol       string `xorm:"sqlcol" json:"sqlcol"`             // 关联字段名
	ColName      string `xorm:"colname" json:"colname"`           // 关联字段中文
	CompareValue string `xorm:"comparevalue" json:"comparevalue"` // 比较值 一个或多个
	CompareText  string `xorm:"comparetext" json:"comparetext"`
	Entid        int    `xorm:"entid" json:"entid"`
	Ver          string `xorm:"ver" json:"ver"`
	Op           string `xorm:"op" json:"op"` // 比较符号 > < = != 等 从sys_matrixcol表中取值保存到此表
}

func (*Sys_matrixdata_v) TableName() string {
	return "sys_matrixdata_v"
}
