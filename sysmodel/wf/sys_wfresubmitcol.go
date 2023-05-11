package wf

/**
 * @Author: weiyg
 * @Date: 2020/3/18 23:19
 * @describe:二次直达退回沟选可修改的字段保存表 -结构体
 */
type Sys_wfresubmitcol struct {
	ID               int    `xorm:"id" json:"id"`
	StepID           string `xorm:"stepid" json:"stepid"`
	BillID           string `xorm:"billid" json:"billid"`
	Pid              int    `xorm:"pid" json:"pid"`
	GridID           int    `xorm:"gridid" json:"gridid"`
	SqlCol           string `xorm:"sqlcol" json:"sqlcol"`
	Name             string `xorm:"name" json:"name"`
	IsMain           int    `xorm:"ismain" json:"ismain"`
	ApproveReturnUid string `xorm:"approvereturnuid" json:"approvereturnuid"`
	WaiteID          string `xorm:"waiteid" json:"waiteid"`
	NewGuid          string `xorm:"newguid" json:"newguid"`
	EntId            int    `xorm:"entid" json:"entid"`
}

func (*Sys_wfresubmitcol) TableName() string {
	return "sys_wfresubmitcol"
}
