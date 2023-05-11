package matrix

type Sys_matrixdata_h struct {
	RangeId   string `xorm:"rangeid" json:"rangeid"` // 主健 数据范围的ID
	MatrixId  string `xorm:"matrixid" json:"matrixid"`
	RangeName string `xorm:"rangename" json:"rangename"` // 数据范围名称
	Isopen    int    `xorm:"isopen" json:"isopen"`       // 是否启用
	EntId     int    `xorm:"entid" json:"entid"`
	IsAny     int    `xorm:"isany" json:"isany"` // 任意数据 例：某角色可以查看所有数据,任意数据范围时不需要指定匹配字段
}

func (*Sys_matrixdata_h) TableName() string {
	return "sys_matrixdata_h"
}
