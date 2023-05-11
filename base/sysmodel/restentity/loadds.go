package restentity

import "base/base/sysmodel"

//查询列表页主数据rest
type BaseDs struct {
	DS              string `xorm:"ds" json:"ds"`
	Filter          []sysmodel.QueryField
	PageSize        int    `xorm:"pagesize" json:"pagesize"`
	PageIndex       int    `xorm:"pageindex" json:"pageindex"`
	Pid             string `xorm:"pid" json:"pid"`
	Gridid          string `xorm:"gridid" json:"gridid"`
	Colname         string `xorm:"colname" json:"colname"`
	Dsw             string `xorm:"dsw" json:"dsw"`       // 仅用作判断是否有数据源where条件，不取前台的数据，从后台redis取。
	IsListPage      bool   `xorm:"islist" json:"islist"` // 是否不使用dsw，例：列表页查询时 可能不使用dsw
	DSGroupId       string `xorm:"dsgroupid" json:"dsgroupid"`
	ParentID        string `xorm:"parentid" json:"parentid"`
	IsNoCnt         int    `xorm:"isnocnt" json:"isnocnt"`
	DswPar          string `xorm:"dswpar" json:"dswpar"`                   //数据源过虑取界面字段 例：csid={main.csid}
	DsWindowShowCol string `xorm:"dswindowshowcol" json:"dswindowshowcol"` //弹窗选择数据源时的自定义显示字段
	IsShowDiscard   bool   `xorm:"isshowdiscard" json:"isshowdiscard"`
	IsHighSearch    bool   `xorm:"ishighsearch" json:"ishighsearch"`
}
