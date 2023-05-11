package wf

type Sys_wfalloweditcols struct {
	ID               int    `xorm:"id" json:"id"`
	StepID           string `xorm:"stepid" json:"stepid"` // 为0表示审后修改
	Pid              int    `xorm:"pid" json:"pid"`
	Grid             int    `xorm:"grid" json:"grid"`
	SqlCol           string `xorm:"sqlcol" json:"sqlcol"`
	Name             string `xorm:"name" json:"name"`
	IsMain           int    `xorm:"ismain" json:"ismain"`
	IsEdit           int    `xorm:"isedit" json:"isedit"`
	IsMust           int    `xorm:"ismust" json:"ismust"`
	IsHide           int    `xorm:"ishide" json:"ishide"`
	EditType         string `xorm:"edittype" json:"edittype"`                     // ByStep  审批中
	ArriveStepUpdate string `xorm:"arrive_step_update" json:"arrive_step_update"` // 到达节点审批时更新为
	EntId            int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfalloweditcols) TableName() string {
	return "sys_wfalloweditcols"
}
