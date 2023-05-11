package invoice

type InvoiceList struct {
	InvocieName   string  `xorm:"invociename" json:"invociename"`
	InvocieType   string  `xorm:"invocietype" json:"invocietype"`
	From          string  `xorm:"from" json:"from"`
	To            string  `xorm:"to" json:"to"`
	Amount_Little float64 `xorm:"amount_little" json:"amount_little"` // 销方名称
	Create_Date   string  `xorm:"create_date" json:"create_date"`     // 销方名称
	Name          string  `xorm:"name" json:"name"`                   // 销方名称
	SellerName    string  `xorm:"seller_name" json:"seller_name"`     // 销方名称
	InvoiceNo     string  `xorm:"invoiceno" json:"invoiceno"`         // 发票号码
	Checked       bool    `xorm:"checked" json:"checked"`
	Id            string  `xorm:"id" json:"id"`
	SpName        string  `xorm:"spname" json:"spname"`
	IsZS          bool    `xorm:"iszs" json:"iszs"`
}

func (*InvoiceList) TableName() string {
	return "InvoiceList"
}
