package restentity

import "github.com/luoliDark/base/sysmodel"

//查询列表页数据rest
type LoadListEntity struct {
	PageIndex int                   `xorm:"pageindex" json:"pageindex"`
	PageSize  int                   `xorm:"pagesize" json:"pagesize"`
	Pid       int                   `xorm:"pid" json:"pid"`
	Filter    []sysmodel.QueryField `xorm:"filter" json:"filter"`
}
