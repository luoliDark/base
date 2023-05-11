package matrix

type Sys_matrixaccess struct {
	AccId      int    `xorm:"accid" json:"accid"`
	RangeId    string `xorm:"rangeid" json:"rangeid"` // 对应数据范围的ID
	MatrixId   string `xorm:"matrixid" json:"matrixid"`
	AccObjID   string `xorm:"accobjid" json:"accobjid"`     // 权限ID
	AccObjType string `xorm:"accobjtype" json:"accobjtype"` // 权限类型 role,user,dept,job
	AccName    string `xorm:"accname" json:"accname"`       // 权限用户名称
	Entid      int    `xorm:"entid" json:"entid"`
	Ver        string `xorm:"ver" json:"ver"`
}

func (*Sys_matrixaccess) TableName() string {
	return "sys_matrixaccess"
}
