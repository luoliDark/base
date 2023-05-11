package matrix

type Sys_matrixtype struct {
	MatrixId   string           `xorm:"matrixid" json:"matrixid"`
	MatrixCode string           `xorm:"matrixcode" json:"matrixcode"` // 矩阵类别编码  注：数据权限时一个PID可以使用多个
	MatrixName string           `xorm:"matrixname" json:"matrixname"` // 矩阵类别名称
	TypeId     int              `xorm:"typeid" json:"typeid"`         // 用途类别 1表示数据权限  2表示审批权限   3表示抄送人
	IsOpen     int              `xorm:"isopen" json:"isopen"`         // 是否启用
	Entid      int              `xorm:"entid" json:"entid"`
	Ver        string           `xorm:"ver" json:"ver"`
	SqlColList []*Sys_matrixcol `xorm:"sqlcollist" json:"sqlcollist"`
}

func (*Sys_matrixtype) TableName() string {
	return "sys_matrixtype"
}
