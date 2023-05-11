package invoice

type Transport_invoice_details struct {
	Id           string  `xorm:"id" json:"id"`
	InvoiceCode  string  `xorm:"invoice_code" json:"invoice_code"` // 发票代码
	InvoiceNo    string  `xorm:"invoice_no" json:"invoice_no"`     // 发票号码
	Date         string  `xorm:"date" json:"date"`
	Time         string  `xorm:"time" json:"time"`
	StationFrom  string  `xorm:"station_from" json:"station_from"`   // 出发站
	StationTo    string  `xorm:"station_to" json:"station_to"`       // 到达站
	AmountLittle float64 `xorm:"amount_little" json:"amount_little"` // 合计人民币（小写）
	TaxAmount    float64 `xorm:"tax_amount" json:"tax_amount"`       // 税额
	NoTaxAmount  float64 `xorm:"no_tax_amount" json:"no_tax_amount"` // 不含税金额
	Billno       string  `xorm:"billno" json:"billno"`
	IsSend       int     `xorm:"is_send" json:"is_send"`             // 是否已发送到外部系统 0：未发送 1：已发送
	Pid          string  `xorm:"pid" json:"pid"`                     // pid
	BillType     string  `xorm:"bill_type" json:"bill_type"`         // 单据类型
	CreateDate   string  `xorm:"create_date" json:"create_date"`     // 创建时间
	AttachmentId int     `xorm:"attachment_id" json:"attachment_id"` // 影像图片id
	CompanyId    int     `xorm:"company_id" json:"company_id"`       // 公司ID
	Userid       string  `xorm:"userid" json:"userid"`
}

func (*Transport_invoice_details) TableName() string {
	return "transport_invoice_details"
}
