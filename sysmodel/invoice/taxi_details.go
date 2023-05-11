package invoice

type Taxi_details struct {
	Id           string  `xorm:"id" json:"id"`
	InvoiceCode  string  `xorm:"invoice_code" json:"invoice_code"`   // 发票代码
	InvoiceNo    string  `xorm:"invoice_no" json:"invoice_no"`       // 发票号码
	Date         string  `xorm:"date" json:"date"`                   // 日期
	AmountLittle float64 `xorm:"amount_little" json:"amount_little"` // 金额
	TimeGetOn    string  `xorm:"time_get_on" json:"time_get_on"`     // 上车时间
	TimeGetOff   string  `xorm:"time_get_off" json:"time_get_off"`   // 下车时间
	WaitingTime  string  `xorm:"waiting_time" json:"waiting_time"`   // 等候时间
	Mileage      float64 `xorm:"mileage" json:"mileage"`             // 里程
	Place        string  `xorm:"place" json:"place"`                 // 发票所在地
	CreateDate   string  `xorm:"create_date" json:"create_date"`     // 创建时间
	Billno       string  `xorm:"billno" json:"billno"`
	IsSend       int     `xorm:"is_send" json:"is_send"`             // 是否已发送到外部系统 0：未发送 1：已发送
	Pid          string  `xorm:"pid" json:"pid"`                     // pid
	BillType     string  `xorm:"bill_type" json:"bill_type"`         // 单据类型
	AttachmentId int     `xorm:"attachment_id" json:"attachment_id"` // 影像图片id
	CompanyId    int     `xorm:"company_id" json:"company_id"`       // 公司ID
	Userid       string  `xorm:"userid" json:"userid"`
}

func (*Taxi_details) TableName() string {
	return "taxi_details"
}
