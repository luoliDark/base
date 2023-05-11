package wf

type Sys_wfstepaccessrulefield struct {
	Pid         int    `xorm:"pid" json:"pid"`
	Selectfleld string `xorm:"selectfleld" json:"selectfleld"`
}

func (*Sys_wfstepaccessrulefield) TableName() string {
	return "sys_wfstepaccessrulefield"
}
