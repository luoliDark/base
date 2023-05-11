package eb

type Eb_dept_col2row struct {
	Id     int    `xorm:"id" json:"id"`
	Deptid string `xorm:"deptid" json:"deptid"`
	Userid string `xorm:"userid" json:"userid"` // 数据源ID，例如:用户的ID值
	Tag    string `xorm:"tag" json:"tag"`       // 数据类型分类
	Memo   string `xorm:"memo" json:"memo"`     // 备注
	EntId  int    `xorm:"entid" json:"entid"`
}

func (*Eb_dept_col2row) TableName() string {
	return "eb_dept_col2row"
}
