package norule

//单号规则
type Sys_norule struct {
	NoId         string            `xorm:"noid" json:"noid"`
	Pid          int               `xorm:"pid" json:"pid"`
	FormPid      int               `xorm:"formpid" json:"formpid"` //规则引用单据PID
	NoMemo       string            `xorm:"nomemo" json:"nomemo"`
	EntId        int               `xorm:"entid" json:"entid"`
	Noruledetail *Sys_noruledetail `xorm:"-" json:"noruledetaillist"`
}

func (*Sys_norule) TableName() string {
	return "sys_norule"
}
