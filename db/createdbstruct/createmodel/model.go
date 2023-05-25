package createmodel

type Glo_setprocess struct {
	ProcessId     int    `xorm:"processid" json:"processid"`     // 主健
	IYear         int    `xorm:"iyear" json:"iyear"`             // 结算年份
	IMonth        int    `xorm:"imonth" json:"imonth"`           // 结算月份
	ScenesId      string `xorm:"scenesid" json:"scenesid"`       // 场景ID
	SaleCompId    string `xorm:"salecompid" json:"salecompid"`   // 销售方公司
	PurCompId     string `xorm:"purcompid" json:"purcompid"`     // 采购方公司
	CsId          string `xorm:"csid" json:"csid"`               // 加盟客户
	SettOjbType   int    `xorm:"settojbtype" json:"settojbtype"` // 结算对象类型 1表示内部公司  2表示加盟商
	SettStatus    int    `xorm:"settstatus" json:"settstatus"`   // 结算状态 1表示待结算 2表示已结算 用于在结算进度中显示进度
	OpUserId      string `xorm:"opuserid" json:"opuserid"`       // 操作人
	OpDate        string `xorm:"opdate" json:"opdate"`           // 结算操作日期
	ScenesLevelID string `xorm:"sceneslevelid" json:"sceneslevelid"`
}

func (*Glo_setprocess) TableName() string {
	return "glo_setprocess"
}
