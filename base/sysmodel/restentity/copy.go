package restentity

import "base/base/sysmodel"

//查询列表页数据rest
type CopyMainEntity struct {
	PageIndex         int                              `xorm:"pageindex" json:"pageindex"`
	PageSize          int                              `xorm:"pagesize" json:"pagesize"`
	SourcePid         int                              `xorm:"sourcepid" json:"sourcepid"`
	TargetPid         int                              `xorm:"targetpid" json:"targetpid"`
	KeyWord           string                           `xorm:"keyword" json:"keyword"` //查询关键字  人员、部门、供应商、备注、单号、单据类型
	Filter            []sysmodel.QueryField            `xorm:"filter" json:"filter"`
	GridListFilter    map[string][]sysmodel.QueryField `xorm:"-" json:"gridListFilter"`
	Ids               string                           `xorm:"ids" json:"ids"`
	CpId              string                           `xorm:"cpid" json:"cpid"`
	IsSecondQueryMain int                              `xorm:"issecondquerymain" json:"issecondquerymain"` //仅用于通过子表查询主表时可用
	SourceGridId      int                              `xorm:"sourcegridid" json:"sourcegridid"`
}
