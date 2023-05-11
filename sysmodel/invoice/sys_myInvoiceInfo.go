package invoice

//个人发票信息
type Sys_MyInvoiceInfo struct {
	InvId       int    `xorm:"invid" json:"invid"`
	UserId      string `xorm:"userid" json:"userid"`
	RateCode    string `xorm:"ratecode" json:"ratecode"`
	CompName    string `xorm:"compname" json:"compname"`
	BankAccount string `xorm:"bankaccount" json:"bankaccount"`
	BankAdd     string `xorm:"bankadd" json:"bankadd"`
	EntId       int    `xorm:"entid" json:"entid"`
}

func (*Sys_MyInvoiceInfo) TableName() string {
	return "sys_myinvoiceinfo"
}
