package invoice

type Sys_mobileocrmapfield struct {
	Fieldid      int    `xorm:"fieldid" json:"fieldid"`
	OcrJsonField string `xorm:"ocrjsonfield" json:"ocrjsonfield"`
	SqlField     string `xorm:"sqlfield" json:"sqlfield"`
	IsOpen       int    `xorm:"isopen" json:"isopen"`
	Memo         string `xorm:"memo" json:"memo"`
	InvoiceType  string `xorm:"invoicetype" json:"invoicetype"`
	ConvertType  string `xorm:"converttype" json:"converttype"`
}

func (*Sys_mobileocrmapfield) TableName() string {
	return "sys_mobileocrmapfield"
}
