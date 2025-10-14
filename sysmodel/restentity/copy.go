package restentity

import "EasyFinance2020/base/sysmodel"

//查询列表页数据rest
type CopyMainEntity struct {
	PageIndex int                   `xorm:"pageindex" json:"pageindex"`
	PageSize  int                   `xorm:"pagesize" json:"pagesize"`
	SourcePid int                   `xorm:"sourcepid" json:"sourcepid"`
	TargetPid int                   `xorm:"targetpid" json:"targetpid"`
	Filter    []sysmodel.QueryField `xorm:"filter" json:"filter"`
	Ids       string                `xorm:"ids" json:"ids"`
}
