package restentity

import "github.com/luoliDark/base/sysmodel"

//查询列表页主数据rest
type BaseDs struct {
	DS        string `xorm:"ds" json:"ds"`
	Filter    []sysmodel.QueryField
	PageSize  int    `xorm:"pagesize" json:"pagesize"`
	PageIndex int    `xorm:"pageindex" json:"pageindex"`
	Pid       string `xorm:"pid" json:"pid"`
	Gridid    string `xorm:"gridid" json:"gridid"`
	Colname   string `xorm:"colname" json:"colname"`
	Dsw       string `xorm:"dsw" json:"dsw"`
	DSGroupId string `xorm:"dsgroupid" json:"dsgroupid"`
	ParentID  string `xorm:"parentid" json:"parentid"`
	IsNoCnt   int    `xorm:"isnocnt" json:"isnocnt"`
	DswPar    string `xorm:"dswpar" json:"dswpar"` //数据源过虑取界面字段 例：csid={main.csid}
}
