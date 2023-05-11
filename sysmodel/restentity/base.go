package restentity

import "github.com/luoliDark/base/sysmodel"

//查询列表页数据rest
type LoadListEntity struct {
	PageIndex   int                    `xorm:"pageindex" json:"pageindex"`
	PageSize    int                    `xorm:"pagesize" json:"pagesize"`
	AppId       string                 `xorm:"appid" json:"appid"`
	Secret      string                 `xorm:"secret" json:"secret"`
	Pid         int                    `xorm:"pid" json:"pid"`
	KeyWord     string                 `xorm:"keyword" json:"keyword"` //查询关键字  人员、部门、供应商、备注、单号、单据类型
	Filter      []sysmodel.QueryField  `xorm:"filter" json:"filter"`
	ExtraFilter map[string]QueryEntity `xorm:"extrafilter" json:"extrafilter"` //延伸 拓展 额外 extra
	Order       string                 `xorm:"order" json:"order"`
}

type QueryEntity struct {
	Ds         string                `xorm:"ds" json:"ds"` //数据源PID
	Col        string                `xorm:"col" json:"col"`
	FilterList []sysmodel.QueryField `xorm:"filterlist" json:"filterlist"`
}
