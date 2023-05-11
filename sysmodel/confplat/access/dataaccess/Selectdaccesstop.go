package dataaccess

type Selectdaccesstop struct {
	DAuthID int    `xorm:"dauthid" json:"dauthid"`
	ID      string `xorm:"id" json:"id"`
	Text    string `xorm:"text" json:"text"`
	Code    string `xorm:"code" json:"code"`
}

func (*Selectdaccesstop) TableName() string {
	return "sys_daccesstop"
}
