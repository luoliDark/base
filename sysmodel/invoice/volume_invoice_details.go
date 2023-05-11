package invoice

type Volume_invoice_details struct {
	Id           string  `xorm:"id" json:"id"`
	InvoiceCode  string  `xorm:"invoice_code" json:"invoice_code"`   // 发票代码
	InvoiceNo    string  `xorm:"invoice_no" json:"invoice_no"`       // 发票号码
	InvoiceName  string  `xorm:"invoice_name" json:"invoice_name"`   // 发票名称
	Date         string  `xorm:"date" json:"date"`                   // 开票日期
	CheckCode    string  `xorm:"check_code" json:"check_code"`       // 校验码
	AmountBig    string  `xorm:"amount_big" json:"amount_big"`       // 合计金额（大写）
	AmountLittle float64 `xorm:"amount_little" json:"amount_little"` // 合计金额（小写）
	BuyerName    string  `xorm:"buyer_name" json:"buyer_name"`       // 购方名称
	BuyerTaxId   string  `xorm:"buyer_tax_id" json:"buyer_tax_id"`   // 购方纳税人识别号
	SellerName   string  `xorm:"seller_name" json:"seller_name"`     // 销方名称
	SellerTaxId  string  `xorm:"seller_tax_id" json:"seller_tax_id"` // 销方纳税人识别号
	CreateDate   string  `xorm:"create_date" json:"create_date"`     // 创建日期
	CheckState   int     `xorm:"check_state" json:"check_state"`     // 验真状态
	Billno       string  `xorm:"billno" json:"billno"`
	IsSend       int     `xorm:"is_send" json:"is_send"`             // 是否已发送到外部系统 0：未发送 1：已发送
	Pid          string  `xorm:"pid" json:"pid"`                     // pid
	BillType     string  `xorm:"bill_type" json:"bill_type"`         // 单据类型
	AttachmentId int     `xorm:"attachment_id" json:"attachment_id"` // 影像图片id
	CompanyId    int     `xorm:"company_id" json:"company_id"`       // 公司id
	IsFirst      int     `xorm:"is_first" json:"is_first"`           // 是否定时查验
	Userid       string  `xorm:"userid" json:"userid"`
}

func (*Volume_invoice_details) TableName() string {
	return "volume_invoice_details"
}
