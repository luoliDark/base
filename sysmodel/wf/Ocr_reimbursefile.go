package wf

type Ocr_reimbursefile struct {
	ReimburseID      string  `xorm:"reimburseid" json:"reimburseid"`           // 档案id
	ContractCode     string  `xorm:"contractcode" json:"contractcode"`         // 合同流水号
	VoucherCode      string  `xorm:"vouchercode" json:"vouchercode"`           // 凭证编码
	VoucherType      string  `xorm:"vouchertype" json:"vouchertype"`           // 凭证类型
	VoucherPeriod    string  `xorm:"voucherperiod" json:"voucherperiod"`       // 凭证期间
	AccountYear      string  `xorm:"accountyear" json:"accountyear"`           // 会计年份
	AccountMonth     string  `xorm:"accountmonth" json:"accountmonth"`         // 会计月份
	ComCode          string  `xorm:"comcode" json:"comcode"`                   // 公司代码
	ReimburselStatus string  `xorm:"reimburselstatus" json:"reimburselstatus"` // 状态
	PreparationTime  string  `xorm:"preparationtime" json:"preparationtime"`   // 制证时间
	CREATEUSER       string  `xorm:"create_user" json:"create_user"`           // 制单人
	ISDeliveryTICKET string  `xorm:"isdeliveryticket" json:"isdeliveryticket"` // 是否交票
	FilingNo         string  `xorm:"filingno" json:"filingno"`                 // 归档编号
	SourceDataOne    string  `xorm:"sourcedataone" json:"sourcedataone"`       // 来源系统1
	SourceDataTwo    string  `xorm:"sourcedatatwo" json:"sourcedatatwo"`       // 来源系统2
	BillNo           string  `xorm:"billno" json:"billno"`                     // 单据号
	BillDate         string  `xorm:"billdate" json:"billdate"`                 // 单据日期
	Suppliper        string  `xorm:"suppliper" json:"suppliper"`               // 供应商
	SubmitUser       string  `xorm:"submituser" json:"submituser"`             // 提交人
	SubmitDept       string  `xorm:"submitdept" json:"submitdept"`             // 提交部门
	Remark           string  `xorm:"remark" json:"remark"`                     // 备注
	TotalMoney       float64 `xorm:"totalmoney" json:"totalmoney"`             // 总金额
	PID              string  `xorm:"pid" json:"pid"`                           // pid
	BillID           string  `xorm:"billid" json:"billid"`                     // billid
	DivisionID       string  `xorm:"divisionid" json:"divisionid"`             // 分册id
	Entid            string  `xorm:"entid" json:"entid"`                       // 企业ID
	Realnumber       string  `xorm:"realnumber" json:"realnumber"`             // 实物编码
}

func (*Ocr_reimbursefile) TableName() string {
	return "ocr_reimbursefile"
}
