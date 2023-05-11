package invoice

type Vat_invoice_details struct {
	Id          string `xorm:"id" json:"id"`
	InvoiceCode string `xorm:"invoice_code" json:"invoice_code"` // 发票代码

	InvoiceNo      string  `xorm:"invoice_no" json:"invoice_no"`             // 发票号码
	Date           string  `xorm:"date" json:"date"`                         // 开票日期
	PretaxAmount   float64 `xorm:"pretax_amount" json:"pretax_amount"`       // 合计金额
	CheckCode      string  `xorm:"check_code" json:"check_code"`             // 校验码
	MachineCode    string  `xorm:"machine_code" json:"machine_code"`         // 机器编码
	TaxAmount      float64 `xorm:"tax_amount" json:"tax_amount"`             // 合计税额
	AmountBig      string  `xorm:"amount_big" json:"amount_big"`             // 价税合计(大写）
	AmountLittle   float64 `xorm:"amount_little" json:"amount_little"`       // 价税合计(小写)
	BuyerName      string  `xorm:"buyer_name" json:"buyer_name"`             // 购方名称
	BuyerTaxId     string  `xorm:"buyer_tax_id" json:"buyer_tax_id"`         // 购方纳购方名称税人识别号
	BuyerAddress   string  `xorm:"buyer_address" json:"buyer_address"`       // 购方地址电话
	BuyerBankInfo  string  `xorm:"buyer_bank_info" json:"buyer_bank_info"`   // 购方开户行及账号
	SellerName     string  `xorm:"seller_name" json:"seller_name"`           // 销方名称
	SellerTaxId    string  `xorm:"seller_tax_id" json:"seller_tax_id"`       // 销方纳税人识别号
	SellerAddress  string  `xorm:"seller_address" json:"seller_address"`     // 销方地址电话
	SellerBankInfo string  `xorm:"seller_bank_info" json:"seller_bank_info"` // 销方开户行及账号
	Receiptor      string  `xorm:"receiptor" json:"receiptor"`               // 收款人
	Checker        string  `xorm:"checker" json:"checker"`                   // 复核
	Issuer         string  `xorm:"issuer" json:"issuer"`                     // 开票人
	CipherText     string  `xorm:"cipher_text" json:"cipher_text"`           // 密码区
	SheetType      string  `xorm:"sheet_type" json:"sheet_type"`             // 联次
	InvoiceName    string  `xorm:"invoice_name" json:"invoice_name"`         // 发票名称
	CreateDate     string  `xorm:"create_date" json:"create_date"`           // 创建时间
	CheckState     int     `xorm:"check_state" json:"check_state"`           // 验真状态
	Billno         string  `xorm:"billno" json:"billno"`
	IsSend         int     `xorm:"is_send" json:"is_send"`             // 是否已发送到外部系统 0：未发送 1：已发送
	Pid            string  `xorm:"pid" json:"pid"`                     // pid
	BillType       string  `xorm:"bill_type" json:"bill_type"`         // 单据类型
	AttachmentId   int     `xorm:"attachment_id" json:"attachment_id"` // 影像图片id
	CompanyId      int     `xorm:"company_id" json:"company_id"`       // 公司id
	IsFirst        int     `xorm:"is_first" json:"is_first"`           // 是否定时查验过
	Userid         string  `xorm:"userid" json:"userid"`
	SpName         string  `xorm:"spname" json:"spname"`
}

func (*Vat_invoice_details) TableName() string {
	return "vat_invoice_details"
}
