package matrix

type Sys_matrixvspid struct {
	Mapid    int    `xorm:"mapid" json:"mapid"` // 主健
	Matrixid string `xorm:"matrixid" json:"matrixid"`
	Pid      int    `xorm:"pid" json:"pid"`
	TypeId   int    `xorm:"typeid" json:"typeid"` // 用途类别 1表示数据权限  2表示审批权限   3表示抄送人
	IsOpen   int    `xorm:"isopen" json:"isopen"`
	EntId    int    `xorm:"EntId" json:"EntId"`
	StepId   string `xorm:"stepid" json:"stepid"`
}

func (*Sys_matrixvspid) TableName() string {
	return "sys_matrixvspid"
}
